package handlers

import (
	"context"
	"github.com/go-playground/validator/v10"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
	"user.com/internal/user"
	"user.com/internal/utils"
	"user.com/internal/utils/errors"
	proto "user.com/pkg/proto/v1"
)

type UserHandler struct {
	proto.UnimplementedUserServer
	userService user.UserServiceClient
	validate    *validator.Validate
}

func NewUserHandler(_userService user.UserServiceClient, _validate *validator.Validate) *UserHandler {
	return &UserHandler{
		userService: _userService,
		validate:    _validate,
	}
}

func (u *UserHandler) CreateUser(ctx context.Context, req *proto.UserRequest) (*proto.UserResponse, error) {

	user := user.UserRequest{Name: req.Name, Email: req.Email, MasterPassword: req.MasterPassword}

	err := u.validate.Struct(user)
	if err != nil {
		return nil, status.Error(codes.InvalidArgument, utils.RequestUserValidate(err))
	}

	err = u.userService.CreateUser(user)
	if err != nil {
		switch err.(type) {
		case *errors.ConflictEmailError:
			return nil, status.Error(codes.AlreadyExists, "Email already exist")
		default:
			return nil, status.Error(codes.Internal, "Internal server error")
		}
	}

	return &proto.UserResponse{Name: req.Name, Email: req.Email}, nil
}
