package spec

import "github.com/cymon1997/go-template/business/outbound/db/model"

type CreateSampleRequest struct {
	Sample struct {
		Data string `json:"data"`
	} `json:"sample"`
}

type CreateSampleResponse struct {
	Sample *model.Sample `json:"sample"`
}
