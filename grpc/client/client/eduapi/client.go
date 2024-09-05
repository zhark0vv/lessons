package eduapi

import (
	"context"

	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"
	educationapi "lessons/grpc/server/gen/go/zhark0vv/grpc/education/api"
)

type Client struct {
	api educationapi.EducationAPIClient
}

func InitClient() (*Client, error) {
	cc, err := grpc.NewClient(
		"localhost:8085",
		grpc.WithTransportCredentials(insecure.NewCredentials()),
	)
	if err != nil {
		return nil, err
	}
	return &Client{api: educationapi.NewEducationAPIClient(cc)}, nil
}

func (c Client) Greet(ctx context.Context, name string) (string, error) {
	res, err := c.api.Greet(
		ctx,
		&educationapi.GreetRequest{
			Name: name,
		},
	)
	if err != nil {
		return "", err
	}
	return res.Message, nil
}
