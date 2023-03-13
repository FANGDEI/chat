package web

import (
	"chat/internal/app/service"
	"chat/internal/pkg/center"
	"log"
)

var (
	friendClient service.FriendServiceClient
)

func init() {
	conn, err := center.Resolver("friend")
	if err != nil {
		defaultLogger.Error("error: failed to resolver friend service")
		log.Fatalln(err)
	}
	friendClient = service.NewFriendServiceClient(conn)
}
