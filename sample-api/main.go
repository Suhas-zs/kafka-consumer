package main

import (
	"developer.zopsmart.com/go/gofr/pkg/gofr"

	"kafka-consumer/sample-api/handlers"
)

func main() {
	app := gofr.New()

	app.POST("/record", handlers.PostHandler)

	app.GET("/hello", handlers.GetHandler)

	app.Start()
}
