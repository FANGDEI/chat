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

func (m *Manager) IsMember(userID, groupID int64) bool {
	var id int64
	err := m.handler.Table(groupMember).Select("id").Where("user_id = ? AND group_id = ?", userID, groupID).Take(&id).Error
	return err == nil && id != 0
}

func (m *Manager) DeleteGroupMemberWithUserIDAndGroupID(userID, groupID int64) error {
	return m.handler.Table(groupMember).Where("user_id = ? AND group_id = ?", userID, groupID).Delete(&GroupMember{}).Error
}
