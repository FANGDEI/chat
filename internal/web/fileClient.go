package web

import (
	"chat/internal/app/service"
	"chat/internal/pkg/center"
	"log"
)

var (
	fileClient service.FileServiceClient
)

func init() {
	conn, err := center.Resolver("file")
	if err != nil {
		defaultLogger.Error("error: failed to resolver file service")
		log.Fatalln(err)
	}
	fileClient = service.NewFileServiceClient(conn)
}
