package handlers

import (
	"auth.com/internal/auth"
	"auth.com/internal/utils/errors"
	proto "auth.com/pkg/proto/v1"
	"context"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
)

type AuthHandler struct {
	proto.UnimplementedLoginServer
	service auth.AuthServiceClient
}

func NewAuthHandler(_authService auth.AuthServiceClient) *AuthHandler {
	return &AuthHandler{
		service: _authService,
	}
}

func (a *AuthHandler) Login(ctx context.Context, req *proto.LoginRequest) (*proto.LoginResponse, error) {

	login := auth.LoginRequest{Email: req.Email, Password: req.Password}

	token, err := a.service.Login(login)
	if err != nil {
		switch err.(type) {
		case *errors.NotFoundEmailError:
			return nil, status.Error(codes.InvalidArgument, "Not found email")
		case *errors.UnauthorizedPasswordError:
			return nil, status.Error(codes.InvalidArgument, "Unauthorized password")
		default:
			return nil, status.Error(codes.Internal, "Internal server error")
		}
	}

	return &proto.LoginResponse{AcessToken: token}, nil
}
