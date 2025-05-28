package main

import (
	"context"
	"log"
	"net/http"

	"github.com/TimNikolaev/drag-chat/internal/delivery/rest"
	"github.com/TimNikolaev/drag-chat/internal/delivery/ws"
	"github.com/TimNikolaev/drag-chat/internal/repository"
	"github.com/TimNikolaev/drag-chat/internal/repository/postgres"
	"github.com/TimNikolaev/drag-chat/internal/server"
	"github.com/TimNikolaev/drag-chat/internal/service"
	"github.com/TimNikolaev/drag-chat/pkg/event/redis"
	"github.com/gorilla/websocket"
)

/*
Разобраться в context.
Разобраться в Redis list.
Разобратся с обработкой ошибок в горутине.
*/

func main() {
	db, err := postgres.New("")
	if err != nil {
		log.Fatalf("failed to initialization db: %s\n", err.Error())
	}
	redisClient, err := redis.InitRedis(context.Background(), "1652")
	if err != nil {
		log.Fatalf("fail to initialization redis %s\n", err.Error())
	}

	repository := repository.New(db)

	service := service.New(repository, redisClient)

	restHandler := rest.NewHandler(service)

	upgrader := websocket.Upgrader{
		CheckOrigin: func(r *http.Request) bool { return true },
	}
	wsHandler := ws.NewWSHandler(service, &upgrader)

	srv := new(server.Server)

	go func() {
		if err := srv.Run("8081", wsHandler.InitConnectRout()); err != nil {
			log.Fatalf("error occurred while running WS server: %s\n", err.Error())
		}
	}()

	if err := srv.Run("8080", restHandler.InitRouts()); err != nil {
		log.Fatalf("error occurred while running REST server: %s\n", err.Error())
	}

}
