package main

import (
	"web-server/internal/config"
	"web-server/internal/utils"
	"web-server/router"

	"github.com/goo-site/log"
)

func main() {
	config.Init()
	utils.LogInit()
	log.Info("WebServer start!")
	defer func() {
		log.Info("WebServer stop!\n")
	}()

	router.Run()
}
