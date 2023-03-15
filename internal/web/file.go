package web

import (
	"chat/internal/app/service"
	"context"
	"io/ioutil"

	"github.com/kataras/iris/v12"
)

func (m *Manager) RouteFile() {
	m.handler.PartyFunc("/file", func(p iris.Party) {
		p.Use(m.tokener.Serve())
		p.Post("/upload", m.upload)
	})
}

func (m *Manager) upload(ctx iris.Context) {
	suffix := ctx.URLParam("suffix")
	_, fileHeader, err := ctx.FormFile("file")
	if err != nil {
		m.sendSimpleMessage(ctx, iris.StatusInternalServerError, err)
		return
	}
	file, err := fileHeader.Open()
	if err != nil {
		m.sendSimpleMessage(ctx, iris.StatusInternalServerError, err)
		return
	}
	defer file.Close()
	data, err := ioutil.ReadAll(file)
	if err != nil {
		m.sendSimpleMessage(ctx, iris.StatusInternalServerError, err)
		return
	}
	response, err := fileClient.Upload(context.Background(), &service.UploadRequest{
		Suffix: suffix,
		Data:   data,
	})
	if err != nil {
		m.sendErrorMessage(ctx, err)
		return
	}
	m.sendGRPCMessage(ctx, iris.StatusOK, response, service.UploadResponse{})
}
