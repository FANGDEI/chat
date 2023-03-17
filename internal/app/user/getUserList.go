package user

import (
	"chat/internal/app/service"
	"chat/internal/pkg/errno"
	"context"
)

func (m *Manager) GetUserList(ctx context.Context, request *service.GetUserListRequest) (*service.GetUserListResponse, error) {
	m.logger.Info("User service, GetUserList service")
	friends, err := m.localer.GetUserListWithID(request.Id)
	if err != nil {
		return nil, errno.ServerErr(errno.ErrDatabase, err.Error())
	}
	var response service.GetUserListResponse
	for _, friend := range friends {
		response.List = append(response.List, &service.SimpleUser{
			Id:        friend.ID,
			Name:      friend.Name,
			Nickname:  friend.NickName,
			Gender:    friend.Gender,
			Avatar:    friend.Avatar,
			Email:     friend.Email,
			Signature: friend.Signature,
		})
	}
	return &response, nil
}
