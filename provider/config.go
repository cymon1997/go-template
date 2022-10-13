package provider

import (
	"sync"

	"github.com/cymon1997/go-template/internal/config"
)

var (
	mainConfig     *config.MainConfig
	syncMainConfig sync.Once
)

func GetMainConfig() *config.MainConfig {
	syncMainConfig.Do(func() {
		mainConfig = config.New()
	})
	return mainConfig
}
