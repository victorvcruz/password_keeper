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
	AuthUserToken(acessToken string) (int64, error)
	AuthServiceToken(acessToken string) error
	Login(service string) (string, error)
}

type authService struct {
	client      pb2.AuthClient
	ctx         context.Context
	apiToken    string
	serviceName string
}

func NewAuthService(_service string) AuthServiceClient {

	var opts []grpc.DialOption
	opts = append(opts, grpc.WithTransportCredentials(insecure.NewCredentials()))

	conn, err := grpc.Dial(os.Getenv("AUTH_HOST"), opts...)
	if err != nil {
		log.Fatal(err)
	}

	auth := authService{
		serviceName: _service,
		ctx:         context.Background(),
	}

	auth.client = pb2.NewAuthClient(conn)
	auth.apiToken = auth.registerService()

	return &auth
}

func (a *authService) AuthUserToken(acessToken string) (int64, error) {
	auth, err := a.client.AuthToken(a.ctx, &pb2.AuthTokenRequest{AcessToken: acessToken})
	if err != nil {
		return 0, err
	}

	if !auth.Authorize {
		return 0, &errors.UnauthorizedTokenError{}
	}

	return auth.Id, nil
}

func (a *authService) AuthServiceToken(acessToken string) error {
	auth, err := a.client.AuthApi(a.ctx, &pb2.AuthTokenService{AcessToken: acessToken})
	if err != nil {
		return err
	}

	if !auth.Authorize {
		return &errors.UnauthorizedTokenError{}
	}

	return nil
}

func (a *authService) Login(service string) (string, error) {
	auth, err := a.client.LoginApi(a.ctx, &pb2.LoginService{Service: a.serviceName, ServiceConn: service, ApiToken: a.apiToken})
	if err != nil {
		return "", err
	}
	return auth.AcessToken, nil
}

func (a *authService) registerService() string {
	auth, err := a.client.RegisterService(a.ctx, &pb2.Register{Service: a.serviceName})
	if err != nil {
		log.Fatalf("failed to register authorization service:%s", err.Error())
	}
	return auth.ApiToken
}
