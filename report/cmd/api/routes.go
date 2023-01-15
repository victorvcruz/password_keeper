package api

import (
	"fmt"
	"google.golang.org/grpc"
	"net"
	"os"
	"report.com/cmd/api/handlers"
	v1 "report.com/pkg/pb"
)

func New(user *handlers.ReportHandler) error {

	listener, err := net.Listen("tcp", fmt.Sprintf(":%s", os.Getenv("API_PORT")))
	if err != nil {
		return err
	}

	app := grpc.NewServer()
	v1.RegisterUserServer(app, user)

	err = app.Serve(listener)
	if err != nil {
		return err
	}
	return nil
}
