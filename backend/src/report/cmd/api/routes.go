package api

import (
	"fmt"
	"github.com/victorvcruz/password_warehouse/src/protobuf/report_pb"
	"google.golang.org/grpc"
	"net"
	"os"
	"report.com/cmd/api/handlers"
)

func New(report *handlers.ReportHandler) error {

	listener, err := net.Listen("tcp", fmt.Sprintf(":%s", os.Getenv("API_PORT")))
	if err != nil {
		return err
	}

	app := grpc.NewServer()
	report_pb.RegisterReportServer(app, report)

	err = app.Serve(listener)
	if err != nil {
		return err
	}
	return nil
}
