package main

import (
	"gofr.dev/pkg/gofr"
)

func main() {
	app := gofr.New()

	//app.EnableOAuth("http://localhost:9001/jwks", 5)

	app.GET("/", func(c *gofr.Context) (interface{}, error) {
		return "ok", nil
	})

	app.Run()
}
