package localer

type GroupMember struct {
	ID      int64 `json:"id"`
	UserID  int64 `json:"user_id"`
	GroupID int64 `json:"group_id"`
}

var groupMember = "group_member"

func (m *Manager) CreateGroupMember(gm GroupMember) error {
	return m.handler.Table(groupMember).Create(&gm).Error
}

func (m *Manager) DeleteGroupMemberWithGroupID(groupID int64) error {
	return m.handler.Table(groupMember).Where("group_id = ?", groupID).Delete(&GroupMember{}).Error
}
