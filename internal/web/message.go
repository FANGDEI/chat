package web

import (
	"chat/internal/pkg/errno"
	"fmt"
	"github.com/kataras/iris/v12"
	"net/http"
)

func (m *Manager) sendJson(ctx iris.Context, code int, v any) {
	ctx.StatusCode(code)
	ctx.JSON(v)
}

func (m *Manager) sendSimpleMessage(ctx iris.Context, code int, errs ...error) {
	if errs == nil {
		m.sendJson(ctx, code, map[string]string{
			"msg": "请求成功",
		})
		return
	}
	for _, err := range errs {
		m.logger.Info(fmt.Sprintf("[%s] CODE(%d) ERROR : %+v\n", ctx.Path(), code, err))
	}
	var msg string
	switch code {
	case iris.StatusBadRequest:
		msg = "请求解析失败"
	case iris.StatusInternalServerError:
		msg = "服务器内部错误，请联系管理员"
	case iris.StatusOK:
		msg = "请求成功"
	case iris.StatusPreconditionFailed:
		msg = "预处理失败"
	case iris.StatusForbidden:
		msg = "权限认证失败"
	default:
		msg = "未知错误"
	}
	m.sendJson(ctx, code, map[string]string{
		"msg": msg,
	})
}

func (m *Manager) sendMessage(ctx iris.Context, code int, msg string, errs ...error) {
	if errs == nil {
		m.sendJson(ctx, code, map[string]string{
			"msg": msg,
		})
		return
	}
	for _, err := range errs {
		m.logger.Info(fmt.Sprintf("[%s] CODE(%d) ERROR : %+v\n", ctx.Path(), code, err))
	}
	m.sendJson(ctx, code, map[string]string{
		"msg": msg,
	})
}

func (m *Manager) sendErrorMessage(ctx iris.Context, err error) {
	code, msg := errno.ParseErr(err)
	m.logger.Error(msg)

	var responseCode int
	switch {
	case code > 20000:
		responseCode = http.StatusBadRequest
	default:
		responseCode = http.StatusInternalServerError
	}
	m.sendSimpleMessage(ctx, responseCode, err)
}
