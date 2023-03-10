package user

import (
	"chat/internal/app/service"
	"chat/internal/pkg/errno"
	"context"
)

func (m *Manager) GetUserInfo(ctx context.Context, request *service.GetUserInfoRequest) (*service.GetUserInfoResponse, error) {
	m.logger.Info("User service, GetUserInfo service")
	userInfo, err := m.localer.GetUserInfoWithUuid(request.Uuid)
	if err != nil {
		return nil, errno.ServerErr(errno.ErrDatabase, err.Error())
	}
	return &service.GetUserInfoResponse{
		User: &service.SimpleUser{
			Id:        userInfo.ID,
			Uuid:      userInfo.Uuid,
			Name:      userInfo.Name,
			Nickname:  userInfo.NickName,
			Gender:    userInfo.Gender,
			Avatar:    userInfo.Avatar,
			Email:     userInfo.Email,
			Signature: userInfo.Signature,
		},
	}, nil
}
