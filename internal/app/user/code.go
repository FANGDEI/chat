package user

import (
	"chat/internal/app/service"
	"chat/internal/pkg/emailer"
	"chat/internal/pkg/errno"
	"context"
	"fmt"
	"math/rand"
)

func (m *Manager) Code(ctx context.Context, request *service.CodeRequest) (*service.Response, error) {
	m.logger.Info("User service, Code service")
	if !m.emailer.IsEmail(request.Email) {
		return nil, errno.ServerErr(errno.ErrEmailFormat, "email format error")
	}
	if m.cacher.IsEmailBanned(request.Email) {
		return nil, errno.ServerErr(errno.ErrEmailBan, "email ban to send")
	}
	code := fmt.Sprintf("%06d", rand.Intn(1000000))
	err := m.cacher.SetEmailCode(request.Email, code)
	if err != nil {
		return nil, errno.ServerErr(errno.ErrRedis, err.Error())
	}
	err = m.cacher.SetEmailToBan(request.Email)
	if err != nil {
		return nil, errno.ServerErr(errno.ErrRedis, err.Error())
	}
	m.emailer.SendEmail(emailer.Information{
		To:   request.Email,
		Code: code,
	})
	return &service.Response{}, nil
}
