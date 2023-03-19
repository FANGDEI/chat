package web

import (
	"chat/internal/app/service"
	"chat/internal/pkg/center"
	"log"
)

var (
	groupClient service.GroupServiceClient
)

func init() {
	conn, err := center.Resolver("group")
	if err != nil {
		defaultLogger.Error("error: failed to resolver group service")
		log.Fatalln(err)
	}
	groupClient = service.NewGroupServiceClient(conn)
}
