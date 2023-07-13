package api

import (
	"fmt"
	"google.golang.org/grpc"
	"net"
	"os"
	"user.com/cmd/api/handlers"
	"user.com/internal/auth"
	"user.com/pkg/middleware"
	v1 "user.com/pkg/pb"
)

func New(user *handlers.UserHandler, auth auth.AuthServiceClient) error {

	listener, err := net.Listen("tcp", fmt.Sprintf(":%s", os.Getenv("API_PORT")))
	if err != nil {
		return err
	}

	app := grpc.NewServer(
		grpc.UnaryInterceptor(middleware.NewInterceptor(auth).ServerInterceptor),
	)
	v1.RegisterUserServer(app, user)

	err = app.Serve(listener)
	if err != nil {
		return err
	}
	return nil
}
