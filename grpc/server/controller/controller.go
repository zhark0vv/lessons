package controller

import (
	"context"

	educationapi "lessons/grpc/server/gen/go/zhark0vv/grpc/education/api"
)

var _ educationapi.EducationAPIServer = (*Controller)(nil)

type Controller struct {
	educationapi.UnimplementedEducationAPIServer
}

func New() *Controller {
	return &Controller{}
}

func (c *Controller) Greet(_ context.Context,
	req *educationapi.GreetRequest) (*educationapi.GreetResponse, error) {
	return &educationapi.GreetResponse{Message: "Hello, " + req.Name}, nil
}
