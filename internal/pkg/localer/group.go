package localer

type Group struct {
	ID     int64  `json:"id"`
	UserID int64  `json:"user_id"`
	Name   string `json:"name"`
	Avatar string `json:"avatar"`
	Notice string `json:"notice"`
}

var group = "group"

func (m *Manager) CreateGroup(g Group) error {
	err := m.execTx(func(m *Manager) error {
		err := m.handler.Table(group).Create(&g).Error
		if err != nil {
			return err
		}
		err = m.CreateGroupMember(GroupMember{
			UserID:  g.UserID,
			GroupID: g.ID,
		})
		if err != nil {
			return err
		}
		return nil
	})
	return err
}

func (m *Manager) DeleteGroupWithID(id int64) error {
	err := m.execTx(func(m *Manager) error {
		err := m.handler.Table(group).Where("id = ?", id).Delete(&Group{}).Error
		if err != nil {
			return err
		}
		err = m.DeleteGroupMemberWithGroupID(id)
		if err != nil {
			return err
		}
		err = m.DeleteGroupApplyWithGroupID(id)
		if err != nil {
			return err
		}
		return nil
	})
	return err
}

func (m *Manager) GetGroupInfoWithID(id int64) (Group, error) {
	var g Group
	err := m.handler.Table(group).Where("id = ?", id).Take(&g).Error
	return g, err
}

func (m *Manager) GetGroupInfoWithName(name string) (Group, error) {
	var g Group
	err := m.handler.Table(group).Where("name = ?", name).Take(&g).Error
	return g, err
}

func (m *Manager) UpdateGroupAvatarWithID(id int64, avatar string) error {
	return m.handler.Table(group).Where("id = ?", id).Update("avatar", avatar).Error
}

func (m *Manager) UpdateGroupNoticeWithID(id int64, notice string) error {
	return m.handler.Table(group).Where("id = ?", id).Update("notice", notice).Error
}
