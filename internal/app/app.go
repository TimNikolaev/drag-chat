package app

import (
	"context"
	"log"
	"os"
	"os/signal"
	"syscall"

	"github.com/TimNikolaev/drag-chat/internal/config"
	"github.com/TimNikolaev/drag-chat/internal/delivery/http"
	"github.com/TimNikolaev/drag-chat/internal/repository"
	"github.com/TimNikolaev/drag-chat/internal/repository/postgres"
	"github.com/TimNikolaev/drag-chat/internal/server"
	"github.com/TimNikolaev/drag-chat/internal/service"
	"github.com/TimNikolaev/drag-chat/pkg/event/redis"
)

type App struct {
	cfg *config.Config
}

func New(cfg *config.Config) *App {
	return &App{
		cfg: cfg,
	}
}

func (a *App) Run() {
	db, err := postgres.New(a.cfg.Postgres.DSN)
	if err != nil {
		log.Fatalf("fail to initialization db: %s\n", err.Error())
	}

	redisClient, err := redis.New(context.Background(), a.cfg.Redis.Port, a.cfg.Redis.Password)
	if err != nil {
		log.Fatalf("fail to initialization redis %s\n", err.Error())
	}

	repository := repository.New(db)

	service := service.New(repository, redisClient, &a.cfg.Auth)

	controller := http.New(service)

	srv := server.New(a.cfg.Api.RestPort, controller.InitRouts())

	go func() {
		if err := srv.Run(); err != nil {
			log.Fatalf("error occurred while running WS server: %s\n", err.Error())
		}
	}()

	quit := make(chan os.Signal, 1)
	signal.Notify(quit, syscall.SIGTERM, syscall.SIGINT)
	<-quit

	if err := srv.Stop(context.Background()); err != nil {
		log.Fatalf("error occurred on server shutting down: %s", err.Error())
	}

	if err := db.Close(); err != nil {
		log.Fatalf("error occurred on db connection close: %s", err.Error())
	}
}
