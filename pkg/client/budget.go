package client

import (
	"context"

	budget "github.com/justIGreK/MoneyKeeper-Budget/pkg/go/budget"
	"github.com/justIGreK/MoneyKeeper-Report/internal/models"
	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"
)

type BudgetClient struct {
	client budget.BudgetServiceClient
}

func NewBudgetClient(serviceAddress string) (*BudgetClient, error) {
	conn, err := grpc.NewClient(serviceAddress, grpc.WithTransportCredentials(insecure.NewCredentials()))
	if err != nil {
		return nil, err
	}
	return &BudgetClient{
		client: budget.NewBudgetServiceClient(conn),
	}, nil
}

func (bc *BudgetClient) GetBudget(ctx context.Context, userID, budgetID string)(*models.Budget, error){
	protoBudget, err := bc.client.GetBudget(ctx, &budget.GetBudgetRequest{BudgetId: budgetID, UserId: userID})
	if err != nil{
		return nil, err
	}
	if protoBudget == nil{
		return nil, nil
	}
	budgetResp := models.Budget{
		ID: protoBudget.Budget.BudgetId,
		Name: protoBudget.Budget.Name,
		Limit: float64(protoBudget.Budget.Limit),
		StartDate: protoBudget.Budget.Start,
		EndDate: protoBudget.Budget.End,
		Category: bc.convertToCategories(protoBudget.Budget.Category),
	}
	return &budgetResp, nil
}

func (bc *BudgetClient) convertToCategories(protoCategories []*budget.Category) []models.Category {
	ctgs := make([]models.Category, len(protoCategories))
	for i, b := range protoCategories {
		ctgs[i] = models.Category{
			ID:       b.CategoryId,
			Name:     b.Name,
			Limit:     float64(b.Limit),
		}
	}
	return ctgs
}