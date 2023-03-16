package localer

import "testing"

func TestCreateFriend(t *testing.T) {
	t.Log(defaultLocalerManager.CreateFriend(1, 2))
}
