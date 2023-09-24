package config

import (
	"github.com/cymon1997/go-logger"
	"github.com/cymon1997/go-template/pkg/db"
	"github.com/cymon1997/go-template/pkg/migrate"
)

const (
	FilePath = "/files/config"
)

type MainConfig struct {
	Server   ServerConfig
	DB       db.Config
	Logger   logger.Config
	Migrator migrate.Config
	Seeder   migrate.Config
}

type ServerConfig struct {
	Host string
	Port int
}
