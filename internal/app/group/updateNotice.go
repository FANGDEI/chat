package group

import (
	"chat/internal/app/service"
	"chat/internal/pkg/errno"
	"context"
)

func (m *Manager) UpdateGroupNotice(ctx context.Context, request *service.UpdateGroupNoticeRequest) (*service.Response, error) {
	m.logger.Info("Group service, UpdateGroupNotice service")
	group, err := m.localer.GetGroupInfoWithID(request.GroupId)
	if err != nil {
		return nil, errno.ServerErr(errno.ErrDatabase, err.Error())
	}
	if group.UserID != request.UserId {
		return nil, errno.ServerErr(errno.ErrGroupRole, "no role")
	}
	err = m.localer.UpdateGroupNoticeWithID(request.GroupId, request.Notice)
	if err != nil {
		return nil, errno.ServerErr(errno.ErrDatabase, err.Error())
	}
	return &service.Response{}, nil
}
