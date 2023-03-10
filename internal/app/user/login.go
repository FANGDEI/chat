package user

import (
	"chat/internal/app/service"
	"chat/internal/pkg/errno"
	"context"
	"gorm.io/gorm"
)

func (m *Manager) Login(ctx context.Context, request *service.UserLoginRequest) (*service.UserLoginResponse, error) {
	m.logger.Info("User service, Login service")
	user, err := m.localer.GetUserInfoWithName(request.Name)
	if err != nil {
		if err == gorm.ErrRecordNotFound {
			return nil, errno.ErrUserNotFound
		}
		return nil, errno.ErrDatabase
	}
	if user.Password != m.cryptoer.ToMd5(request.Password) {
		return nil, errno.ErrUserPassword
	}
	token, err := m.tokener.NewToken(user.ID, user.Uuid)
	if err != nil {
		return nil, errno.ErrTokenGenerate
	}
	return &service.UserLoginResponse{
		Token: token,
	}, nil
}
