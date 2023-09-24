package main

import (
	"fmt"
	"github.com/cymon1997/go-logger"
	"github.com/cymon1997/go-template/provider"
)

func main() {
	cfg := provider.GetMainConfig()

	provider.Init(cfg)
	logger.WithMeta(logger.Meta{
		"cfg": cfg,
	}).Debug("init config")

	err := provider.GetMigrator().Up()
	if err != nil {
		logger.WithError(err).Fatal("error apply db migrations")
	}

	// Note: uncomment if there's seeders
	//err = provider.GetSeeder().Up()
	//if err != nil && err != fs.ErrNotExist {
	//	logger.WithError(err).Fatal("error apply db seeds")
	//}

	logger.Infof("serving http at %s:%d\n", cfg.Server.Host, cfg.Server.Port)
	err = provider.GetHttpServer().Run(fmt.Sprintf(":%d", cfg.Server.Port))
	if err != nil {
		logger.WithError(err).Fatal("error serving http: ", err)
	}
}
