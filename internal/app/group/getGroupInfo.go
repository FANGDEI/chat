package group

import (
	"chat/internal/app/service"
	"chat/internal/pkg/errno"
	"context"
	"gorm.io/gorm"
)

func (m *Manager) GetGroupInfo(ctx context.Context, request *service.GetGroupInfoRequest) (*service.GetGroupInfoResponse, error) {
	m.logger.Info("Group service, GetGroupInfo service")
	group, err := m.localer.GetGroupInfoWithName(request.GroupName)
	if err != nil {
		if err == gorm.ErrRecordNotFound {
			return nil, errno.ServerErr(errno.ErrGroupNotFound, err.Error())
		}
		return nil, errno.ServerErr(errno.ErrDatabase, err.Error())
	}
	return &service.GetGroupInfoResponse{
		Group: &service.Group{
			Id:     group.ID,
			UserId: group.UserID,
			Name:   group.Name,
			Avatar: m.obser.GetURL(group.Avatar),
			Notice: group.Notice,
		},
	}, nil
}
