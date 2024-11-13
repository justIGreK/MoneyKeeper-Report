package main

import (
	"log"
	"net"

	"github.com/justIGreK/MoneyKeeper-Report/cmd/handler"
	"github.com/justIGreK/MoneyKeeper-Report/internal/service"
	"github.com/justIGreK/MoneyKeeper-Report/pkg/client"
	"google.golang.org/grpc"
	"google.golang.org/grpc/reflection"
)

func main() {
	user, err := client.NewUserClient("localhost:50052")
	if err != nil {
		log.Fatal(err)
	}
	tx, err := client.NewTransactionClient("localhost:50053")
	if err != nil {
		log.Fatal(err)
	}
	budget, err := client.NewBudgetClient("localhost:50051")
	if err != nil {
		log.Fatal(err)
	}

	reportSRV := service.NewReportService(tx, user, budget)
	lis, err := net.Listen("tcp", ":50054")
	if err != nil {
		log.Fatalf("failed to listen: %v", err)
	}
	grpcServer := grpc.NewServer()

	handler := handler.NewHandler(grpcServer, reportSRV)
	handler.RegisterServices()
	reflection.Register(grpcServer)

	log.Printf("Starting gRPC server on :50054")
	if err := grpcServer.Serve(lis); err != nil {
		log.Fatalf("failed to serve: %v", err)
	}
}
