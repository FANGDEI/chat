package localer

type FriendApply struct {
	ID       int64  `json:"id"`
	UserID   int64  `json:"user_id"`
	FriendID int64  `json:"friend_id"`
	ApplyMsg string `json:"apply_msg"`
}

var friendApply = "friend_apply"

// CreateFriendApply 被请求的用户ID做UserID
func (m *Manager) CreateFriendApply(apply FriendApply) error {
	return m.handler.Table(friendApply).Create(&FriendApply{
		UserID:   apply.FriendID,
		FriendID: apply.UserID,
		ApplyMsg: apply.ApplyMsg,
	}).Error
}

func (m *Manager) DeleteApply(userID, friendID int64) error {
	return m.handler.Table(friendApply).Where("user_id = ? AND friend_id = ?", userID, friendID).Delete(&FriendApply{}).Error
}

func (m *Manager) GetApply(id int64) ([]FriendApply, error) {
	var fs []FriendApply
	err := m.handler.Table(friendApply).Where("user_id = ?", id).Find(&fs).Error
	return fs, err
}
