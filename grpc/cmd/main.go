package main

import (
	"context"
	"fmt"

	"lessons/grpc/client/client/eduapi"
)

func main() {
	ctx := context.Background()

	c, err := eduapi.InitClient()
	if err != nil {
		panic(err)
	}

	resp, err := c.Greet(ctx, "John")
	if err != nil {
		panic(err)
	}

	fmt.Println(resp)
}
