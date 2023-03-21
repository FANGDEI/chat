package user

import (
	"chat/internal/app/service"
	"chat/internal/pkg/errno"
	"context"
)

func (m *Manager) GetUserGroupList(ctx context.Context, request *service.GetUserGroupListRequest) (*service.GetUserGroupListResponse, error) {
	m.logger.Info("User service, GetUserGroupList service")
	list, err := m.localer.GetUserGroupListWithID(request.Id)
	if err != nil {
		return nil, errno.ServerErr(errno.ErrDatabase, err.Error())
	}
	var response service.GetUserGroupListResponse
	for _, g := range list {
		response.List = append(response.List, &service.Group{
			Id:     g.ID,
			UserId: g.UserID,
			Name:   g.Name,
			Avatar: g.Avatar,
			Notice: g.Notice,
		})
	}
	return &response, nil
}
