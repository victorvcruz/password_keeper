package api

import (
	"fmt"
	"google.golang.org/grpc"
	"internal.com/cmd/api/handlers"
	"internal.com/pkg/pb"
	"net"
	"os"
)

func New(internal *handlers.InternalHandler) error {

	listener, err := net.Listen("tcp", fmt.Sprintf(":%s", os.Getenv("API_PORT")))
	if err != nil {
		return err
	}

	app := grpc.NewServer()
	pb.RegisterInternalServer(app, internal)

	err = app.Serve(listener)
	if err != nil {
		return err
	}
	return nil
}
