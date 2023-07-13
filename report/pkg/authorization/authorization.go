package authorization

import (
	"context"
	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"
	"log"
	"os"
	pb2 "report.com/pkg/pb"
)

var (
	serviceName string
	client      pb2.AuthClient
	apiToken    string
)

func Setup(service string) {
	serviceName = service
	client = connectClient()
	apiToken = registerService()
}

func Login(service string) (string, error) {
	auth, err := client.LoginApi(context.Background(), &pb2.LoginService{Service: serviceName, ServiceConn: service, ApiToken: apiToken})
	if err != nil {
		return "", err
	}
	return auth.AcessToken, err
}

func connectClient() pb2.AuthClient {
	var opts []grpc.DialOption
	opts = append(opts, grpc.WithTransportCredentials(insecure.NewCredentials()))

	conn, err := grpc.Dial(os.Getenv("AUTH_HOST"), opts...)
	if err != nil {
		log.Fatal(err)
	}
	return pb2.NewAuthClient(conn)
}

func registerService() string {
	auth, err := client.RegisterService(context.Background(), &pb2.Register{Service: serviceName})
	if err != nil {
		log.Fatalf("failed to register authorization service:%s", err.Error())
	}
	return auth.ApiToken
}
