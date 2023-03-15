package friend

import (
	"chat/internal/app/service"
	"chat/internal/pkg/cacher"
	"chat/internal/pkg/constanter"
	"chat/internal/pkg/errno"
	"context"
	"time"
)

func (m *Manager) AddFriend(ctx context.Context, request *service.AddFriendRequest) (*service.Response, error) {
	m.logger.Info("Friend Service, AddFriend Service")
	// 获取当前用户信息
	uuid, friendName := request.Uuid, request.FriendName
	info, err := m.localer.GetUserInfoWithUuid(uuid)
	if err != nil {
		return nil, errno.ServerErr(errno.ErrDatabase, err.Error())
	}
	// 查询要添加好友的信息
	finfo, err := m.localer.GetUserInfoWithName(friendName)
	if err != nil {
		return nil, errno.ServerErr(errno.ErrDatabase, err.Error())
	}
	// 判断是否已经是好友
	if exists := m.localer.FriendExists(info.ID, finfo.ID); exists {
		return nil, errno.ServerErr(errno.ErrFriendExists, "friend already exists")
	}
	err = m.localer.CreateFriend(info.ID, finfo.ID)
	if err != nil {
		return nil, errno.ServerErr(errno.ErrDuplicateRequest, err.Error())
	}
	err = m.cacher.Send(&cacher.Message{
		From:        info.Uuid,
		To:          finfo.Uuid,
		Content:     info.Name + "请求添加你为好友",
		ContentType: constanter.FRIEND_REQUEST,
		MessageType: constanter.MESSAGE_TYPE_USER,
		Time:        time.Now().Format("2006.01.02 15:04:05"),
	})
	if err != nil {
		return nil, errno.ServerErr(errno.ErrRedis, err.Error())
	}
	return &service.Response{}, nil
}
