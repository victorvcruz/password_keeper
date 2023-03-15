package handlers

import (
	"auth.com/internal/auth"
	"auth.com/internal/utils/errors"
	pb2 "auth.com/pkg/pb"
	"context"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
)

type AuthHandler struct {
	pb2.UnimplementedAuthServer
	service auth.AuthServiceClient
}

func NewAuthHandler(_authService auth.AuthServiceClient) *AuthHandler {
	return &AuthHandler{
		service: _authService,
	}
}

func (a *AuthHandler) Login(ctx context.Context, req *pb2.LoginRequest) (*pb2.LoginResponse, error) {

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

	return &pb2.LoginResponse{AcessToken: token}, nil
}

func (a *AuthHandler) AuthToken(ctx context.Context, req *pb2.AuthTokenRequest) (*pb2.AuthTokenResponse, error) {

	login := auth.AuthTokenRequest{AcessToken: req.AcessToken}

	id, err := a.service.ValidateToken(login)
	if err != nil {
		switch err.(type) {
		case *errors.NotFoundIdError:
			return nil, status.Error(codes.InvalidArgument, "Not found user id")
		case *errors.ExpiredTokenError:
			return nil, status.Error(codes.Unauthenticated, "expired token")
		case *errors.InvalidTokenError:
			return nil, status.Error(codes.Unauthenticated, "invalid token")
		default:
			return nil, status.Error(codes.Internal, "Internal server error")
		}
	}

	valid := id != 0
	return &pb2.AuthTokenResponse{Id: id, Authorize: valid}, nil
}
