package config

import "github.com/cymon1997/go-template/pkg/migrate"

type MainConfig struct {
	App      AppConfig
	Migrator migrate.Config
	Seeder   migrate.Config
}

type AppConfig struct {
	Host string
	Port int
}
