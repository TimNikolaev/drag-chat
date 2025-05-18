package main

import (
	"log"
	"net/http"

	"github.com/TimNikolaev/drag-chat/internal/delivery/rest"
	"github.com/TimNikolaev/drag-chat/internal/delivery/ws"
	"github.com/TimNikolaev/drag-chat/internal/server"
	"github.com/gorilla/websocket"
)

func main() {
	upgrader := websocket.Upgrader{
		CheckOrigin: func(r *http.Request) bool { return true },
	}

	ws := ws.NewWSHandler(&upgrader)
	handler := rest.NewHandler()

	srv := new(server.Server)

	go func() {
		if err := srv.Run("8081", ws.InitConnectRout()); err != nil {
			log.Fatalf("error occurred while running WS server: %s", err.Error())
		}
	}()

	if err := srv.Run("8080", handler.InitRouts()); err != nil {
		log.Fatalf("error occurred while running REST server: %s", err.Error())
	}

}
