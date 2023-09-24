package db

import (
	"context"
	"log"

	"github.com/cymon1997/go-template/business/outbound/db/model"
	"github.com/cymon1997/go-template/internal/errors"
)

func (r *repositoryImpl) InsertSample(ctx context.Context, data model.Sample) (*model.Sample, error) {
	if err := r.db.Instance().WithContext(ctx).Create(&data).Error; err != nil {
		log.Print("error insert: ", err)
		return nil, errors.ErrDatabaseOp
	}
	return &data, nil
}
