package user

import (
	"chat/internal/app/service"
	"chat/internal/pkg/errno"
	"chat/internal/pkg/localer"
	"context"
	"strings"
)

func (m *Manager) UpdateUserInfo(ctx context.Context, request *service.UpdateUserInfoRequest) (*service.Response, error) {
	m.logger.Info("User service, UpdateUserInfo service")
	str := strings.Split(request.User.Avatar, "/")
	err := m.localer.UpdateUserInfoWithUuid(request.Uuid, localer.SimpleUser{
		NickName:  request.User.Nickname,
		Gender:    request.User.Gender,
		Avatar:    str[len(str)-1],
		Signature: request.User.Signature,
	})
	if err != nil {
		return nil, errno.ServerErr(errno.ErrDatabase, err.Error())
	}
	return &service.Response{}, nil
}
