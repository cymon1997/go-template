package provider

import (
	"sync"

	"github.com/cymon1997/go-template/pkg/migrate"
)

var (
	migrator     migrate.Migrator
	syncMigrator sync.Once

	seeder     migrate.Migrator
	syncSeeder sync.Once
)

func GetMigrator() migrate.Migrator {
	syncMigrator.Do(func() {
		cfg := GetMainConfig()
		migrator = migrate.New(cfg.Migrator)
	})
	return migrator
}

func GetSeeder() migrate.Migrator {
	syncSeeder.Do(func() {
		cfg := GetMainConfig()
		seeder = migrate.New(cfg.Seeder)
	})
	return seeder
}
