package provider

import (
	"github.com/cymon1997/go-logger"
	"github.com/cymon1997/go-template/internal/config"
	iCtx "github.com/cymon1997/go-template/internal/context"
)

func Init(cfg *config.MainConfig) {
	cfg.Logger.IDKey = iCtx.IDKey
	logger.Init(cfg.Logger)
}
