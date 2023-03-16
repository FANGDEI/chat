package localer

type User struct {
	ID        int64  `json:"id"`
	Name      string `json:"name"`
	Password  string `json:"password"`
	NickName  string `json:"nick_name"`
	Gender    string `json:"gender"`
	Avatar    string `json:"avatar"`
	Email     string `json:"email"`
	Signature string `json:"signature"`
}

type SimpleUser struct {
	ID        int64  `json:"id"`
	Name      string `json:"name"`
	NickName  string `json:"nick_name"`
	Gender    string `json:"gender"`
	Avatar    string `json:"avatar"`
	Email     string `json:"email"`
	Signature string `json:"signature"`
}

var user = "user"

func (m *Manager) CreateUser(u User) error {
	return m.handler.Table(user).Create(&u).Error
}

func (m *Manager) GetUserInfoWithID(id int64) (SimpleUser, error) {
	var u SimpleUser
	err := m.handler.Table(user).Where("id = ?", id).Take(&u).Error
	return u, err
}

func (m *Manager) GetUserInfoWithName(name string) (SimpleUser, error) {
	var u SimpleUser
	err := m.handler.Table(user).Where("name = ?", name).Take(&u).Error
	return u, err
}

func (m *Manager) GetUserInformationWithID(id int64) (User, error) {
	var u User
	err := m.handler.Table(user).Where("id = ?", id).Take(&u).Error
	return u, err
}

func (m *Manager) GetUserInformationWithName(name string) (User, error) {
	var u User
	err := m.handler.Table(user).Where("name = ?", name).Take(&u).Error
	return u, err
}

func (m *Manager) GetUserListWithID(id int64) ([]SimpleUser, error) {
	var us []SimpleUser
	err := m.handler.Table(user).Raw("select u.id, u.uuid, u.name, u.nick_name, u.gender, u.avatar, u.email, u.signature from friend as f join user as u on f.friend_id = u.id where f.user_id = ?", id).Scan(&us).Error
	return us, err
}

func (m *Manager) UpdateUserInfoWithID(id int64, u SimpleUser) error {
	return m.handler.Table(user).Where("id = ?", id).Omit("id", "name", "email").Updates(&u).Error
}

func (m *Manager) UpdateUserPasswordWithID(id int64, password string) error {
	return m.handler.Table(user).Where("id = ?", id).Update("password", password).Error
}

func (m *Manager) DeleteUserWithID(id int64) error {
	return m.handler.Table(user).Where("id = ?", id).Delete(&User{}).Error
}

//func (m *Manager) GetUserInfoWithUuid(uuid string) (SimpleUser, error) {
//	var u SimpleUser
//	err := m.handler.Table(user).Where("uuid = ?", uuid).Take(&u).Error
//	return u, err
//}

//func (m *Manager) GetUserListWithUuid(uuid string) ([]SimpleUser, error) {
//	queryUser, err := m.GetUserInfoWithUuid(uuid)
//	if err != nil {
//		return nil, err
//	}
//
//	var us []SimpleUser
//	err = m.handler.Table(user).Raw("select u.id, u.uuid, u.name, u.nick_name, u.gender, u.avatar, u.email, u.signature from friend as f join user as u on f.friend_id = u.id where f.user_id = ?", queryUser.ID).Scan(&us).Error
//	return us, err
//}

//func (m *Manager) GetUserInformationWithUuid(uuid string) (User, error) {
//	var u User
//	err := m.handler.Table(user).Where("uuid = ?", uuid).Take(&u).Error
//	return u, err
//}

//func (m *Manager) UpdateUserInfoWithUuid(uuid string, u SimpleUser) error {
//	return m.handler.Table(user).Where("uuid = ?", uuid).Omit("id", "uuid", "name", "email").Updates(&u).Error
//}

//func (m *Manager) UpdateUserPasswordWithUuid(uuid string, password string) error {
//	return m.handler.Table(user).Where("uuid = ?", uuid).Update("password", password).Error
//}

//func (m *Manager) DeleteUserWithUuid(uuid int64) error {
//	return m.handler.Table(user).Where("uuid = ?", uuid).Delete(&User{}).Error
//}
