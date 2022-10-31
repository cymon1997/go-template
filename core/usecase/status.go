package usecase

import (
	"context"

	"github.com/cymon1997/go-template/core/inbound/http/spec"
)

func (uc *usecaseImpl) Status(ctx context.Context) (*spec.StatusResponse, error) {
	instances := make([]spec.Instance, 0)
	for _, instance := range uc.instances {
		status := spec.StatusOK
		if err := instance.Ping(ctx); err != nil {
			status = spec.StatusError
		}
		instances = append(instances, spec.Instance{
			Name:   instance.Name(),
			Status: status,
		})
	}
	return &spec.StatusResponse{
		Instances: instances,
	}, nil
}
