package api

import (
	"fmt"
	"github.com/victorvcruz/password_warehouse/src/protobuf/user_pb"
	"google.golang.org/grpc"
	"net"
	"os"
	"user.com/cmd/api/handlers"
	"user.com/internal/auth"
	"user.com/pkg/middleware"
)

func New(user *handlers.UserHandler, auth auth.AuthServiceClient) error {

	listener, err := net.Listen("tcp", fmt.Sprintf(":%s", os.Getenv("API_PORT")))
	if err != nil {
		return err
	}

	app := grpc.NewServer(
		grpc.UnaryInterceptor(middleware.NewInterceptor(auth).ServerInterceptor),
	)
	user_pb.RegisterUserServer(app, user)

	err = app.Serve(listener)
	if err != nil {
		return err
	}
	return nil
}
