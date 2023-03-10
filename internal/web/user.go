package web

import (
	"chat/internal/app/service"
	"chat/internal/pkg/center"
	"context"
	"github.com/kataras/iris/v12"
)

func (m *Manager) RouteUser() {
	m.handler.PartyFunc("/user", func(p iris.Party) {
		p.Get("/email/{email}", m.getEmail)
		p.Get("/info", m.tokener.Serve(), m.getUserInfo)
		p.Post("/register", m.register)
		p.Post("/login", m.login)
		p.Post("/update", m.tokener.Serve(), m.updateInfo)
		p.Post("/change/password", m.tokener.Serve(), m.changePwd)
	})
}

func (m *Manager) getEmail(ctx iris.Context) {
	email := ctx.Params().Get("email")
	conn, err := center.Resolver("user")
	if err != nil {
		m.sendSimpleMessage(ctx, iris.StatusInternalServerError)
		return
	}
	c := service.NewUserServiceClient(conn)

	_, err = c.Code(context.Background(), &service.CodeRequest{
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
	conn, err := center.Resolver("user")
	if err != nil {
		m.sendSimpleMessage(ctx, iris.StatusInternalServerError)
		return
	}
	c := service.NewUserServiceClient(conn)

	_, err = c.Register(context.Background(), &service.UserRegisterRequest{
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
	conn, err := center.Resolver("user")
	if err != nil {
		m.sendSimpleMessage(ctx, iris.StatusInternalServerError)
		return
	}
	c := service.NewUserServiceClient(conn)

	response, err := c.Login(context.Background(), &service.UserLoginRequest{
		Name:     msg.Name,
		Password: msg.Password,
	})
	if err != nil {
		m.sendErrorMessage(ctx, err)
		return
	}
	m.sendJson(ctx, iris.StatusOK, map[string]any{
		"msg":   "请求成功",
		"token": response.Token,
	})
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
	uuid := m.tokener.GetUUID(ctx)
	conn, err := center.Resolver("user")
	if err != nil {
		m.sendSimpleMessage(ctx, iris.StatusInternalServerError)
		return
	}
	c := service.NewUserServiceClient(conn)

	response, err := c.GetUserInfo(context.Background(), &service.GetUserInfoRequest{
		Uuid: uuid,
	})
	if err != nil {
		m.sendErrorMessage(ctx, err)
		return
	}
	m.sendJson(ctx, iris.StatusOK, map[string]any{
		"data": response.User,
	})
}

type updateInfoMessage struct {
	NickName  string `json:"nick_name"`
	Gender    string `json:"gender"`
	Avatar    string `json:"avatar"`
	Signature string `json:"signature"`
}

func (m *Manager) updateInfo(ctx iris.Context) {
	uuid := m.tokener.GetUUID(ctx)
	var msg updateInfoMessage
	if err := ctx.ReadJSON(&msg); err != nil {
		m.sendSimpleMessage(ctx, iris.StatusBadRequest, err)
		return
	}
	conn, err := center.Resolver("user")
	if err != nil {
		m.sendSimpleMessage(ctx, iris.StatusInternalServerError)
		return
	}
	c := service.NewUserServiceClient(conn)

	_, err = c.UpdateUserInfo(context.Background(), &service.UpdateUserInfoRequest{
		Uuid: uuid,
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
	uuid := m.tokener.GetUUID(ctx)
	var msg changePwdMessage
	if err := ctx.ReadJSON(&msg); err != nil {
		m.sendSimpleMessage(ctx, iris.StatusBadRequest, err)
		return
	}
	conn, err := center.Resolver("user")
	if err != nil {
		m.sendSimpleMessage(ctx, iris.StatusInternalServerError)
		return
	}
	c := service.NewUserServiceClient(conn)

	_, err = c.ChangePassword(context.Background(), &service.UserChangePasswordRequest{
		Uuid:   uuid,
		OldPwd: msg.OldPwd,
		NewPwd: msg.NewPwd,
	})
	if err != nil {
		m.sendErrorMessage(ctx, err)
		return
	}
	m.sendSimpleMessage(ctx, iris.StatusOK)
}
