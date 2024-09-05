package main

import (
	"context"

	"lessons/grpc/server/app"
	"lessons/grpc/server/controller"
)

func main() {
	c := controller.New()
	a, err := app.Init(context.Background(), c)

	if err != nil {
		panic(err)
	}

	if err := a.Run(); err != nil {
		panic(err)
	}
}
