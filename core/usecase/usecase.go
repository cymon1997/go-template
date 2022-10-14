package usecase

import (
	"context"

	"github.com/cymon1997/go-template/core/inbound/http/spec"
	"github.com/cymon1997/go-template/core/outbound/db"
	"github.com/cymon1997/go-template/internal/ping"
)

type Usecase interface {
	Status(ctx context.Context) (*spec.StatusResponse, error)
	Version(ctx context.Context) (*spec.VersionResponse, error)
}

type usecaseImpl struct {
	dbRepo    db.Repository
	instances []ping.Pingable
}

func New(dbRepo db.Repository, instances []ping.Pingable) Usecase {
	return &usecaseImpl{
		dbRepo:    dbRepo,
		instances: instances,
	}
}
