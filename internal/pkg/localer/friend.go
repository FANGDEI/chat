package localer

type Friend struct {
	ID       int64 `json:"id"`
	UserID   int64 `json:"user_id"`
	FriendID int64 `json:"friend_id"`
	Agree    bool  `json:"agree"`
}

var friend = "friend"

func (m *Manager) CreateFriend(userID, friendID int64) error {
	return m.handler.Table(friend).Create(&Friend{
		UserID:   userID,
		FriendID: friendID,
	}).Error
}

func (m *Manager) DeleteFriend(userID, friendID int64) error {
	return m.handler.Table(friend).Where("user_id = ? AND friend_id = ?", userID, friendID).Delete(&Friend{}).Error
}

// TODO:
func (m *Manager) AcceptFriend(selfID, otherID int64) error {
	return m.handler.Table(friend).Where("user_id = ? AND friend_id = ?", selfID, otherID).Update("agree", 1).Error
}

func (m *Manager) DeAcceptFriend(selfID, otherID int64) error {
	return m.handler.Table(friend).Where("user_id = ? AND friend_id = ?", selfID, otherID).Delete(&Friend{}).Error
}

func (m *Manager) FriendExists(userID, friendID int64) bool {
	var u SimpleUser
	err := m.handler.Table(friend).Where("user_id = ? AND friend_id = ? AND agree = 1", userID, friendID).Take(&u).Error
	return (err == nil) && (u.ID != 0)
}
