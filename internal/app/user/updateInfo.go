package user

import (
	"chat/internal/app/service"
	"chat/internal/pkg/errno"
	"chat/internal/pkg/localer"
	"context"
)

func (m *Manager) UpdateUserInfo(ctx context.Context, request *service.UpdateUserInfoRequest) (*service.Response, error) {
	m.logger.Info("User service, UpdateUserInfo service")
	err := m.localer.UpdateUserInfoWithUuid(request.Uuid, localer.SimpleUser{
		NickName:  request.User.Nickname,
		Gender:    request.User.Gender,
		Avatar:    request.User.Avatar,
		Email:     request.User.Email,
		Signature: request.User.Signature,
	})
	if err != nil {
		return nil, errno.ErrDatabase
	}
	return &service.Response{}, nil
}
