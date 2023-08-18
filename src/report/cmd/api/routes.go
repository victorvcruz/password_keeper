package api

import (
	"fmt"
	"github.com/victorvcruz/password_warehouse/src/protobuf/report_pb"
	"google.golang.org/grpc"
	"net"
	"os"
	"report.com/cmd/api/handlers"
	"report.com/internal/auth"
	"report.com/pkg/middleware"
)

func New(report *handlers.ReportHandler, auth auth.AuthServiceClient) error {

	listener, err := net.Listen("tcp", fmt.Sprintf(":%s", os.Getenv("API_PORT")))
	if err != nil {
		return err
	}

	app := grpc.NewServer(
		grpc.UnaryInterceptor(middleware.NewInterceptor(auth).ServerInterceptor),
	)
	report_pb.RegisterReportServer(app, report)

	err = app.Serve(listener)
	if err != nil {
		return err
	}
	return nil
}
