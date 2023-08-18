package handlers

import (
	"context"
	"github.com/go-playground/validator/v10"
	"github.com/victorvcruz/password_warehouse/src/protobuf/user_pb"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
	"google.golang.org/protobuf/types/known/timestamppb"
	"user.com/internal/auth"
	"user.com/internal/user"
	"user.com/internal/utils"
	"user.com/internal/utils/errors"
)

type UserHandler struct {
	user_pb.UnimplementedUserServer
	userService user.UserServiceClient
	authService auth.AuthServiceClient
	validate    *validator.Validate
}

func NewUserHandler(
	_userService user.UserServiceClient,
	_authService auth.AuthServiceClient,
	_validate *validator.Validate,
) *UserHandler {
	return &UserHandler{
		userService: _userService,
		authService: _authService,
		validate:    _validate,
	}
}

func (u *UserHandler) CreateUser(_ context.Context, req *user_pb.UserRequest) (*user_pb.UserResponse, error) {

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

	return &user_pb.UserResponse{Name: req.Name, Email: req.Email}, nil
}

func (u *UserHandler) FindUser(ctx context.Context, _ *user_pb.Empty) (*user_pb.UserResponse, error) {

	token, err := utils.BearerToken(ctx)
	if err != nil {
		return nil, status.Error(codes.InvalidArgument, "Add token")
	}

	id, err := u.authService.AuthUserToken(token)
	if err != nil {
		return nil, status.Error(codes.InvalidArgument, err.Error())
	}

	user, err := u.userService.FindUserById(id)
	if err != nil {
		return nil, status.Error(codes.Internal, "Internal server error")
	}

	return &user_pb.UserResponse{Name: user.Name, Email: user.Email}, nil
}

func (u *UserHandler) UpdateUser(ctx context.Context, req *user_pb.UserRequest) (*user_pb.UserResponse, error) {

	token, err := utils.BearerToken(ctx)
	if err != nil {
		return nil, status.Error(codes.InvalidArgument, "Add token")
	}

	id, err := u.authService.AuthUserToken(token)
	if err != nil {
		return nil, status.Error(codes.InvalidArgument, err.Error())
	}

	user := user.UserRequest{Name: req.Name, Email: req.Email, MasterPassword: req.MasterPassword}
	err = u.validate.Struct(user)
	if err != nil {
		return nil, status.Error(codes.InvalidArgument, utils.RequestUserValidate(err))
	}

	err = u.userService.UpdateUser(id, user)
	if err != nil {
		return nil, status.Error(codes.Internal, "Internal server error")
	}

	return &user_pb.UserResponse{Name: req.Name, Email: req.Email}, nil
}

func (u *UserHandler) DeleteUser(ctx context.Context, _ *user_pb.Empty) (*user_pb.MessageResponse, error) {

	token, err := utils.BearerToken(ctx)
	if err != nil {
		return nil, status.Error(codes.InvalidArgument, "Add token")
	}

	id, err := u.authService.AuthUserToken(token)
	if err != nil {
		return nil, status.Error(codes.InvalidArgument, err.Error())
	}

	err = u.userService.DeleteUser(id)
	if err != nil {
		return nil, status.Error(codes.Internal, "Internal server error")
	}

	return &user_pb.MessageResponse{Message: "User deleted"}, nil
}

func (u *UserHandler) FindUserByData(ctx context.Context, _ *user_pb.Empty) (*user_pb.DetailedUserResponse, error) {

	token, err := utils.BearerToken(ctx)
	if err != nil {
		return nil, status.Error(codes.InvalidArgument, "Add token")
	}

	err = u.authService.AuthServiceToken(token)
	if err != nil {
		return nil, status.Error(codes.InvalidArgument, err.Error())
	}

	user, err := u.userService.FindUserByData(utils.GetMetadata(ctx, "id"), utils.GetMetadata(ctx, "userName"), utils.GetMetadata(ctx, "userEmail"))
	if err != nil {
		return nil, status.Error(codes.Internal, "Internal server error")
	}

	return &user_pb.DetailedUserResponse{
		Id: user.Id, Name: user.Name, Email: user.Email, MasterPassword: user.MasterPassword,
		CreatedAt: timestamppb.New(user.CreatedAt), UpdatedAt: timestamppb.New(user.UpdatedAt),
	}, nil
}
