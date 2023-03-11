package user

import (
	"chat/internal/app/service"
	"chat/internal/pkg/errno"
	"context"
	"gorm.io/gorm"
)

func (m *Manager) GetUserList(ctx context.Context, request *service.GetUserListRequest) (*service.GetUserListResponse, error) {
	m.logger.Info("User service, GetUserList service")
	//friends, err := m.localer.GetUserListWithUuid(request.Uuid)
	// 根据 uuid 获取用户ID
	info, err := m.localer.GetUserInfoWithUuid(request.Uuid)
	if err != nil {
		if err == gorm.ErrRecordNotFound {
			return nil, errno.ServerErr(errno.ErrUserNotFound, err.Error())
		}
		return nil, errno.ServerErr(errno.ErrDatabase, err.Error())
	}
	// 根据 用户ID 获取好友列表
	friends, err := m.localer.GetUserListWithID(info.ID)
	if err != nil {
		return nil, errno.ServerErr(errno.ErrDatabase, err.Error())
	}
	fs := make([]*service.SimpleUser, 0)
	for _, friend := range friends {
		fs = append(fs, &service.SimpleUser{
			Id:        friend.ID,
			Uuid:      friend.Uuid,
			Name:      friend.Name,
			Nickname:  friend.NickName,
			Gender:    friend.Gender,
			Avatar:    friend.Avatar,
			Email:     friend.Email,
			Signature: friend.Signature,
		})
	}
	return &service.GetUserListResponse{
		Friends: fs,
	}, nil
}
