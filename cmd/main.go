package main

import (
	"log"
	"net/http"

	"github.com/TimNikolaev/drag-chat/internal/delivery/rest"
	"github.com/TimNikolaev/drag-chat/internal/delivery/ws"
	"github.com/TimNikolaev/drag-chat/internal/repository"
	"github.com/TimNikolaev/drag-chat/internal/server"
	"github.com/TimNikolaev/drag-chat/internal/service"
	"github.com/gorilla/websocket"
)

func main() {
	repository := repository.New()

	service := service.New(repository)

	restHandler := rest.NewHandler(service)

	upgrader := websocket.Upgrader{
		CheckOrigin: func(r *http.Request) bool { return true },
	}
	wsHandler := ws.NewWSHandler(service, &upgrader)

	srv := new(server.Server)

	go func() {
		if err := srv.Run("8081", wsHandler.InitConnectRout()); err != nil {
			log.Fatalf("error occurred while running WS server: %s", err.Error())
		}
	}()

	if err := srv.Run("8080", restHandler.InitRouts()); err != nil {
		log.Fatalf("error occurred while running REST server: %s", err.Error())
	}

}
