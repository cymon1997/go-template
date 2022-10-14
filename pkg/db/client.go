package db

import (
	"context"
	"log"
	"time"

	"github.com/cymon1997/go-template/internal/ping"
	_ "github.com/go-sql-driver/mysql"
	_ "github.com/lib/pq"
	"gorm.io/gorm"
)

type Client interface {
	Instance() *gorm.DB
	ping.Pingable
}

type clientImpl struct {
	*gorm.DB
}

func New(cfg Config) Client {
	client, err := gorm.Open(getDialector(cfg), &gorm.Config{
		SkipDefaultTransaction:   false,
		NowFunc:                  getNowFunc(cfg),
		PrepareStmt:              true,
		DisableAutomaticPing:     false,
		DisableNestedTransaction: false,
	})
	if err != nil {
		log.Fatal("error init database")
	}
	db, err := client.DB()
	if err != nil {
		log.Fatal("error connect database")
	}
	db.SetMaxOpenConns(cfg.MaxOpenConn)
	db.SetMaxIdleConns(cfg.MaxIdleConn)
	db.SetConnMaxIdleTime(time.Duration(cfg.ConnMaxIdleTime) * time.Second)
	db.SetConnMaxLifetime(time.Duration(cfg.ConnMaxLifeTime) * time.Second)
	return &clientImpl{
		DB: client,
	}
}

func (c *clientImpl) Instance() *gorm.DB {
	return c.DB
}

func (c *clientImpl) Ping(_ context.Context) error {
	db, err := c.DB.DB()
	if err != nil {
		return err
	}
	return db.Ping()
}
