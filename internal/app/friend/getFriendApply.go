package friend

import (
	"chat/internal/app/service"
	"chat/internal/pkg/errno"
	"context"
)

func (m *Manager) GetFriendApply(ctx context.Context, request *service.GetFriendApplyRequest) (*service.GetFriendApplyResponse, error) {
	m.logger.Info("Friend service, GetFriendApply service")
	list, err := m.localer.GetFriendApplyWithID(request.UserId)
	if err != nil {
		return nil, errno.ServerErr(errno.ErrDatabase, err.Error())
	}
	var response service.GetFriendApplyResponse
	for _, f := range list {
		response.List = append(response.List, &service.FriendApply{
			Id:       f.ID,
			UserId:   f.UserID,
			FriendId: f.FriendID,
			ApplyMsg: f.ApplyMsg,
		})
	}
	return &response, nil
}
