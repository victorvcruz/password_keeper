package authorization

import (
	"context"
	"github.com/victorvcruz/password_warehouse/src/protobuf/auth_pb"
	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"
	"log"
	"os"
)

type AuthorizationClient interface {
	Login(service string) (string, error)
}

type authService struct {
	ctx         context.Context
	apiToken    string
	serviceName string
	host        string
}

func NewAuthorization(_service string) AuthorizationClient {
	auth := authService{
		host:        os.Getenv("AUTH_HOST"),
		serviceName: _service,
		ctx:         context.Background(),
	}
	return &auth
}

func (a *authService) client() auth_pb.AuthClient {
	var opts []grpc.DialOption
	opts = append(opts, grpc.WithTransportCredentials(insecure.NewCredentials()))

	conn, err := grpc.Dial(a.host, opts...)
	if err != nil {
		log.Fatal(err)
	}

	return auth_pb.NewAuthClient(conn)
}

func (a *authService) Login(service string) (string, error) {
	auth, err := a.client().LoginApi(a.ctx, &auth_pb.LoginService{Service: a.serviceName, ServiceConn: service, ApiToken: a.apiToken})
	if err != nil {
		return "", err
	}
	return auth.AcessToken, nil
}
