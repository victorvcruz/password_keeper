package auth

import (
	"context"
	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"
	"log"
	"os"
	"user.com/internal/utils/errors"
	pb2 "user.com/pkg/pb"
)

type AuthServiceClient interface {
	AuthToken(acessToken string) (string, error)
}

type authService struct {
	client pb2.AuthClient
	ctx    context.Context
}

func NewAuthService() AuthServiceClient {

	var opts []grpc.DialOption
	opts = append(opts, grpc.WithTransportCredentials(insecure.NewCredentials()))

	conn, err := grpc.Dial(os.Getenv("AUTH_HOST"), opts...)
	if err != nil {
		log.Fatal(err)
	}
	_client := pb2.NewAuthClient(conn)
	return &authService{
		client: _client,
		ctx:    context.Background(),
	}
}

func (a authService) AuthToken(acessToken string) (string, error) {
	auth, err := a.client.AuthToken(a.ctx, &pb2.AuthTokenRequest{AcessToken: acessToken})
	if err != nil {
		return "", err
	}

	if !auth.Authorize {
		return "", &errors.UnauthorizedTokenError{}
	}

	return auth.Id, nil
}
