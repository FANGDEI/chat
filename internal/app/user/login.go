package user

import (
	"chat/internal/app/service"
	"chat/internal/pkg/errno"
	"context"

	"gorm.io/gorm"
)

func (m *Manager) Login(ctx context.Context, request *service.UserLoginRequest) (*service.UserLoginResponse, error) {
	m.logger.Info("User service, Login service")
	user, err := m.localer.GetUserInformationWithName(request.Name)
	if err != nil {
		if err == gorm.ErrRecordNotFound {
			return nil, errno.ServerErr(errno.ErrUserNotFound, err.Error())
		}
		return nil, errno.ServerErr(errno.ErrDatabase, err.Error())
	}
	if user.Password != m.cryptoer.ToMd5(request.Password) {
		return nil, errno.ServerErr(errno.ErrUserPassword, "name or password error")
	}
	token, err := m.tokener.NewToken(user.ID)
	if err != nil {
		return nil, errno.ServerErr(errno.ErrTokenGenerate, err.Error())
	}
	return &service.UserLoginResponse{
		Token: token,
	}, nil
}
