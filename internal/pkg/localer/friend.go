package localer

type Friend struct {
	ID       int64 `json:"id"`
	UserID   int64 `json:"user_id"`
	FriendID int64 `json:"friend_id"`
}

var friend = "friend"

func (m *Manager) CreateFriend(userID, friendID int64) error {
	err := m.execTx(func(m *Manager) error {
		err := m.handler.Table(friend).Create(&Friend{
			UserID:   userID,
			FriendID: friendID,
		}).Error
		if err != nil {
			return err
		}
		err = m.handler.Table(friend).Create(&Friend{
			UserID:   friendID,
			FriendID: userID,
		}).Error
		if err != nil {
			return err
		}
		return nil
	})
	return err
}

func (m *Manager) DeleteFriend(userID, friendID int64) error {
	err := m.execTx(func(m *Manager) error {
		err := m.handler.Table(friend).Where("user_id = ? AND friend_id = ?", userID, friendID).Delete(&Friend{}).Error
		if err != nil {
			return err
		}
		err = m.handler.Table(friend).Where("user_id = ? AND friend_id = ?", friendID, userID).Delete(&Friend{}).Error
		if err != nil {
			return err
		}
		return nil
	})
	return err
}
