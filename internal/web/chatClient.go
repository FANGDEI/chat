package web

import (
	"chat/internal/app/service"
	"chat/internal/pkg/center"
	"log"
)

var (
	chatClient service.ChatServiceClient
)

func init() {
	conn, err := center.Resolver("chat")
	if err != nil {
		defaultLogger.Error("error: failed to resolver chat service")
		log.Fatalln(err)
	}
	chatClient = service.NewChatServiceClient(conn)
}
