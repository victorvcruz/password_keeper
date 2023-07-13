package api

import (
	"fmt"
	"google.golang.org/grpc"
	"net"
	"os"
	"report.com/cmd/api/handlers"
	"report.com/internal/auth"
	"report.com/pkg/middleware"
	v1 "report.com/pkg/pb"
)

func New(user *handlers.ReportHandler, auth auth.AuthServiceClient) error {

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
