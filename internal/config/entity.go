package config

import (
	"github.com/cymon1997/go-template/pkg/db"
	"github.com/cymon1997/go-template/pkg/migrate"
)

type MainConfig struct {
	App      AppConfig
	DB       db.Config
	Migrator migrate.Config
	Seeder   migrate.Config
}

type AppConfig struct {
	Host string
	Port int
}
