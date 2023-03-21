package group

import (
	"chat/internal/app/service"
	"chat/internal/pkg/errno"
	"context"
)

func (m *Manager) GetGroupApply(ctx context.Context, request *service.GetGroupApplyRequest) (*service.GetGroupApplyResponse, error) {
	m.logger.Info("Group service, GetGroupApply service")
	list, err := m.localer.GetGroupApplyWithUserID(request.UserId)
	if err != nil {
		return nil, errno.ServerErr(errno.ErrDatabase, err.Error())
	}
	var response service.GetGroupApplyResponse
	for _, g := range list {
		response.List = append(response.List, &service.GroupApply{
			Id:       g.ID,
			UserId:   g.UserID,
			ApplyId:  g.ApplyID,
			GroupId:  g.GroupID,
			ApplyMsg: g.ApplyMsg,
		})
	}
	return &response, nil
}
