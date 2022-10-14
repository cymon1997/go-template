package provider

import (
	"sync"

	ob "github.com/cymon1997/go-template/core/outbound/db"
	"github.com/cymon1997/go-template/pkg/db"
)

var (
	dbClient     db.Client
	syncDbClient sync.Once

	dbRepo     ob.Repository
	syncDbRepo sync.Once
)

func GetDbClient() db.Client {
	syncDbClient.Do(func() {
		cfg := GetMainConfig()
		dbClient = db.New(cfg.DB)
	})
	return dbClient
}

func GetDbRepository() ob.Repository {
	syncDbRepo.Do(func() {
		dbRepo = ob.New(GetDbClient())
	})
	return dbRepo
}
