package api

import (
	"fmt"
	"github.com/victorvcruz/password_warehouse/protobuf/user_pb"
	"google.golang.org/grpc"
	"net"
	"os"
	"user.com/cmd/api/handlers"
)

func New(user *handlers.UserHandler) error {

	listener, err := net.Listen("tcp", fmt.Sprintf(":%s", os.Getenv("API_PORT")))
	if err != nil {
		return err
	}

	app := grpc.NewServer()
	user_pb.RegisterUserServer(app, user)

	err = app.Serve(listener)
	if err != nil {
		return err
	}
	return nil
}
