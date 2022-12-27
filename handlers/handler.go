package handlers

import (
	"developer.zopsmart.com/go/gofr/pkg/gofr"
	bulkLoader "github.int.mcafee.com/mcafee/bulk-processing-framework/pkg/loader"

	"kafka-consumer/services"
)

const (
	Layer1 = 1
	Layer2 = 2
)

type loader struct {
	transformer services.TransformerService
}

func New(t services.TransformerService) loader {
	return loader{transformer: t}
}

func (l loader) Load(ctx *gofr.Context) (interface{}, error) {
	b, err := bulkLoader.New(ctx, ctx.PathParam("config"))
	if err != nil {
		return nil, err
	}

	id, err := b.AddConfigLayerWithOptions(l.transformer.CustomTransformer, 0, Layer2, Layer1)
	if err != nil {
		ctx.Logger.Error("error in transformation at layer %s", id)
	}

	if err := b.Start(ctx); err != nil {
		return nil, err
	}

	return "Done!", nil
}
