package main

import (
	"fmt"
	"log"

	"github.com/cymon1997/go-template/provider"
)

func main() {
	cfg := provider.GetMainConfig()

	err := provider.GetMigrator().Up()
	if err != nil {
		log.Fatal("error apply db migrations: ", err)
	}

	err = provider.GetSeeder().Up()
	if err != nil {
		log.Fatal("error apply db seeders: ", err)
	}

	log.Printf("serve http at %s:%d\n", cfg.App.Host, cfg.App.Port)
	err = provider.GetHttpServer().Run(fmt.Sprintf(":%d", cfg.App.Port))
	if err != nil {
		log.Fatal("error serve http: ", err)
	}
}
