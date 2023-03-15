package handlers

import (
	"context"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
	"internal.com/internal/service"
	pb2 "internal.com/pkg/pb"
)

type InternalHandler struct {
	pb2.UnimplementedInternalServer
	service service.InternalServiceClient
}

func NewInternalHandler(_service service.InternalServiceClient) *InternalHandler {
	return &InternalHandler{
		service: _service,
	}
}

func (a *InternalHandler) TokenService(ctx context.Context, req *pb2.TokenRequest) (*pb2.TokenResponse, error) {

	login := service.TokenRequest{Service: req.Service, Password: req.Password}

	token, err := a.service.TokenRequest(login)
	if err != nil {
		switch err.(type) {
		default:
			return nil, status.Error(codes.Internal, "Internal server error")
		}
	}

	return &pb2.TokenResponse{AcessToken: token}, nil
}
