package user

import (
	"chat/internal/app/service"
	"chat/internal/pkg/errno"
	"chat/internal/pkg/localer"
	"context"
)

func (m *Manager) Register(ctx context.Context, request *service.UserRegisterRequest) (*service.Response, error) {
	m.logger.Info("User service, Register service")
	if len(request.Name) < 6 {
		return nil, errno.ServerErr(errno.ErrUserNameLength, "error length of Name")
	}
	if len(request.Password) < 8 || len(request.Password) > 16 {
		return nil, errno.ServerErr(errno.ErrUserPasswordLength, "error length of Password")
	}
	// 判断验证码
	err := m.cacher.AuthEmail(request.Email, request.Code)
	if err != nil {
		return nil, errno.ServerErr(errno.ErrEmailAuth, err.Error())
	}
	// 创建用户
	err = m.localer.CreateUser(localer.User{
		Name:     request.Name,
		Password: m.cryptoer.ToMd5(request.Password),
		NickName: m.cryptoer.GetUUIDWithoutSplit(),
		Email:    request.Email,
		Avatar:   "default.jpg",
	})
	if err != nil {
		return nil, errno.ServerErr(errno.ErrUserExists, err.Error())
	}
	return &service.Response{}, nil
}
