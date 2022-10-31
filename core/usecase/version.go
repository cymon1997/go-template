package usecase

import (
	"context"

	"github.com/cymon1997/go-template/core/inbound/http/spec"
)

func (uc *usecaseImpl) Version(_ context.Context) (*spec.VersionResponse, error) {
	return &spec.VersionResponse{
		Version: "1.0.0",
	}, nil
}
