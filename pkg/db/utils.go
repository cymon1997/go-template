package db

import (
	"time"

	"gorm.io/gorm"

	"gorm.io/driver/mysql"
	"gorm.io/driver/postgres"
)

func getDialector(cfg Config) gorm.Dialector {
	switch cfg.Driver {
	case MySQL:
		return mysql.Open(cfg.DSN)
	case Postgres:
		return postgres.Open(cfg.DSN)
	default:
		return nil
	}
}

func getNowFunc(cfg Config) func() time.Time {
	if cfg.Timezone != "" {
		InitTimezone(cfg.Timezone)
		return NowTZ
	}
	return NowUTC
}
