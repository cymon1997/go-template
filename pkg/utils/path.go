package utils

import (
	"path/filepath"
	"runtime"
)

func GetRootPath() string {
	_, caller, _, _ := runtime.Caller(1)
	return filepath.Join(filepath.Dir(caller), "../..")
}
