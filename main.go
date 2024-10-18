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

	fmt.Println(string(op), err)

	app.GET("/", func(c *gofr.Context) (interface{}, error) {
		op, err := exec.Command("docker", "ps").CombinedOutput()

		return string(op), err
	})

	app.Run()
}
