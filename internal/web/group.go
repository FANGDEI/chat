package web

import (
	"chat/internal/app/service"
	"context"

	"github.com/kataras/iris/v12"
)

func (m *Manager) RouteGroup() {
	m.handler.PartyFunc("/group", func(p iris.Party) {
		p.Use(m.tokener.Serve())
		p.Get("/list", m.groupApplyList)
		p.Post("/create", m.createGroup)
		p.Post("/info", m.getGroupInfo)
		p.Post("/delete", m.deleteGroup)
		p.Post("/add", m.addGroup)
		p.Post("/exit", m.exitGroup)
		p.Post("/accept", m.accept)
		p.Post("/update/avatar", m.updateGroupAvatar)
		p.Post("/update/notice", m.updateGroupNotice)
	})
}

type createGroupMessage struct {
	Name string `json:"name"`
}

func (m *Manager) createGroup(ctx iris.Context) {
	var msg createGroupMessage
	if err := ctx.ReadJSON(&msg); err != nil {
		m.sendSimpleMessage(ctx, iris.StatusBadRequest, err)
		return
	}
	_, err := groupClient.Create(context.Background(), &service.CreateRequest{
		UserId:    m.tokener.GetID(ctx),
		GroupName: msg.Name,
	})
	if err != nil {
		m.sendErrorMessage(ctx, err)
		return
	}
	m.sendSimpleMessage(ctx, iris.StatusOK)
}

type deleteGroupMessage struct {
	GroupID int64 `json:"group_id"`
}

func (m *Manager) deleteGroup(ctx iris.Context) {
	var msg deleteGroupMessage
	if err := ctx.ReadJSON(&msg); err != nil {
		m.sendSimpleMessage(ctx, iris.StatusBadRequest, err)
		return
	}
	_, err := groupClient.Delete(context.Background(), &service.DeleteRequest{
		UserId:  m.tokener.GetID(ctx),
		GroupId: msg.GroupID,
	})
	if err != nil {
		m.sendErrorMessage(ctx, err)
		return
	}
	m.sendSimpleMessage(ctx, iris.StatusOK)
}

type getGroupInfoMessage struct {
	GroupName string `json:"group_name"`
}

func (m *Manager) getGroupInfo(ctx iris.Context) {
	var msg getGroupInfoMessage
	if err := ctx.ReadJSON(&msg); err != nil {
		m.sendSimpleMessage(ctx, iris.StatusBadRequest, err)
		return
	}
	response, err := groupClient.GetGroupInfo(context.Background(), &service.GetGroupInfoRequest{
		GroupName: msg.GroupName,
	})
	if err != nil {
		m.sendErrorMessage(ctx, err)
		return
	}
	m.sendGRPCMessage(ctx, iris.StatusOK, response, service.GetGroupInfoResponse{})
}

type updateGroupAvatarMessage struct {
	GroupID int64  `json:"group_id"`
	Avatar  string `json:"avatar"`
}

func (m *Manager) updateGroupAvatar(ctx iris.Context) {
	var msg updateGroupAvatarMessage
	if err := ctx.ReadJSON(&msg); err != nil {
		m.sendSimpleMessage(ctx, iris.StatusBadRequest, err)
		return
	}
	_, err := groupClient.UpdateGroupAvatar(context.Background(), &service.UpdateGroupAvatarRequest{
		UserId:  m.tokener.GetID(ctx),
		GroupId: msg.GroupID,
		Avatar:  msg.Avatar,
	})
	if err != nil {
		m.sendErrorMessage(ctx, err)
		return
	}
	m.sendSimpleMessage(ctx, iris.StatusOK)
}

type updateGroupNoticeMessage struct {
	GroupID int64  `json:"group_id"`
	Notice  string `json:"notice"`
}

func (m *Manager) updateGroupNotice(ctx iris.Context) {
	var msg updateGroupNoticeMessage
	if err := ctx.ReadJSON(&msg); err != nil {
		m.sendSimpleMessage(ctx, iris.StatusBadRequest, err)
		return
	}
	_, err := groupClient.UpdateGroupNotice(context.Background(), &service.UpdateGroupNoticeRequest{
		UserId:  m.tokener.GetID(ctx),
		GroupId: msg.GroupID,
		Notice:  msg.Notice,
	})
	if err != nil {
		m.sendErrorMessage(ctx, err)
		return
	}
	m.sendSimpleMessage(ctx, iris.StatusOK)
}

type addGroupMessage struct {
	GroupID  int64  `json:"group_id"`
	ApplyMsg string `json:"apply_msg"`
}

func (m *Manager) addGroup(ctx iris.Context) {
	var msg addGroupMessage
	if err := ctx.ReadJSON(&msg); err != nil {
		m.sendSimpleMessage(ctx, iris.StatusBadRequest, err)
		return
	}
	_, err := groupClient.AddGroup(context.Background(), &service.AddGroupRequest{
		UserId:   m.tokener.GetID(ctx),
		GroupId:  msg.GroupID,
		ApplyMsg: msg.ApplyMsg,
	})
	if err != nil {
		m.sendErrorMessage(ctx, err)
		return
	}
	m.sendSimpleMessage(ctx, iris.StatusOK)
}

type acceptApplyMessage struct {
	UserID  int64 `json:"user_id"`
	GroupID int64 `json:"group_id"`
	Accept  bool  `json:"accept"`
}

func (m *Manager) accept(ctx iris.Context) {
	var msg acceptApplyMessage
	if err := ctx.ReadJSON(&msg); err != nil {
		m.sendSimpleMessage(ctx, iris.StatusBadRequest, err)
		return
	}
	_, err := groupClient.AcceptApply(context.Background(), &service.AcceptApplyRequest{
		UserId:  msg.UserID,
		GroupId: msg.GroupID,
		Accept:  msg.Accept,
	})
	if err != nil {
		m.sendErrorMessage(ctx, err)
		return
	}
	m.sendSimpleMessage(ctx, iris.StatusOK)
}

type exitGroupMessage struct {
	GroupID int64 `json:"group_id"`
}

func (m *Manager) exitGroup(ctx iris.Context) {
	var msg exitGroupMessage
	if err := ctx.ReadJSON(&msg); err != nil {
		m.sendSimpleMessage(ctx, iris.StatusBadRequest, err)
		return
	}
	_, err := groupClient.ExitGroup(context.Background(), &service.ExitGroupRequest{
		UserId:  m.tokener.GetID(ctx),
		GroupId: msg.GroupID,
	})
	if err != nil {
		m.sendErrorMessage(ctx, err)
		return
	}
	m.sendSimpleMessage(ctx, iris.StatusOK)
}

func (m *Manager) groupApplyList(ctx iris.Context) {
	response, err := groupClient.GetGroupApply(context.Background(), &service.GetGroupApplyRequest{
		UserId: m.tokener.GetID(ctx),
	})
	if err != nil {
		m.sendErrorMessage(ctx, err)
		return
	}
	m.sendGRPCMessage(ctx, iris.StatusOK, response, service.GetGroupApplyResponse{})
}
