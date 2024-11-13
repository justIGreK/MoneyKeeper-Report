package service

import (
	"context"
	"errors"
	"fmt"
	"log"
	"time"

	"github.com/justIGreK/MoneyKeeper-Report/internal/models"
)

type TransactionService interface {
	GetTransactionList(ctx context.Context, userID string) ([]models.Transaction, error)
	GetTXByTimeFrame(ctx context.Context, userID, start, end string) ([]models.Transaction, error)
}

type UserService interface {
	GetUser(ctx context.Context, id string) (string, string, error)
}

type BudgetSevice interface {
	GetBudget(ctx context.Context, userID, budgetID string) (*models.Budget, error)
}

type ReportService struct {
	Transaction TransactionService
	User        UserService
	Budget      BudgetSevice
}

func NewReportService(tx TransactionService, user UserService, budget BudgetSevice) *ReportService {
	return &ReportService{Transaction: tx, User: user, Budget: budget}
}

const (
	Dateformat     string = "2006-01-02"
	DateTimeformat string = "2006-01-02T15:04:05"
)

func (s *ReportService) GetPeriodSummary(ctx context.Context, userID, period string) (*models.ReportResponse, error) {
	id, _, err := s.User.GetUser(ctx, userID)
	if err != nil {
		log.Println(err)
		return nil, err
	}
	if id == "" {
		return nil, errors.New("user not found")
	}
	startDate, endDate := s.getPeriodDates(period)
	if startDate.IsZero() {
		return nil, errors.New("invalid period")
	}
	txs, err := s.Transaction.GetTXByTimeFrame(ctx, userID, startDate.Format(Dateformat), endDate.Format(Dateformat))
	if err != nil {
		log.Println(err)
		return nil, err
	}

	report := models.ReportResponse{
		Period:     fmt.Sprintf("%s - %s", startDate.Format("2006-01-02"), endDate.Format("2006-01-02")),
		Categories: []models.CategoryReport{},
	}

	categoryTotals := make(map[string]float32)
	categoryCounts := make(map[string]int)
	transactionCount := 0
	totalSpent := float32(0.0)

	for _, txn := range txs {
		totalSpent += txn.Cost
		transactionCount++

		if txn.Category != "" {
			categoryTotals[txn.Category] += txn.Cost
			categoryCounts[txn.Category]++
		}
	}

	report.TotalSpent = totalSpent
	report.TransactionCount = transactionCount

	for category, total := range categoryTotals {
		report.Categories = append(report.Categories, models.CategoryReport{
			Name:  category,
			Total: total,
			Count: categoryCounts[category],
		})
	}

	return &report, nil
}

func (s *ReportService) getPeriodDates(period string) (time.Time, time.Time) {
	now := time.Now().UTC()
	switch period {
	case "day":
		start := now
		return start.AddDate(0, 0, -1), start
	case "week":
		start := now
		return start.AddDate(0, 0, -7), start
	case "month":
		start := now
		return start.AddDate(0, -1, 0), start
	default:
		return time.Time{}, time.Time{}
	}
}

func (s *ReportService) GetBudgetReport(ctx context.Context, userID, budgetID string) (*models.BudgetReport, error) {
	id, _, err := s.User.GetUser(ctx, userID)
	if err != nil {
		log.Println(err)
		return nil, err
	}
	if id == "" {
		return nil, errors.New("user not found")
	}
	budget, err := s.Budget.GetBudget(ctx, userID, budgetID)
	if err != nil {
		log.Println(err)
		return nil, err
	}
	if budget == nil {
		return nil, errors.New("budget is not found")
	}
	daysLeft := 0.0
	now := time.Now()
	endDate, err := time.Parse(Dateformat, budget.EndDate)
	if err != nil{
		log.Println("invalid endtime")
		return nil, err
	}
	if !endDate.Before(now){
		duration := endDate.Sub(now)
		daysLeft = duration.Hours()/24.0
	}
	txs, err := s.Transaction.GetTXByTimeFrame(ctx, userID, budget.StartDate, budget.EndDate)
	if err != nil {
		log.Println(err)
		return nil, err
	}

	report := models.BudgetReport{
		BudgetName: budget.Name,
		Period:     fmt.Sprintf("%s - %s", budget.StartDate, budget.EndDate),
		Limit: float32(budget.Limit),
		LeftDays: float32(daysLeft),
		RequiredCategories: []models.RequiredCategoryReport{},
		Categories: []models.CategoryReport{},
	}
	categoryTotals := make(map[string]float32)
	categoryCounts := make(map[string]int)
	transactionCount := 0
	totalSpent := float32(0.0)
	requiredCategories := make(map[string]float32)
	for _, categories := range budget.Category{
		requiredCategories[categories.Name] = float32(categories.Limit)
	}
	for _, txn := range txs {
		totalSpent += txn.Cost
		transactionCount++

		if txn.Category != "" {
			categoryTotals[txn.Category] += txn.Cost
			categoryCounts[txn.Category]++
		}
	}

	report.TotalSpent = totalSpent
	report.TransactionCount = transactionCount

	for category, total := range categoryTotals {
		if limit, exists := requiredCategories[category]; exists {
			report.RequiredCategories = append(report.RequiredCategories, models.RequiredCategoryReport{
				Name: category,
				Total: total,
				Limit: limit,
				Count: categoryCounts[category],
			})
		}else {
			report.Categories = append(report.Categories, models.CategoryReport{
				Name:  category,
				Total: total,
				Count: categoryCounts[category],
			})
		}
	}
	report.TransactionCount = transactionCount
	return &report, nil
}
