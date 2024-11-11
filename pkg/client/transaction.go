package client

import (
	"context"

	tx "github.com/justIGreK/MoneyKeeper-Transaction/pkg/go/transaction"
	"github.com/justIGreK/MoneyKeeper-Report/internal/models"
	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"
)

type TransactionClient struct {
	client tx.TransactionServiceClient
}

func NewTransactionClient(serviceAddress string) (*TransactionClient, error) {
	conn, err := grpc.NewClient(serviceAddress, grpc.WithTransportCredentials(insecure.NewCredentials()))
	if err != nil {
		return nil, err
	}
	return &TransactionClient{
		client: tx.NewTransactionServiceClient(conn),
	}, nil
}

func (tc *TransactionClient) GetTransactionList(ctx context.Context, userID string) ([]models.Transaction, error) {
	req := &tx.GetTransactionListRequest{
		UserId: userID,
	}
	res, err := tc.client.GetTransactionList(ctx, req)
	if err != nil {
		return nil, err
	}
	txs := tc.convertToTxs(res.Transactions)
	return txs, nil
}

func (tc *TransactionClient) convertToTxs(protoBudgets []*tx.Transaction) []models.Transaction {
	txs := make([]models.Transaction, len(protoBudgets))
	for i, b := range protoBudgets {
		txs[i] = models.Transaction{
			ID:       b.Id,
			UserID:   b.UserId,
			Category: b.Category,
			Name:     b.Name,
			Cost:     b.Cost,
			Date:     b.Date,
		}
	}
	return txs
}


func (tc *TransactionClient) GetTXByTimeFrame(ctx context.Context, userID, start, end string) ([]models.Transaction, error) {
	req := &tx.GetTXByTimeFrameRequest{
		UserId:    userID,
		StartDate: start,
		EndDate:   end,
	}
	res, err := tc.client.GetTXByTimeFrame(ctx, req)
	if err != nil {
		return nil, err
	}

	txs := tc.convertToTxs(res.Transactions)

	return txs, nil
}
