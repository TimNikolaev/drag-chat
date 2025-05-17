package main

import (
	"log"

	dragchat "github.com/TimNikolaev/drag-chat"
	"github.com/TimNikolaev/drag-chat/internal/handler"
)

func main() {

	handler := handler.New()

	srv := new(dragchat.Server)

	if err := srv.Run("8000", handler.InitRouts()); err != nil {
		log.Fatalf("error occurred while running http server: %s", err.Error())
	}

}
