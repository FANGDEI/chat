package web

import (
	"github.com/kataras/iris/v12"
)

func (m *Manager) RouteFriend() {
	m.handler.PartyFunc("/friend", func(p iris.Party) {
		p.Post("/add")
		p.Post("/del")
	})
}
