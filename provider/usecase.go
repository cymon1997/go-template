package provider

import (
	"sync"

	"github.com/cymon1997/go-template/core/usecase"
)

var (
	uc     usecase.Usecase
	syncUc sync.Once
)

func GetUsecase() usecase.Usecase {
	syncUc.Do(func() {
		uc = usecase.New()
	})
	return uc
}
