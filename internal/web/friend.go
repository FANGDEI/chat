package web

//
//import (
//	"chat/internal/app/service"
//	"context"
//
//	"github.com/kataras/iris/v12"
//)
//
//func (m *Manager) RouteFriend() {
//	m.handler.PartyFunc("/friend", func(p iris.Party) {
//		p.Use(m.tokener.Serve())
//		p.Post("/{name}", m.addFriend)
//		p.Delete("/{name}", m.delFriend)
//		p.Post("/accept", m.accept)
//	})
//}
//
//func (m *Manager) addFriend(ctx iris.Context) {
//	uuid, name := m.tokener.GetUUID(ctx), ctx.Params().Get("name")
//	_, err := friendClient.AddFriend(context.Background(), &service.AddFriendRequest{
//		Uuid:       uuid,
//		FriendName: name,
//	})
//	if err != nil {
//		m.sendErrorMessage(ctx, err)
//		return
//	}
//	m.sendSimpleMessage(ctx, iris.StatusOK)
//}
//
//func (m *Manager) delFriend(ctx iris.Context) {
//	uuid, name := m.tokener.GetUUID(ctx), ctx.Params().Get("name")
//	_, err := friendClient.DelFriend(context.Background(), &service.DelFriendRequest{
//		Uuid:       uuid,
//		FriendName: name,
//	})
//	if err != nil {
//		m.sendErrorMessage(ctx, err)
//	}
//	m.sendSimpleMessage(ctx, iris.StatusOK)
//}
//
//type acceptMessage struct {
//	Uuid  string `json:"uuid"`
//	Agree int    `json:"agree"`
//}
//
//func (m *Manager) accept(ctx iris.Context) {
//	uuid := m.tokener.GetUUID(ctx)
//	var msg acceptMessage
//	if err := ctx.ReadJSON(&msg); err != nil {
//		m.sendSimpleMessage(ctx, iris.StatusBadRequest, err)
//		return
//	}
//	_, err := friendClient.AcceptFriend(context.Background(), &service.AcceptFriendRequest{
//		SelfUuid: uuid,
//		Uuid:     msg.Uuid,
//		Agree:    msg.Agree == 1,
//	})
//	if err != nil {
//		m.sendErrorMessage(ctx, err)
//		return
//	}
//	m.sendSimpleMessage(ctx, iris.StatusOK)
//}
