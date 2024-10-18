package main

import (
	"fmt"
	"gofr.dev/pkg/gofr"
	"os/exec"
)

func main() {
	app := gofr.New()

	//app.EnableOAuth("http://localhost:9001/jwks", 5)

	op, err := exec.Command("docker", "ps").CombinedOutput()
	if err != nil {
		panic(err)
	}

	fmt.Println(string(op))

	app.GET("/", func(c *gofr.Context) (interface{}, error) {
		return "ok", nil
	})

	app.Run()
}
