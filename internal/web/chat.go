package web

import (
	"chat/internal/app/service"
	"context"
	"encoding/json"
	"fmt"
	"log"
	"math/rand"
	"net/http"
	"time"

	"github.com/gorilla/websocket"
	"github.com/kataras/iris/v12"
)

func (m *Manager) RouteChat() {
	m.handler.PartyFunc("/chat", func(p iris.Party) {
		//p.Use(m.tokener.Serve())
		p.Get("/ws", m.ws)
		p.Get("/ws/test", m.wsTest)
	})
}

type Message struct {
	From        string `json:"from"`
	To          string `json:"to"`
	Content     string `json:"content"`
	MessageType int64  `json:"message_type"` // 1 群聊 2 私聊
	ContentType int64  `json:"content_type"`
	Time        string `json:"time"`
}

type Client struct {
	Uuid string
	// Name string
	Conn  *websocket.Conn
	Close chan struct{}
}

func (c *Client) Read() {
	defer func() {
		MyHub.UnRegister <- c
		c.Conn.Close()
	}()
	for {
		// json
		// {"to": "uuid", "content": "hello im", "message_type": 1, "content_type": 1}
		_, message, err := c.Conn.ReadMessage()
		if err != nil {
			defaultLogger.Info("client close conn")
			break
		}
		// 反序列化收到的 json 数据
		var req service.SendRequest
		if err := json.Unmarshal(message, &req); err != nil {
			defaultLogger.Error(err.Error())
			c.Conn.WriteMessage(websocket.TextMessage, []byte(err.Error()))
			break
		}
		req.From = c.Uuid
		if req.To == c.Uuid {
			defaultLogger.Error("error: send message self")
			c.Conn.WriteMessage(websocket.TextMessage, []byte("send message to yourself"))
			break
		}
		if req.To == "" {
			defaultLogger.Error("error: wrong target uuid")
			c.Conn.WriteMessage(websocket.TextMessage, []byte("error: wrong target uuid"))
			break
		}
		// chatClient 发送消息到对应的 Redis 消息队列中
		if _, err := chatClient.Send(context.Background(), &req); err != nil {
			defaultLogger.Error(err.Error())
			c.Conn.WriteMessage(websocket.TextMessage, []byte(err.Error()))
		}
	}
}

func (c *Client) Write() {
	defer func() {
		c.Conn.Close()
	}()
	for {
		ctx, cancel := context.WithDeadline(context.Background(), time.Now().Add(time.Hour))
		go func() {
			// 客户端关闭websocket取消请求
			<-c.Close
			cancel()
		}()
		// response.Msg 中返回收到的所有消息的 json 字符串
		response, err := chatClient.Get(ctx, &service.GetRequest{
			Uuid: c.Uuid,
		})
		if err != nil {
			defaultLogger.Error(err.Error())
			c.Conn.WriteMessage(websocket.TextMessage, []byte(err.Error()))
			return
		}
		// 将消息 json 串发送给前端解析
		for _, msg := range response.Msg {
			c.Conn.WriteMessage(websocket.TextMessage, []byte(msg))
		}
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
	log.Println(uuid)
	client := &Client{
		Uuid:  uuid,
		Conn:  conn,
		Close: make(chan struct{}),
	}

	// 向 Hub 注册 websocket 连接
	// 用于查找当前在线用户
	MyHub.Register <- client
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
	uuid := fmt.Sprintf("%06d", rand.Intn(100000))
	log.Println(uuid)
	client := &Client{
		Uuid:  uuid,
		Conn:  conn,
		Close: make(chan struct{}),
	}

	MyHub.Register <- client
	go client.Read()
	go client.Write()
}
