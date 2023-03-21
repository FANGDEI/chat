package web

import (
	"chat/internal/app/service"
	"chat/internal/pkg/cacher"
	"context"
	"encoding/json"
	"log"
	"math/rand"
	"net/http"
	"time"

	"github.com/gorilla/websocket"
	"github.com/kataras/iris/v12"
)

func (m *Manager) RouteChat() {
	m.handler.PartyFunc("/chat", func(p iris.Party) {
		p.Use(m.tokener.Serve())
		p.Get("/ws", m.ws)
		p.Get("/ws/test", m.wsTest)
		p.Post("/user/history", m.getUserHistory)
		p.Post("/group/history", m.getGroupHistory)
	})
}

// type Message struct {
// 	From        string `json:"from"`
// 	To          string `json:"to"`
// 	Content     string `json:"content"`
// 	MessageType int64  `json:"message_type"` // 1 群聊 2 私聊
// 	ContentType int64  `json:"content_type"`
// 	Time        string `json:"time"`
// }

var (
	redis = cacher.GetDefaultCacherManager()
)

type Client struct {
	UserID int64
	Conn   *websocket.Conn
	Close  chan struct{}
}

func (c *Client) Read() {
	defer func() {
		MyHub.UnRegister <- c
		c.Conn.Close()
	}()
	for {
		// json
		// {"to": 1, "content": "hello im", "message_type": 1, "content_type": 1}
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
		req.From = c.UserID
		if req.To == c.UserID {
			defaultLogger.Error("error: send message self")
			c.Conn.WriteMessage(websocket.TextMessage, []byte("send message to yourself"))
			break
		}
		if req.To == 0 {
			defaultLogger.Error("error: wrong target uuid")
			c.Conn.WriteMessage(websocket.TextMessage, []byte("error: wrong target uuid"))
			break
		}
		// chatClient 发送消息到对应的 Redis 消息队列中
		if _, err := chatClient.Send(context.Background(), &req); err != nil {
			defaultLogger.Error(err.Error())
			c.Conn.WriteMessage(websocket.TextMessage, []byte(err.Error()))
			break
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
			Id: c.UserID,
		})
		if err != nil {
			defaultLogger.Error(err.Error())
			c.Conn.WriteMessage(websocket.TextMessage, []byte(err.Error()))
			return
		}
		// 将消息 json 串发送给前端解析
		for _, msg := range response.Msg {
			if err := c.Conn.WriteMessage(websocket.TextMessage, []byte(msg)); err != nil {
				reErr := redis.RewriteAndPop(c.UserID, response.Msg) // conn closed, rewrite msgs to redis and pop the message history
				if reErr != nil {
					defaultLogger.Error(err.Error())
				}
				return
			}
		}
	}
}

func (m *Manager) ws(ctx iris.Context) {
	m.logger.Info("websocket connection")
	var upGrader = websocket.Upgrader{
		CheckOrigin: func(r *http.Request) bool {
			return true
		},
		Subprotocols: []string{ctx.Request().Header.Get("Sec-WebSocket-Protocol")},
	}
	conn, err := upGrader.Upgrade(ctx.ResponseWriter(), ctx.Request(), nil)
	if err != nil {
		m.sendSimpleMessage(ctx, iris.StatusInternalServerError, err)
		return
	}
	id := m.tokener.GetID(ctx)
	client := &Client{
		UserID: id,
		Conn:   conn,
		Close:  make(chan struct{}),
	}

	// 向 Hub 注册 websocket 连接
	MyHub.Register <- client
	go client.Read()
	go client.Write()
}

type getUserHistoryMessage struct {
	OtherID    int64 `json:"other_id"`
	Offset     int64 `json:"offset"`
	Limit      int64 `json:"limit"`
	Pagination bool  `json:"pagination"`
}

func (m *Manager) getUserHistory(ctx iris.Context) {
	var msg getUserHistoryMessage
	if err := ctx.ReadJSON(&msg); err != nil {
		m.sendSimpleMessage(ctx, iris.StatusBadRequest, err)
		return
	}
	response, err := chatClient.GetUserHistory(context.Background(), &service.GetUserHistoryRequest{
		UserId:     m.tokener.GetID(ctx),
		OtherId:    msg.OtherID,
		Offset:     msg.Offset,
		Limit:      msg.Limit,
		Pagination: msg.Pagination,
	})
	if err != nil {
		m.sendErrorMessage(ctx, err)
		return
	}
	m.sendGRPCMessage(ctx, iris.StatusOK, response, service.GetUserHistoryResponse{})
}

type getGroupHistoryMessage struct {
	GroupID    int64 `json:"group_id"`
	Offset     int64 `json:"offset"`
	Limit      int64 `json:"limit"`
	Pagination bool  `json:"pagination"`
}

func (m *Manager) getGroupHistory(ctx iris.Context) {
	var msg getGroupHistoryMessage
	if err := ctx.ReadJSON(&msg); err != nil {
		m.sendSimpleMessage(ctx, iris.StatusBadRequest, err)
		return
	}
	response, err := chatClient.GetGroupHistory(context.Background(), &service.GetGroupHistoryRequest{
		GroupId:    msg.GroupID,
		Offset:     msg.Offset,
		Limit:      msg.Limit,
		Pagination: msg.Pagination,
	})
	if err != nil {
		m.sendErrorMessage(ctx, err)
		return
	}
	m.sendGRPCMessage(ctx, iris.StatusOK, response, service.GetGroupHistoryResponse{})
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
	id := rand.Intn(100000)
	log.Println(id)
	client := &Client{
		UserID: int64(id),
		Conn:   conn,
		Close:  make(chan struct{}),
	}

	MyHub.Register <- client
	go client.Read()
	go client.Write()
}
