package main

import (
	"context"
	"os"
	"os/signal"
	"syscall"
	"time"

	"github.com/nurhudajoantama/stthmauto/app/hmstt"
	"github.com/nurhudajoantama/stthmauto/app/server"
	"github.com/nurhudajoantama/stthmauto/internal/bbolt"
	"github.com/nurhudajoantama/stthmauto/internal/config"
	"github.com/nurhudajoantama/stthmauto/internal/instrumentation"

	log "github.com/rs/zerolog/log"
)

func main() {
	// initialize config
	configPath := os.Getenv("CONFIG_PATH")
	if configPath == "" {
		configPath = "conf/conf.yaml"
	}
	config, err := config.InitializeConfig(configPath)
	if err != nil {
		log.Fatal().Err(err).Msg("failed to load config")
	}

	// initialize logger
	cleanupLog := instrumentation.InitializeLogger(config.Log)
	defer cleanupLog()

	// initialize bbolt
	bboltDB := bbolt.InitializeBolt(config.KV)

	// initialize server
	srv := server.New(config.HTTP.Addr())

	// HTSTT
	{
		hmsttStore := hmstt.NewStore(bboltDB)
		hmsttService := hmstt.NewService(hmsttStore)
		hmstt.RegisterHandlers(srv, hmsttService)
	}

	// start server implemented graceful shutdown
	go func() {
		log.Info().Msgf("starting server on %s", config.HTTP.Addr())
		if err := srv.Start(); err != nil {
			log.Error().Err(err).Msg("server stopped with error")
		}
	}()

	quit := make(chan os.Signal, 1)
	signal.Notify(quit, os.Interrupt, syscall.SIGTERM, syscall.SIGQUIT)

	<-quit

	log.Info().Msg("shutting down server...")
	{
		gracefulPeriod := 10 * time.Second
		shutdownCtx, cancel := context.WithTimeout(context.Background(), gracefulPeriod)
		defer cancel()

		if err := srv.Shutdown(shutdownCtx); err != nil {
			log.Error().Err(err).Msg("failed to gracefully shutdown server")
		}
		log.Info().Msg("server stopped")

		bbolt.CloseBolt(shutdownCtx, bboltDB)
	}
}
