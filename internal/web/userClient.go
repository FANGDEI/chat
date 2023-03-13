package web

import (
	"chat/internal/app/service"
	"chat/internal/pkg/center"
	"log"
)

var (
	userClient service.UserServiceClient
)

func init() {
	conn, err := center.Resolver("user")
	if err != nil {
		defaultLogger.Error("error: failed to resolver user service")
		log.Fatalln(err)
	}
	userClient = service.NewUserServiceClient(conn)
}
