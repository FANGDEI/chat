package file

import (
	"chat/internal/app/service"
	"chat/internal/pkg/errno"
	"context"
)

func (m *Manager) Upload(ctx context.Context, request *service.UploadRequest) (*service.UploadResponse, error) {
	m.logger.Info("File Service, Upload service")
	fileName := m.cryptoer.GetUUIDWithoutSplit()
	info, err := m.obser.Upload(context.Background(), fileName, request.Suffix, request.Data)
	if err != nil {
		return nil, errno.ServerErr(errno.ErrUploadFile, err.Error())
	}
	return &service.UploadResponse{
		Url: m.obser.GetURL(info.Key),
	}, nil
}
