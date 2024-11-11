package handler

import (
	reportProto "github.com/justIGreK/MoneyKeeper-Report/pkg/go/report"
	"google.golang.org/grpc"
)
type Handler struct {
	server grpc.ServiceRegistrar
	report ReportService
}

func NewHandler(grpcServer grpc.ServiceRegistrar, reportSRV ReportService) *Handler{
	return &Handler{server: grpcServer, report: reportSRV}
}
func(h *Handler) RegisterServices(){
	h.registerReportService(h.server, h.report)
}

func(h *Handler) registerReportService(server grpc.ServiceRegistrar, report ReportService){
	reportProto.RegisterReportServiceServer(server, &ReportServiceServer{ReportSRV: report})
}