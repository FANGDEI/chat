package localer

type GroupApply struct {
	ID       int64  `json:"id"`
	UserID   int64  `json:"user_id"`
	GroupID  int64  `json:"group_id"`
	ApplyMsg string `json:"apply_msg"`
}

var groupApply = "group_apply"

func (m *Manager) CreateGroupApply(apply GroupApply) error {
	return nil
}

func (m *Manager) DeleteGroupApplyWithGroupID(groupID int64) error {
	return m.handler.Table(groupApply).Where("group_id = ?", groupID).Delete(&GroupApply{}).Error
}
