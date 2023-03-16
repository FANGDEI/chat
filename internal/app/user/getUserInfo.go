package user

import (
	"chat/internal/app/service"
	"chat/internal/pkg/errno"
	"context"
)

func (m *Manager) GetUserInfo(ctx context.Context, request *service.GetUserInfoRequest) (*service.GetUserInfoResponse, error) {
	m.logger.Info("User service, GetUserInfo service")
	userInfo, err := m.localer.GetUserInfoWithID(request.Id)
	if err != nil {
		return nil, errno.ServerErr(errno.ErrDatabase, err.Error())
	}
	return &service.GetUserInfoResponse{
		User: &service.SimpleUser{
			Id:        userInfo.ID,
			Name:      userInfo.Name,
			Nickname:  userInfo.NickName,
			Gender:    userInfo.Gender,
			Avatar:    m.obser.GetURL(userInfo.Avatar),
			Email:     userInfo.Email,
			Signature: userInfo.Signature,
		},
	}, nil
}
