package postgres

import (
	"time"

	log "github.com/rs/zerolog/log"

	"github.com/nurhudajoantama/stthmauto/internal/config"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

func NewGorm(c config.SQL) *gorm.DB {
	db, err := gorm.Open(postgres.Open(c.DataSourceName()), &gorm.Config{})
	if err != nil {
		log.Fatal().Err(err).Msg("failed to connect to database")
	}

	sqlDB, err := db.DB()
	if err != nil {
		log.Fatal().Err(err).Msg("failed to get database instance")
	}

	sqlDB.SetMaxIdleConns(c.MaxIdleConn)
	sqlDB.SetMaxOpenConns(c.MaxOpenConn)
	sqlDB.SetConnMaxLifetime(time.Hour)

	return db
}
