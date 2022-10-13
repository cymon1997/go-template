package usecase

import (
	"context"

	"github.com/cymon1997/go-template/core/inbound/http/spec"
)

type Usecase interface {
	Version(ctx context.Context) (*spec.VersionResponse, error)
}

type usecaseImpl struct {
}

func New() Usecase {
	return &usecaseImpl{}
}
