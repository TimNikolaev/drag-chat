package main

import (
	"context"
	"log"

	"github.com/TimNikolaev/drag-chat/internal/config"
	"github.com/TimNikolaev/drag-chat/internal/delivery/http"
	"github.com/TimNikolaev/drag-chat/internal/repository"
	"github.com/TimNikolaev/drag-chat/internal/repository/postgres"
	"github.com/TimNikolaev/drag-chat/internal/server"
	"github.com/TimNikolaev/drag-chat/internal/service"
	"github.com/TimNikolaev/drag-chat/pkg/event/redis"
	_ "github.com/lib/pq"
)

/*
Разобраться в context.

Разобратся с обработкой ошибок в горутине.
*/

func main() {
	cfg, err := config.Init()
	if err != nil {
		log.Fatalf("error initialization configs: %s\n", err.Error())
	}

	db, err := postgres.New(cfg.Postgres.DSN)
	if err != nil {
		log.Fatalf("fail to initialization db: %s\n", err.Error())
	}

	redisClient, err := redis.New(context.Background(), cfg.Redis.Port, cfg.Redis.Password)
	if err != nil {
		log.Fatalf("fail to initialization redis %s\n", err.Error())
	}

	repository := repository.New(db)

	service := service.New(repository, redisClient, &cfg.Auth)

	controller := http.New(service)

	srv := new(server.Server)

	go func() {
		if err := srv.Run(cfg.Api.WSPort, controller.InitWSRouts()); err != nil {
			log.Fatalf("error occurred while running WS server: %s\n", err.Error())
		}
	}()

	if err := srv.Run(cfg.Api.RestPort, controller.InitRouts()); err != nil {
		log.Fatalf("error occurred while running REST server: %s\n", err.Error())
	}

}
