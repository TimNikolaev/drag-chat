package main

import (
	"log"

	"github.com/TimNikolaev/drag-chat/internal/app"
	"github.com/TimNikolaev/drag-chat/internal/config"
	_ "github.com/lib/pq"
)

func main() {
	cfg, err := config.Init()
	if err != nil {
		log.Fatalf("error initialization configs: %s\n", err.Error())
	}

	app := app.New(cfg)

	app.Run()
}
