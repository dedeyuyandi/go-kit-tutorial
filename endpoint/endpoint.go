package endpoint

import (
	"context"

	"github.com/go-kit/kit/endpoint"
	"github.com/testing/go-kit-tutorial/model"
	"github.com/testing/go-kit-tutorial/service"
)

type Endpoints struct {
	CreateUser endpoint.Endpoint
	GetUser    endpoint.Endpoint
}

func MakeEndpoints(s service.Service) Endpoints {
	return Endpoints{
		CreateUser: makeCreateUserEndpoint(s),
		GetUser:    makeGetUserEndpoint(s),
	}
}

func makeCreateUserEndpoint(s service.Service) endpoint.Endpoint {
	return func(ctx context.Context, request interface{}) (interface{}, error) {
		req := request.(model.WriteUserRequest)
		err := s.CreateUser(ctx, req)
		if err != nil {
			return nil, err
		}
		return model.CreateUserResponse{Message: "ok"}, nil
	}
}

func makeGetUserEndpoint(s service.Service) endpoint.Endpoint {
	return func(ctx context.Context, request interface{}) (interface{}, error) {
		req := request.(model.GetUserRequest)
		resp, err := s.GetUser(ctx, req.Id)
		return resp, err
	}
}
