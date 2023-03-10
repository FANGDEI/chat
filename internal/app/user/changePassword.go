package user

import (
	"chat/internal/app/service"
	"chat/internal/pkg/errno"
	"context"
)

func (m *Manager) ChangePassword(ctx context.Context, request *service.UserChangePasswordRequest) (*service.Response, error) {
	m.logger.Info("User service, ChangePassword service")
	uuid, oldPwd, newPwd := request.Uuid, request.OldPwd, request.NewPwd
	user, err := m.localer.GetUserInformationWithUuid(uuid)
	if err != nil {
		return nil, errno.ErrDatabase
	}
	if user.Password != m.cryptoer.ToMd5(oldPwd) {
		return nil, errno.ErrUserOldPassword
	}
	err = m.localer.UpdateUserPasswordWithUuid(uuid, newPwd)
	if err != nil {
		return nil, errno.ErrDatabase
	}
	return &service.Response{}, nil
}
