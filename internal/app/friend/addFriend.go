package friend

import (
	"chat/internal/app/service"
	"chat/internal/pkg/cacher"
	"chat/internal/pkg/constanter"
	"chat/internal/pkg/errno"
	"chat/internal/pkg/localer"
	"context"
	"time"
)

func (m *Manager) AddFriend(ctx context.Context, request *service.AddFriendRequest) (*service.Response, error) {
	m.logger.Info("Friend Service, AddFriend Service")
	id, friendID := request.Id, request.FriendId
	// 判断是否已经是好友
	if exists := m.localer.IsFriend(id, friendID); exists {
		return nil, errno.ServerErr(errno.ErrFriendExists, "friend already exists")
	}
	err := m.localer.CreateFriendApply(localer.FriendApply{
		UserID:   id,
		FriendID: friendID,
		ApplyMsg: request.ApplyMsg,
	})
	if err != nil {
		return nil, errno.ServerErr(errno.ErrDuplicateRequest, err.Error())
	}
	err = m.cacher.Send(&cacher.Message{
		From:        id,
		To:          friendID,
		Content:     request.ApplyMsg,
		ContentType: constanter.FRIEND_REQUEST,
		MessageType: constanter.MESSAGE_TYPE_USER,
		Time:        time.Now().Format("2006.01.02 15:04:05"),
	})
	if err != nil {
		return nil, errno.ServerErr(errno.ErrRedis, err.Error())
	}
	return &service.Response{}, nil
}
