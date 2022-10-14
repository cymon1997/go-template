package db

import (
	"github.com/cymon1997/go-template/pkg/db"
)

type Repository interface {
}

type repositoryImpl struct {
	db db.Client
}

func New(db db.Client) Repository {
	return &repositoryImpl{
		db: db,
	}
}
