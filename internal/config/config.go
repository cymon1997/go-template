package config

import (
	"log"

	"github.com/cymon1997/go-template/pkg/utils"

	"github.com/cymon1997/go-template/pkg/config"
)

const (
	configPath = "/files/config"
)

func New() *MainConfig {
	var cfg MainConfig
	err := config.New(&cfg, config.Param{
		Name:   "config",
		Path:   utils.GetRootPath() + configPath,
		Prefix: "",
	})
	if err != nil {
		log.Fatal("error load config: ", err)
	}
	return &cfg
}
