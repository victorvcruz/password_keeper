package handlers

import (
	"context"
	"github.com/golang/protobuf/ptypes/empty"
	"github.com/victorvcruz/password_warehouse/protobuf/report_pb"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
	"report.com/internal/report"
	"report.com/internal/utils"
)

type ReportHandler struct {
	report_pb.UnimplementedReportServer
	reportService report.ServiceClient
}

func NewReportHandler(_reportService report.ServiceClient) *ReportHandler {
	return &ReportHandler{
		reportService: _reportService,
	}
}

func (r *ReportHandler) CreateReport(_ context.Context, req *report_pb.ReportRequest) (*report_pb.ReportResponse, error) {

	report := report.Request{
		Action:      req.Action,
		Description: req.Description,
		UserId:      req.UserId,
		VaultId:     req.VaultId,
	}

	err := r.reportService.CreateReport(report)
	if err != nil {
		return nil, status.Error(codes.Internal, err.Error())
	}

	return &report_pb.ReportResponse{Action: req.Action, Description: req.Description, UserId: req.UserId,
		VaultId: req.VaultId}, nil
}

func (r *ReportHandler) ReportByUserId(ctx context.Context, _ *empty.Empty) (*report_pb.ListReportResponse, error) {

	userId, err := utils.GetMetadataByKey(ctx, "userId")
	if err != nil {
		return nil, status.Error(codes.InvalidArgument, err.Error())
	}

	reports, err := r.reportService.FindReportsByUserId(userId)
	if err != nil {
		return nil, status.Error(codes.Internal, err.Error())
	}

	var listReport report_pb.ListReportResponse
	for _, report := range reports {
		listReport.Reports = append(listReport.Reports, &report_pb.ReportResponse{
			Action:      report.Action,
			Description: report.Description,
			UserId:      report.UserId,
			VaultId:     utils.NullableInt(report.VaultId),
			UpdatedAt:   report.UpdatedAt.Format("2006-01-02 15:04:05"),
		})
	}

	return &listReport, nil
}
