package user

import (
	"chat/internal/app/service"
	"chat/internal/pkg/errno"
	"context"
	"gorm.io/gorm"
)

func (m *Manager) GetOtherUserInfo(ctx context.Context, request *service.GetOtherUserInfoRequest) (*service.GetOtherUserInfoResponse, error) {
	m.logger.Info("User service, GetOtherUserInfo")
	userInfo, err := m.localer.GetUserInfoWithName(request.Name)
	if err != nil {
		if err == gorm.ErrRecordNotFound {
			return nil, errno.ServerErr(errno.ErrUserNotFound, err.Error())
		}
		return nil, errno.ServerErr(errno.ErrDatabase, err.Error())
	}
	return &service.GetOtherUserInfoResponse{
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
