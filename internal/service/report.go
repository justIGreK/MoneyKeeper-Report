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

type ReportService struct {
	Transaction TransactionService
	User        UserService
}

func NewReportService(tx TransactionService, user UserService) *ReportService {
	return &ReportService{Transaction: tx, User: user}
}

const(
	Dateformat string = "2006-01-02"
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
		UserID:     userID,
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

// func (s *ReportService) GetBudgetReport(ctx context.Context, userID, budgetID string) (*models.BudgetReport, error) {
// 	budget, err := s.BudgetRepo.GetBudget(ctx, userID, budgetID)
// 	if err != nil {
// 		log.Println(err)
// 		return nil, err
// 	}
// 	if budget == nil {
// 		return nil, errors.New("budget is not found")
// 	}
// 	now := time.Now()
// 	if now.Before(budget.EndDate) {
// 		return nil, errors.New("budget is not over yet")
// 	}
// 	txs, err := s.Transaction.GetTXByTimeFrame(ctx, userID, models.TimeFrame{budget.StartDate, budget.EndDate})
// 	if err != nil {
// 		log.Println(err)
// 		return nil, err
// 	}

// 	var totalSpent float64
// 	var failedDays, successfulDays int
// 	dailyExpenses := make(map[string]float64)
// 	itemCounts := make(map[string]int)
// 	var mostExpensive models.Transaction
// 	var mostFrequent string

// 	for _, tx := range txs {
// 		totalSpent += tx.Cost

// 		day := tx.Date.Format("2006-01-02")
// 		dailyExpenses[day] += tx.Cost

// 		if tx.Cost > mostExpensive.Cost {
// 			mostExpensive = tx
// 		}

// 		itemCounts[tx.Name]++
// 		if itemCounts[tx.Name] > itemCounts[mostFrequent] {
// 			mostFrequent = tx.Name
// 		}
// 	}
// 	for _, dailyTotal := range dailyExpenses {
// 		if dailyTotal > budget.DailyAmount {
// 			failedDays++
// 		} else {
// 			successfulDays++
// 		}
// 	}

// 	budgetExceeded := totalSpent > budget.Amount
// 	report := models.BudgetReport{
// 		TotalSpent:     totalSpent,
// 		BudgetExceeded: budgetExceeded,
// 		FailedDays:     failedDays,
// 		SuccessfulDays: successfulDays,
// 		MostExpensive:  mostExpensive,
// 		MostFrequent:   mostFrequent,
// 	}
// 	return &report, nil
// }
