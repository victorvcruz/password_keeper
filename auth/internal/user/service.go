package user

import (
	pb2 "auth.com/pkg/pb"
	"context"
	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"
	"google.golang.org/grpc/metadata"
	"log"
	"os"
	"strconv"
)

type UserServiceClient interface {
	UserById(id int64) (*UserDTO, error)
	UserByEmail(email string) (*UserDTO, error)
}

type userService struct {
	client pb2.UserClient
	ctx    context.Context
}

func NewUserService() UserServiceClient {

	var opts []grpc.DialOption
	opts = append(opts, grpc.WithTransportCredentials(insecure.NewCredentials()))

	conn, err := grpc.Dial(os.Getenv("USER_HOST"), opts...)
	if err != nil {
		log.Fatal(err)
	}
	_client := pb2.NewUserClient(conn)
	return &userService{
		client: _client,
		ctx:    context.Background(),
	}
}

func (u *userService) UserByEmail(email string) (*UserDTO, error) {
	md := metadata.New(map[string]string{"userEmail": email})
	ctx := metadata.NewOutgoingContext(context.Background(), md)
	user, err := u.client.FindUserByData(ctx, &pb2.Empty{})
	if err != nil {
		return nil, err
	}

	return u.userResponseToDto(user), nil
}

func (u *userService) UserById(id int64) (*UserDTO, error) {
	md := metadata.New(map[string]string{"id": strconv.FormatInt(id, 10)})
	ctx := metadata.NewOutgoingContext(context.Background(), md)
	user, err := u.client.FindUserByData(ctx, &pb2.Empty{})
	if err != nil {
		return nil, err
	}

	return u.userResponseToDto(user), nil
}

func (u *userService) userResponseToDto(response *pb2.DetailedUserResponse) *UserDTO {
	return &UserDTO{
		Id:             response.Id,
		Name:           response.Name,
		Email:          response.Email,
		MasterPassword: response.MasterPassword,
		CreatedAt:      response.CreatedAt.AsTime(),
		UpdatedAt:      response.UpdatedAt.AsTime(),
		DeletedAt:      response.DeletedAt.AsTime(),
	}
}
