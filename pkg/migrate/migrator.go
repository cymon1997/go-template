package migrate

import (
	"log"

	"github.com/golang-migrate/migrate"
	_ "github.com/golang-migrate/migrate/database/mysql"
	_ "github.com/golang-migrate/migrate/database/postgres"
	_ "github.com/golang-migrate/migrate/source/file"
)

type Migrator interface {
	Up() error
	Down() error
}

type migratorImpl struct {
	migrator *migrate.Migrate
}

func New(cfg Config) Migrator {
	schemaTableName := cfg.SchemaTableName
	if schemaTableName == "" {
		schemaTableName = DefaultSchemaTableName
	}
	dsn, err := addSchemaTableName(cfg.DSN, cfg.SchemaTableName)
	if err != nil {
		log.Fatal("error parse migrator database dsn: ", err)
	}
	migrator, err := migrate.New(cfg.FilePath, dsn)
	if err != nil {
		log.Fatal("error init migrator: ", err)
	}
	return &migratorImpl{
		migrator: migrator,
	}
}

func (m *migratorImpl) Up() error {
	if err := m.migrator.Up(); err != migrate.ErrNoChange {
		return err
	}
	return nil
}

func (m *migratorImpl) Down() error {
	if err := m.migrator.Down(); err != migrate.ErrNoChange {
		return err
	}
	return nil
}
