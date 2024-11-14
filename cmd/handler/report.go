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
	GetPeriodSummary(ctx context.Context, period models.GetPeriodSummary) (*models.ReportResponse, error)
	GetBudgetReport(ctx context.Context, userID, budgetID string) (*models.BudgetReport, error)
}

func (s *ReportServiceServer) GetSummaryReport(ctx context.Context, req *reportProto.GetSummaryReportRequest) (*reportProto.GetSummaryReportResponse, error) {
	period := models.GetPeriodSummary{UserID: req.UserId}
	if req.Start != "" {
		period.Start = &req.Start
	} 
	if req.End != "" {
		period.End = &req.End
	} 
	if req.Period != "" {
		period.Period = &req.Period
	} 
	report, err := s.ReportSRV.GetPeriodSummary(ctx, period)
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
	return &reportProto.ReportResponse{Period: report.Period,
		TotalSpent: float32(report.TotalSpent), TransactionCount: int32(report.TransactionCount),
		Categories: categories}
}

func (s *ReportServiceServer) GetBudgetReport(ctx context.Context, req *reportProto.GetBudgetReportRequest) (*reportProto.GetBudgetReportResponse, error) {
	report, err := s.ReportSRV.GetBudgetReport(ctx, req.UserId, req.BudgetId)
	if err != nil {
		return nil, err
	}
	protoReport := convertToProtoBudgetReport(report)
	return &reportProto.GetBudgetReportResponse{
		Report: protoReport,
	}, nil
}

func convertToProtoBudgetReport(report *models.BudgetReport) *reportProto.BudgetReport {
	reqCategories := make([]*reportProto.RequiredCategoryReport, len(report.Categories)) 
	categories := make([]*reportProto.CategoryReport, len(report.Categories))
	for i, b := range report.Categories {
		categories[i] = &reportProto.CategoryReport{
			Name:  b.Name,
			Total: float32(b.Total),
			Count: int32(b.Count),
		}
	}
	for i, b := range report.RequiredCategories {
		reqCategories[i] = &reportProto.RequiredCategoryReport{
			Name:  b.Name,
			Total: float32(b.Total),
			Limit: float32(b.Limit),
			Count: int32(b.Count),
		}
	}

	return &reportProto.BudgetReport{
		BudgetName: report.BudgetName, 
		Period: report.Period,
		LeftDays: report.LeftDays,
		Limit: report.Limit,
		TotalSpent: float32(report.TotalSpent), 
		TransactionCount: int32(report.TransactionCount),
		ReqCategories: reqCategories,
		Categories: categories,
	}
}