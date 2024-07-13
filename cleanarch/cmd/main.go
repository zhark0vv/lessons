package main

import (
	"lessons/cleanarch/internal/app"
)

func main() {
	a := app.New(app.Config{
		InMemory: true,
	})

	a.Run(":8085")
}
