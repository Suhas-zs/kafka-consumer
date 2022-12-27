package handlers

import (
	"developer.zopsmart.com/go/gofr/pkg/gofr"

	"kafka-consumer/sample-api/models"
)

func PostHandler(ctx *gofr.Context) (interface{}, error) {
	var per models.Person

	if err := ctx.Bind(&per); err != nil {
		return nil, err
	}

	return per, nil
}

func GetHandler(ctx *gofr.Context) (interface{}, error) {
	name := ctx.Param("name")

	return "Hello " + name, nil
}
