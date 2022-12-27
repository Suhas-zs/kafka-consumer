package main

import (
	"developer.zopsmart.com/go/gofr/pkg/gofr"

	"kafka-consumer/handlers"
	"kafka-consumer/services/transformers"
)

func main() {
	app := gofr.NewCMD()

	t := transformers.New()
	h := handlers.New(t)

	app.GET("load", h.Load)

	app.Start()
}
