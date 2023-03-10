package main

import (
	"chat/internal/web"
	"log"
)

func main() {
	app := web.New()
	if err := app.Run(); err != nil {
		log.Fatalf("failed to server, %v", err)
	}
}
