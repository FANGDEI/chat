package web

import (
	mylog "chat/internal/pkg/logger"
	"github.com/gorilla/websocket"
	"github.com/kataras/iris/v12"
	"net/http"
)

func (m *Manager) RouteChat() {
	m.handler.PartyFunc("/chat", func(p iris.Party) {
		//p.Use(m.tokener.Serve())
		p.Get("/ws", m.ws)
		p.Get("/ws/test", m.wsTest)
	})
}

type Client struct {
	Uuid string
	// Name string
	Conn  *websocket.Conn
	Close chan struct{}
}

var defaultLogger = mylog.GetDefaultLoggerManager()

func (c *Client) Read() {
	defer func() {
		close(c.Close)
		c.Conn.Close()
	}()
	for {
		_, message, err := c.Conn.ReadMessage()
		if err != nil {
			defaultLogger.Info("client close conn")
			break
		}
		c.Conn.WriteMessage(websocket.TextMessage, message)
	}
}

func (c *Client) Write() {
	defer func() {
		c.Conn.Close()
	}()
	for {
	}
}

func (m *Manager) ws(ctx iris.Context) {
	m.logger.Info("websocket connection " + m.tokener.GetUUID(ctx))
	var upGrader = websocket.Upgrader{
		CheckOrigin: func(r *http.Request) bool {
			return true
		},
		// Subprotocols: []string{ctx.Request().Header.Get("Sec-WebSocket-Protocol")},
	}
	conn, err := upGrader.Upgrade(ctx.ResponseWriter(), ctx.Request(), nil)
	if err != nil {
		m.sendSimpleMessage(ctx, iris.StatusInternalServerError, err)
		return
	}
	uuid := m.tokener.GetUUID(ctx)
	client := &Client{
		Uuid:  uuid,
		Conn:  conn,
		Close: make(chan struct{}),
	}

	go client.Read()
	go client.Write()
}

func (m *Manager) wsTest(ctx iris.Context) {
	var upGrader = websocket.Upgrader{
		CheckOrigin: func(r *http.Request) bool {
			return true
		},
		// Subprotocols: []string{ctx.Request().Header.Get("Sec-WebSocket-Protocol")},
	}
	conn, err := upGrader.Upgrade(ctx.ResponseWriter(), ctx.Request(), nil)
	if err != nil {
		m.sendSimpleMessage(ctx, iris.StatusInternalServerError, err)
		return
	}
	client := &Client{
		Uuid:  "uuid",
		Conn:  conn,
		Close: make(chan struct{}),
	}

	go client.Read()
	go client.Write()
}
