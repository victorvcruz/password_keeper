package handlers

import (
	"auth.com/cmd/api/model"
	"auth.com/internal/auth"
	"auth.com/internal/utils/errors"
	"context"
	"github.com/victorvcruz/password_warehouse/protobuf/auth_pb"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
)

type AuthHandler struct {
	auth_pb.UnimplementedAuthServer
	service auth.ServiceClient
}

func NewAuthHandler(_authService auth.ServiceClient) *AuthHandler {
	return &AuthHandler{
		service: _authService,
	}
}

func (a *AuthHandler) Login(_ context.Context, req *auth_pb.LoginRequest) (*auth_pb.LoginResponse, error) {
	token, err := a.service.Login(model.LoginFromProto(req))
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
	return &auth_pb.LoginResponse{AcessToken: token}, nil
}

func (a *AuthHandler) AuthToken(_ context.Context, req *auth_pb.AuthTokenRequest) (*auth_pb.AuthTokenResponse, error) {
	id, err := a.service.ValidateToken(model.AuthFromProto(req))
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
	return &auth_pb.AuthTokenResponse{Id: id, Authorize: id != 0}, nil
}
