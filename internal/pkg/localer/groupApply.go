package localer

type GroupApply struct {
	ID       int64  `json:"id"`
	UserID   int64  `json:"user_id"`
	ApplyID  int64  `json:"apply_id"`
	GroupID  int64  `json:"group_id"`
	ApplyMsg string `json:"apply_msg"`
}

var groupApply = "group_apply"

func (m *Manager) CreateGroupApply(apply GroupApply) error {
	return m.handler.Table(groupApply).Create(&apply).Error
}

func (m *Manager) DeleteGroupApply(applyID, groupID int64) error {
	return m.handler.Table(groupApply).Where("apply_id = ? AND group_id = ?", applyID, groupID).Delete(&GroupApply{}).Error
}

func (m *Manager) DeleteGroupApplyWithGroupID(groupID int64) error {
	return m.handler.Table(groupApply).Where("group_id = ?", groupID).Delete(&GroupApply{}).Error
}

func (m *Manager) GetGroupApplyWithUserID(userID int64) ([]GroupApply, error) {
	var gs []GroupApply
	err := m.handler.Table(groupApply).Where("user_id = ?", userID).Find(&gs).Error
	return gs, err
}
