package user

import (
	"context"
	"github.com/victorvcruz/password_warehouse/src/protobuf/user_pb"
	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"
	"google.golang.org/grpc/metadata"
	"log"
	"os"
	"strconv"
)

type ServiceClient interface {
	UserById(id int64) (*UserDTO, error)
	UserByEmail(email string) (*UserDTO, error)
}

type service struct {
	client user_pb.UserClient
	ctx    context.Context
}

func NewUserService() ServiceClient {

	var opts []grpc.DialOption
	opts = append(opts, grpc.WithTransportCredentials(insecure.NewCredentials()))

	conn, err := grpc.Dial(os.Getenv("USER_HOST"), opts...)
	if err != nil {
		log.Fatal(err)
	}
	_client := user_pb.NewUserClient(conn)
	return &service{
		client: _client,
		ctx:    context.Background(),
	}
}

func (u *service) UserByEmail(email string) (*UserDTO, error) {
	md := metadata.New(map[string]string{"userEmail": email})
	ctx := metadata.NewOutgoingContext(context.Background(), md)
	user, err := u.client.FindUserByData(ctx, &user_pb.Empty{})
	if err != nil {
		return nil, err
	}

	return u.userResponseToDto(user), nil
}

func (u *service) UserById(id int64) (*UserDTO, error) {
	md := metadata.New(map[string]string{"id": strconv.FormatInt(id, 10)})
	ctx := metadata.NewOutgoingContext(context.Background(), md)
	user, err := u.client.FindUserByData(ctx, &user_pb.Empty{})
	if err != nil {
		return nil, err
	}

	return u.userResponseToDto(user), nil
}

func (u *service) userResponseToDto(response *user_pb.DetailedUserResponse) *UserDTO {
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
