package group

import (
	"chat/internal/app/service"
	"chat/internal/pkg/errno"
	"chat/internal/pkg/localer"
	"context"
)

func (m *Manager) AcceptApply(ctx context.Context, request *service.AcceptApplyRequest) (*service.Response, error) {
	m.logger.Info("Group service, AcceptApply service")
	if request.Accept {
		err := m.localer.CreateGroupMember(localer.GroupMember{
			UserID:  request.UserId,
			GroupID: request.GroupId,
		})
		if err != nil {
			return nil, errno.ServerErr(errno.ErrDatabase, err.Error())
		}
	}
	err := m.localer.DeleteGroupApply(request.UserId, request.GroupId)
	if err != nil {
		return nil, errno.ServerErr(errno.ErrDatabase, err.Error())
	}
	return &service.Response{}, nil
}
