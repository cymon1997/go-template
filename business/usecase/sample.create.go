package usecase

import (
	"context"
	"log"
	"time"

	"github.com/cymon1997/go-template/business/inbound/http/spec"
	"github.com/cymon1997/go-template/business/outbound/db/model"
	"github.com/google/uuid"
	"gopkg.in/guregu/null.v4"
)

func (uc *usecaseImpl) CreateSample(ctx context.Context, input spec.CreateSampleRequest) (*spec.CreateSampleResponse, error) {
	output, err := uc.dbRepo.InsertSample(ctx, model.Sample{
		UUID:       uuid.New().String(),
		Data:       null.StringFrom(input.Sample.Data),
		CreateTime: null.TimeFrom(time.Now().UTC()),
		// TODO: Add request source to context data
		//CreateBy:   null.String{},
	})
	if err != nil {
		log.Println("error execute dbRepo.InsertSample: ", err)
		return nil, err
	}
	return &spec.CreateSampleResponse{
		Sample: output,
	}, nil
}
