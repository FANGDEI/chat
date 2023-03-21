package web

import (
	"chat/internal/app/service"
	"context"

	"github.com/kataras/iris/v12"
)

func (m *Manager) RouteFriend() {
	m.handler.PartyFunc("/friend", func(p iris.Party) {
		p.Use(m.tokener.Serve())
		p.Get("/list", m.applyList)
		p.Post("/add", m.addFriend)
		p.Post("/del", m.delFriend)
		p.Post("/agree", m.agree)
	})
}

type addFriendMessage struct {
	FriendID int64  `json:"friend_id"`
	ApplyMsg string `json:"apply_msg"`
}

func (m *Manager) addFriend(ctx iris.Context) {
	userID := m.tokener.GetID(ctx)
	var msg addFriendMessage
	if err := ctx.ReadJSON(&msg); err != nil {
		m.sendSimpleMessage(ctx, iris.StatusBadRequest, err)
		return
	}
	_, err := friendClient.AddFriend(context.Background(), &service.AddFriendRequest{
		UserId:   userID,
		FriendId: msg.FriendID,
		ApplyMsg: msg.ApplyMsg,
	})
	if err != nil {
		m.sendErrorMessage(ctx, err)
		return
	}
	m.sendSimpleMessage(ctx, iris.StatusOK)
}

type delFriendMessage struct {
	FriendID int64 `json:"friend_id"`
}

func (m *Manager) delFriend(ctx iris.Context) {
	var msg delFriendMessage
	if err := ctx.ReadJSON(&msg); err != nil {
		m.sendSimpleMessage(ctx, iris.StatusBadRequest, err)
		return
	}
	_, err := friendClient.DelFriend(context.Background(), &service.DelFriendRequest{
		UserId:   m.tokener.GetID(ctx),
		FriendId: msg.FriendID,
	})
	if err != nil {
		m.sendErrorMessage(ctx, err)
		return
	}
	m.sendSimpleMessage(ctx, iris.StatusOK)
}

type agreeMessage struct {
	FriendID int64 `json:"friend_id"`
	Agree    bool  `json:"agree"`
}

func (m *Manager) agree(ctx iris.Context) {
	var msg agreeMessage
	if err := ctx.ReadJSON(&msg); err != nil {
		m.sendSimpleMessage(ctx, iris.StatusBadRequest, err)
		return
	}
	_, err := friendClient.AgreeApply(context.Background(), &service.AgreeApplyRequest{
		UserId:   m.tokener.GetID(ctx),
		FriendId: msg.FriendID,
		Agree:    msg.Agree,
	})
	if err != nil {
		m.sendErrorMessage(ctx, err)
		return
	}
	m.sendSimpleMessage(ctx, iris.StatusOK)
}

func (m *Manager) applyList(ctx iris.Context) {
	response, err := friendClient.GetFriendApply(context.Background(), &service.GetFriendApplyRequest{
		UserId: m.tokener.GetID(ctx),
	})
	if err != nil {
		m.sendErrorMessage(ctx, err)
		return
	}
	m.sendGRPCMessage(ctx, iris.StatusOK, response, service.GetFriendApplyResponse{})
}
