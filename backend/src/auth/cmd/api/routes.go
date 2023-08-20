package api

import (
	"auth.com/cmd/api/handlers"
	"fmt"
	"github.com/victorvcruz/password_warehouse/src/protobuf/auth_pb"
	"google.golang.org/grpc"
	"net"
	"os"
)

func New(auth *handlers.AuthHandler) error {

	listener, err := net.Listen("tcp", fmt.Sprintf(":%s", os.Getenv("API_PORT")))
	if err != nil {
		return err
	}

	app := grpc.NewServer()
	auth_pb.RegisterAuthServer(app, auth)

	err = app.Serve(listener)
	if err != nil {
		return err
	}
	return nil
}
