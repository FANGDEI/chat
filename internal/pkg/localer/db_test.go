package localer

import "testing"

func TestCreateFriend(t *testing.T) {
	t.Log(defaultLocalerManager.CreateFriend(1, 2))
}

func TestGetUserList(t *testing.T) {
	t.Log(defaultLocalerManager.GetUserListWithID(4))
}

func TestGetUserGroupList(t *testing.T) {
	list, _ := defaultLocalerManager.GetUserGroupListWithID(4)
	t.Log(list)
}
