package web

import "github.com/kataras/iris/v12/websocket"

type Client struct {
	Uuid string
	Name string
	Conn *websocket.Conn
	Send chan []byte
}
