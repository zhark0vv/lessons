package main

import (
	"context"
	"fmt"
)

const key = "userID"

func main() {
	ctx := context.WithValue(context.Background(), key, 12345)
	if userID, ok := ctx.Value(key).(string); ok {
		fmt.Println("User ID:", userID)
	}
}
