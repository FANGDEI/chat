package localer

type Friend struct {
	ID       int64 `json:"id"`
	UserID   int64 `json:"user_id"`
	FriendID int64 `json:"friend_id"`
	Agree    bool  `json:"agree"`
	Re       bool  `json:"re"`
}

var friend = "friend"

func (m *Manager) CreateFriend(userID, friendID int64) error {
	return m.handler.Table(friend).Create(&Friend{
		UserID:   userID,
		FriendID: friendID,
		Re:       false,
	}).Error
}

// DeleteFriend 软删除, 后台做定时任务异步删除
func (m *Manager) DeleteFriend(userID, friendID int64) error {
	return m.handler.Table(friend).Where("user_id = ? AND friend_id = ?", userID, friendID).Update("re = ?", 1).Error
}

func (m *Manager) FriendExists(userID, friendID int64) bool {
	var u SimpleUser
	err := m.handler.Table(friend).Where("user_id = ? AND friend_id = ?", userID, friendID).Take(&u).Error
	return (err == nil) && (u.ID != 0)
}
