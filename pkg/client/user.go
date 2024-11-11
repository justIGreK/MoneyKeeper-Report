package client

import (
	"context"
	"errors"

		user "github.com/justIGreK/MoneyKeeper-User/pkg/go/user"
	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"
)

type UserClient struct {
	client user.UserServiceClient
}

func NewUserClient(serviceAddress string) (*UserClient, error) {
	conn, err := grpc.NewClient(serviceAddress, grpc.WithTransportCredentials(insecure.NewCredentials()))
	if err != nil {
		return nil, err
	}
	return &UserClient{
		client: user.NewUserServiceClient(conn),
	}, nil
}

func (uc *UserClient) GetUser(ctx context.Context, id string) (string, string, error) {
	req := &user.GetUserRequest{UserId: id}
	res, err := uc.client.GetUser(ctx, req)
	if err != nil {
		return "", "", err
	}
	if req == nil {
		return "", "", errors.New("user is not find")
	}
	return res.Id, res.Name, nil
}
