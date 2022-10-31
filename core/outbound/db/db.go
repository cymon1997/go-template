package db

import (
	"context"

	"github.com/cymon1997/go-template/core/outbound/db/model"
	"github.com/cymon1997/go-template/pkg/db"
)

type Repository interface {
	InsertSample(ctx context.Context, data model.Sample) (*model.Sample, error)
}

type repositoryImpl struct {
	db db.Client
}

func New(db db.Client) Repository {
	return &repositoryImpl{
		db: db,
	}
}
