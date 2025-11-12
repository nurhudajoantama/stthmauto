package monitoring

import (
	"context"
	"errors"
	"time"

	"github.com/cenkalti/backoff/v4"
	"github.com/nurhudajoantama/stthmauto/app/hmstt"
	"github.com/nurhudajoantama/stthmauto/app/worker"
	"github.com/nurhudajoantama/stthmauto/internal/config"
	"github.com/rs/zerolog/log"
)

type MonitoringWorker struct {
	service       *hmstt.HmsttService
	intercheckCfg config.InternetCheck
}

func RegisterWorkers(s *worker.Worker, svc *hmstt.HmsttService, intercheckCfg config.InternetCheck) {
	hw := &MonitoringWorker{
		service:       svc,
		intercheckCfg: intercheckCfg,
	}

	s.Go(func(ctx context.Context) func() error {
		return hw.internetWorker(ctx)
	})
}

func (w *MonitoringWorker) internetWorker(ctx context.Context) func() error {
	return func() error {
		interval, err := time.ParseDuration(w.intercheckCfg.Interval)
		if err != nil {
			log.Error().Err(err).Msg("invalid internet check interval duration, using default 1 minute")
			interval = 2 * time.Minute
		}

		ticker := time.NewTicker(interval)
		defer ticker.Stop()

		for {
			select {
			case <-ctx.Done():
				log.Info().Msg("hmstt internet worker stopped")
				return nil
			case <-ticker.C:
				pingCheckNetOk := pingInternet(w.intercheckCfg.CheckAddress)
				if !pingCheckNetOk {
					log.Print("modem connection is down, just wait")
					err := w.internetWorkerSwitchModem(ctx)
					if err != nil {
						log.Error().Err(err).Msg("hmstt internet worker switch error")
					}
				}
			}
		}
	}
}

func (w *MonitoringWorker) internetWorkerSwitchModem(ctx context.Context) error {
	exp := backoff.NewExponentialBackOff()
	exp.InitialInterval = 30 * time.Second
	exp.MaxInterval = 10 * time.Minute
	exp.MaxElapsedTime = 0
	exp.RandomizationFactor = 0.3
	exp.Multiplier = 3.0

	bo := backoff.WithContext(exp, ctx)

	return backoff.Retry(func() error {

		pingCheckModemOk := pingInternet(w.intercheckCfg.ModemAddress)
		if !pingCheckModemOk {
			log.Print("modem connection is down")
			return errors.New("modem connection is down, cannot restart modem (will retry)")
		}

		pingCheckNetOk := pingInternet(w.intercheckCfg.CheckAddress)
		if pingCheckNetOk {
			log.Print("internet connection is down")
			return nil
		}

		log.Print("internet connection is down, restarting modem")

		err := w.service.RestartModem(ctx)
		if err != nil {
			log.Printf("restart modem failed: %v (will retry)", err)
			return err
		}
		log.Print("restart modem success")

		return errors.New("internet still down after modem restart (will retry)")
	}, bo)

}
