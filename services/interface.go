package services

import (
	"developer.zopsmart.com/go/gofr/pkg/gofr"
	"github.int.mcafee.com/mcafee/bulk-processing-framework/pkg/models"
)

type TransformerService interface {
	CustomTransformer(ctx *gofr.Context, record models.Record) (*models.Record, error)
}
