package controller

import (
	"context"

	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
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
	return nil, status.Error(codes.FailedPrecondition, "you are not allowed to greet")
}
