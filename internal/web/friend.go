package web

import (
	"chat/internal/app/service"
	"context"

	"github.com/kataras/iris/v12"
)

func (m *Manager) RouteFriend() {
	m.handler.PartyFunc("/friend", func(p iris.Party) {
		p.Use(m.tokener.Serve())
		p.Post("/add/{name}", m.addFriend)
		p.Post("/del")
	})
}

func (m *Manager) addFriend(ctx iris.Context) {
	uuid, name := m.tokener.GetUUID(ctx), ctx.Params().Get("name")
	_, err := friendClient.AddFriend(context.Background(), &service.AddFriendRequest{
		Uuid:       uuid,
		FriendName: name,
	})
	if err != nil {
		m.sendErrorMessage(ctx, err)
		return
	}
	m.sendSimpleMessage(ctx, iris.StatusOK)
}
