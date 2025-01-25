package server

import (
	"context"
	"grpc_service_logger/pkg/domain/audit"
)

type AuditService interface {
	Insert(ctx context.Context, req *audit.LogRequest) error
}

type AuditServer struct {
	service AuditService
}

func NewAuditServer(service AuditService) *AuditServer {
	return &AuditServer{service: service}
}

func (h *AuditServer) Log(ctx context.Context, req *audit.LogRequest) (*audit.Empty, error) {
	err := h.service.Insert(ctx, req)
	if err != nil {
		return nil, err
	}

	return &audit.Empty{}, err
}
