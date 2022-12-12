package api

import (
	"auth.com/cmd/api/handlers"
	v1 "auth.com/pkg/proto/v1"
	"fmt"
	"google.golang.org/grpc"
	"net"
	"os"
)

func New(user *handlers.AuthHandler) error {

	listener, err := net.Listen("tcp", fmt.Sprintf(":%s", os.Getenv("API_PORT")))
	if err != nil {
		return err
	}

	app := grpc.NewServer()
	v1.RegisterLoginServer(app, user)

	err = app.Serve(listener)
	if err != nil {
		return err
	}
	return nil
}
