package provider

import (
	"github.com/cymon1997/go-logger"
	"sync"

	iCfg "github.com/cymon1997/go-template/internal/config"
	"github.com/cymon1997/go-template/pkg/config"
	"github.com/cymon1997/go-template/pkg/utils"
)

var (
	mainConfig     *iCfg.MainConfig
	syncMainConfig sync.Once
)

func GetMainConfig() *iCfg.MainConfig {
	syncMainConfig.Do(func() {
		err := config.New(&mainConfig, config.Param{
			Name:   "config",
			Path:   utils.GetRootPath() + iCfg.FilePath,
			Prefix: "",
		})
		if err != nil {
			logger.Fatal("error load config: ", err)
		}
	})
	return mainConfig
}
