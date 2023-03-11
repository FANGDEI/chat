package friend

import (
	"chat/internal/app/service"
	"chat/internal/pkg/errno"
	"context"
	"gorm.io/gorm"
)

func (m *Manager) DelFriend(ctx context.Context, request *service.DelFriendRequest) (*service.Response, error) {
	m.logger.Info("Friend Service, DelFriend Service")
	// 获取当前用户信息
	uuid, friendName := request.Uuid, request.FriendName
	info, err := m.localer.GetUserInfoWithUuid(uuid)
	if err != nil {
		return nil, errno.ServerErr(errno.ErrDatabase, err.Error())
	}
	// 查询要添加好友的信息
	// 查询不到就返回 ErrUserNotFound
	finfo, err := m.localer.GetUserInfoWithName(friendName)
	if err != nil {
		if err == gorm.ErrRecordNotFound {
			return nil, errno.ServerErr(errno.ErrUserNotFound, err.Error())
		}
		return nil, errno.ServerErr(errno.ErrDatabase, err.Error())
	}
	if exists := m.localer.FriendExists(info.ID, finfo.ID); !exists {
		return nil, errno.ServerErr(errno.ErrFriendNotExists, "not friend")
	}
	err = m.localer.DeleteFriend(info.ID, finfo.ID)
	if err != nil {
		return nil, errno.ErrDatabase
	}
	return &service.Response{}, nil
}
