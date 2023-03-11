package web

import (
	"chat/internal/pkg/logger"
	"go.uber.org/zap"
)

type Server struct {
	Clients    map[string]*Client
	Broadcast  chan []byte
	Register   chan *Client
	UnRegister chan *Client
	logger     *logger.Manager
}

func NewServer() *Server {
	return &Server{
		Clients:    make(map[string]*Client),
		Broadcast:  make(chan []byte),
		Register:   make(chan *Client),
		UnRegister: make(chan *Client),
		logger:     logger.GetDefaultLoggerManager(),
	}
}

func (s *Server) Run() {
	s.logger.Info("Run server", zap.Any("start server", "start server ..."))
	for {
		select {
		case conn := <-s.Register:
			// 用户获得websocket连接后向Server进行注册
			s.logger.Info(conn.Name + " connect to websocket")
			s.Clients[conn.Name] = conn
		case conn := <-s.UnRegister:
			// 用户释放连接
			s.logger.Info(conn.Name + " give up the websocket")
			if _, ok := s.Clients[conn.Name]; ok {
				close(conn.Send)
				delete(s.Clients, conn.Name)
			}
			//case msg := <-s.Broadcast:

		}
	}
}

type Msg struct {
	MessageType int    `json:"message_type"`
	ContentType int    `json:"content_type"`
	Content     string `json:"content"`
	From        string `json:"from"`
	Time        string `json:"time"`
}
