package provider

import (
	"sync"

	"github.com/cymon1997/go-template/internal/ping"
)

var (
	instances     []ping.Pingable
	syncInstances sync.Once
)

func GetInstances() []ping.Pingable {
	syncInstances.Do(func() {
		instances = []ping.Pingable{
			GetDbClient(),
		}
	})
	return instances
}
