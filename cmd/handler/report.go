package handler

import (
	"context"
	"github.com/justIGreK/MoneyKeeper-Report/internal/models"

	reportProto "github.com/justIGreK/MoneyKeeper-Report/pkg/go/report"
)

type ReportServiceServer struct {
	reportProto.UnimplementedReportServiceServer
	ReportSRV ReportService
}

type ReportService interface {
	GetPeriodSummary(ctx context.Context, userID, period string) (*models.ReportResponse, error)
}

func (s *ReportServiceServer) GetSummaryReport(ctx context.Context, req *reportProto.GetSummaryReportRequest) (*reportProto.GetSummaryReportResponse, error) {

	report, err := s.ReportSRV.GetPeriodSummary(ctx, req.UserId, req.Period)
	if err != nil {
		return nil, err
	}

	protoReport := convertToProtoReportResponse(report)
	return &reportProto.GetSummaryReportResponse{
		Report: protoReport,
	}, nil

}

func convertToProtoReportResponse(report *models.ReportResponse) *reportProto.ReportResponse {
	categories := make([]*reportProto.CategoryReport, len(report.Categories))
	for i, b := range report.Categories {
		categories[i] = &reportProto.CategoryReport{
			Name:  b.Name,
			Total: float32(b.Total),
			Count: int32(b.Count),
		}
	}

	return &reportProto.ReportResponse{UserId: report.UserID, Period: report.Period,
		TotalSpent: float32(report.TotalSpent), TransactionCount: int32(report.TransactionCount),
		Categories: categories}
}
