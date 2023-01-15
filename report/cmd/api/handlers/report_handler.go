package handlers

import (
	"context"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
	"report.com/internal/report"
	"report.com/internal/utils"
	"report.com/pkg/pb"
)

type ReportHandler struct {
	pb.UnimplementedUserServer
	reportService report.ReportServiceClient
}

func NewReportHandler(_reportService report.ReportServiceClient) *ReportHandler {
	return &ReportHandler{
		reportService: _reportService,
	}
}

func (r *ReportHandler) CreateReport(ctx context.Context, req *pb.ReportRequest) (*pb.ReportResponse, error) {

	report := report.ReportRequest{
		Action:      req.Action,
		Description: req.Description,
		UserId:      req.UserId,
		VaultId:     req.VaultId,
	}

	err := r.reportService.CreateReport(report)
	if err != nil {
		return nil, status.Error(codes.Internal, err.Error())
	}

	return &pb.ReportResponse{Action: req.Action, Description: req.Description, UserId: req.UserId,
		VaultId: req.VaultId}, nil
}

func (r *ReportHandler) ReportByUserId(ctx context.Context, req *pb.Empty) (*pb.ListReportResponse, error) {

	userId, err := utils.GetMetadataByKey(ctx, "userId")
	if err != nil {
		return nil, status.Error(codes.InvalidArgument, err.Error())
	}

	reports, err := r.reportService.FindReportsByUserId(userId)
	if err != nil {
		return nil, status.Error(codes.Internal, err.Error())
	}

	var listReport pb.ListReportResponse
	for _, report := range reports {
		listReport.Reports = append(listReport.Reports, &pb.ReportResponse{
			Action:      report.Action,
			Description: report.Description,
			UserId:      report.UserId,
			VaultId:     utils.NullableInt(report.VaultId),
			UpdatedAt:   report.UpdatedAt.Format("2006-01-02 15:04:05"),
		})
	}

	return &listReport, nil
}
