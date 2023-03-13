package chat

import (
	"chat/internal/app/service"
	"context"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
)

func (m *Manager) Get(ctx context.Context, request *service.GetRequest) (*service.GetResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method Get not implemented")
}
