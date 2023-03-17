package web

import (
	"chat/internal/app/service"
	"context"

	"github.com/kataras/iris/v12"
)

func (m *Manager) RouteUser() {
	m.handler.PartyFunc("/user", func(p iris.Party) {
		p.Get("/email/{email}", m.getEmail)
		p.Get("/info", m.tokener.Serve(), m.getUserInfo)
		p.Get("/info/{name}", m.tokener.Serve(), m.getOtherUserInfo)
		p.Get("/friends", m.tokener.Serve(), m.friends)
		p.Post("/register", m.register)
		p.Post("/login", m.login)
		p.Post("/update", m.tokener.Serve(), m.updateInfo)
		p.Post("/change/password", m.tokener.Serve(), m.changePwd)
	})
}

func (m *Manager) getEmail(ctx iris.Context) {
	email := ctx.Params().Get("email")
	_, err := userClient.Code(context.Background(), &service.CodeRequest{
		Email: email,
	})
	if err != nil {
		m.sendErrorMessage(ctx, err)
		return
	}
	m.sendSimpleMessage(ctx, iris.StatusOK)
}

type registerMessage struct {
	Name     string `json:"name"`
	Password string `json:"password"`
	Email    string `json:"email"`
	Code     string `json:"code"`
}

func (m *Manager) register(ctx iris.Context) {
	var msg registerMessage
	if err := ctx.ReadJSON(&msg); err != nil {
		m.sendSimpleMessage(ctx, iris.StatusBadRequest, err)
		return
	}
	_, err := userClient.Register(context.Background(), &service.UserRegisterRequest{
		Name:     msg.Name,
		Password: msg.Password,
		Email:    msg.Email,
		Code:     msg.Code,
	})
	if err != nil {
		m.sendErrorMessage(ctx, err)
		return
	}
	m.sendSimpleMessage(ctx, iris.StatusOK)
}

type loginMessage struct {
	Name     string `json:"name"`
	Password string `json:"password"`
}

func (m *Manager) login(ctx iris.Context) {
	var msg loginMessage
	if err := ctx.ReadJSON(&msg); err != nil {
		m.sendSimpleMessage(ctx, iris.StatusBadRequest, err)
		return
	}
	response, err := userClient.Login(context.Background(), &service.UserLoginRequest{
		Name:     msg.Name,
		Password: msg.Password,
	})
	if err != nil {
		m.sendErrorMessage(ctx, err)
		return
	}
	// m.sendJson(ctx, iris.StatusOK, map[string]any{
	// 	"token": response.Token,
	// })
	m.sendGRPCMessage(ctx, iris.StatusOK, response, service.UserLoginResponse{})
}

//type getUserInfoMessage struct {
//	Uuid string `json:"uuid"`
//}

func (m *Manager) getUserInfo(ctx iris.Context) {
	//var msg getUserInfoMessage
	//if err := ctx.ReadJSON(&msg); err != nil {
	//	m.sendSimpleMessage(ctx, iris.StatusBadRequest, err)
	//	return
	//}
	id := m.tokener.GetID(ctx)
	response, err := userClient.GetUserInfo(context.Background(), &service.GetUserInfoRequest{
		Id: id,
	})
	if err != nil {
		m.sendErrorMessage(ctx, err)
		return
	}
	//m.sendJson(ctx, iris.StatusOK, map[string]any{
	//	"data": response.User,
	//})
	m.sendGRPCMessage(ctx, iris.StatusOK, response, service.GetUserInfoResponse{})
}

func (m *Manager) getOtherUserInfo(ctx iris.Context) {
	name := ctx.Params().Get("name")
	response, err := userClient.GetOtherUserInfo(context.Background(), &service.GetOtherUserInfoRequest{
		Name: name,
	})
	if err != nil {
		m.sendErrorMessage(ctx, err)
		return
	}
	m.sendGRPCMessage(ctx, iris.StatusOK, response, service.GetOtherUserInfoResponse{})
}

type updateInfoMessage struct {
	NickName  string `json:"nick_name"`
	Gender    string `json:"gender"`
	Avatar    string `json:"avatar"`
	Signature string `json:"signature"`
}

func (m *Manager) updateInfo(ctx iris.Context) {
	id := m.tokener.GetID(ctx)
	var msg updateInfoMessage
	if err := ctx.ReadJSON(&msg); err != nil {
		m.sendSimpleMessage(ctx, iris.StatusBadRequest, err)
		return
	}
	_, err := userClient.UpdateUserInfo(context.Background(), &service.UpdateUserInfoRequest{
		Id: id,
		User: &service.SimpleUser{
			Nickname:  msg.NickName,
			Gender:    msg.Gender,
			Avatar:    msg.Avatar,
			Signature: msg.Signature,
		},
	})
	if err != nil {
		m.sendErrorMessage(ctx, err)
		return
	}
	m.sendSimpleMessage(ctx, iris.StatusOK)
}

type changePwdMessage struct {
	OldPwd string `json:"old_pwd"`
	NewPwd string `json:"new_pwd"`
}

func (m *Manager) changePwd(ctx iris.Context) {
	id := m.tokener.GetID(ctx)
	var msg changePwdMessage
	if err := ctx.ReadJSON(&msg); err != nil {
		m.sendSimpleMessage(ctx, iris.StatusBadRequest, err)
		return
	}
	_, err := userClient.ChangePassword(context.Background(), &service.UserChangePasswordRequest{
		Id:     id,
		OldPwd: msg.OldPwd,
		NewPwd: msg.NewPwd,
	})
	if err != nil {
		m.sendErrorMessage(ctx, err)
		return
	}
	m.sendSimpleMessage(ctx, iris.StatusOK)
}

func (m *Manager) friends(ctx iris.Context) {
	id := m.tokener.GetID(ctx)
	response, err := userClient.GetUserList(context.Background(), &service.GetUserListRequest{Id: id})
	if err != nil {
		m.sendErrorMessage(ctx, err)
		return
	}
	m.sendGRPCMessage(ctx, iris.StatusOK, response, service.GetUserInfoResponse{})
}
