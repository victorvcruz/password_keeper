package authorization

import (
	"context"
	"github.com/victorvcruz/password_warehouse/protobuf/auth_pb"
	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"
	"log"
	"os"
)

var (
	serviceName string
	client      auth_pb.AuthClient
	apiToken    string
)

func Setup(service string) {
	serviceName = service
	client = connectClient()
	apiToken = registerService()
}

func Login(service string) (string, error) {
	auth, err := client.LoginApi(context.Background(), &auth_pb.LoginService{Service: serviceName, ServiceConn: service, ApiToken: apiToken})
	if err != nil {
		return "", err
	}
	return auth.AcessToken, err
}

func connectClient() auth_pb.AuthClient {
	var opts []grpc.DialOption
	opts = append(opts, grpc.WithTransportCredentials(insecure.NewCredentials()))

	conn, err := grpc.Dial(os.Getenv("AUTH_HOST"), opts...)
	if err != nil {
		log.Fatal(err)
	}
	return auth_pb.NewAuthClient(conn)
}

func registerService() string {
	auth, err := client.RegisterService(context.Background(), &auth_pb.Register{Service: serviceName})
	if err != nil {
		log.Fatalf("failed to register authorization service:%s", err.Error())
	}
	return auth.ApiToken
}
