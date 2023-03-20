package group

import (
	"chat/internal/app/service"
	"chat/internal/pkg/cacher"
	"chat/internal/pkg/constanter"
	"chat/internal/pkg/errno"
	"chat/internal/pkg/localer"
	"context"

	"gorm.io/gorm"
)

func (m *Manager) AddGroup(ctx context.Context, request *service.AddGroupRequest) (*service.Response, error) {
	m.logger.Info("Group service, AddGroup service")
	if exists := m.localer.IsMember(request.UserId, request.GroupId); exists {
		return nil, errno.ServerErr(errno.ErrGroupMemberAlready, "group member already")
	}
	group, err := m.localer.GetGroupInfoWithID(request.GroupId)
	if err != nil {
		if err == gorm.ErrRecordNotFound {
			return nil, errno.ServerErr(errno.ErrGroupNotFound, err.Error())
		}
		return nil, errno.ServerErr(errno.ErrDatabase, err.Error())
	}
	err = m.localer.CreateGroupApply(localer.GroupApply{
		UserID:   group.UserID,
		ApplyID:  request.UserId,
		GroupID:  request.GroupId,
		ApplyMsg: request.ApplyMsg,
	})
	if err != nil {
		return nil, errno.ServerErr(errno.ErrGroupDuplicateRequest, err.Error())
	}
	err = m.cacher.Send(&cacher.Message{
		From:        request.UserId,
		To:          group.UserID,
		GroupID:     request.GroupId,
		Content:     request.ApplyMsg,
		ContentType: constanter.GROUP_REQUEST,
		MessageType: constanter.MESSAGE_TYPE_USER,
	})
	if err != nil {
		return nil, errno.ServerErr(errno.ErrRedis, err.Error())
	}
	return &service.Response{}, nil
}
