package main

import (
	"github.com/Liki4/Kyoto/internal/conf"
	"github.com/Liki4/Kyoto/internal/db"
	tgbot2 "github.com/Liki4/Kyoto/internal/tgbot"
	"github.com/Liki4/Kyoto/internal/web"
	log "unknwon.dev/clog/v2"
)

func main() {
	defer log.Stop()
	err := log.NewConsole()
	if err != nil {
		panic(err)
	}

	if err = conf.Load(); err != nil {
		log.Fatal("Failed to load config: %v", err)
	}

	if err = db.ConnDB(); err != nil {
		log.Fatal("Failed to connect to MySQL database: %v", err)
	}

	tgbot2.Run()
	web.Run()
}
