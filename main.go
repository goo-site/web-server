package main

import (
	"github.com/goo-site/log"
	"web-server/internal/config"
	"web-server/internal/dal/mysql"
	"web-server/internal/utils"
	"web-server/router"
)

func main() {
	config.Init()
	utils.LogInit()
	mysql.InitMySQL()
	log.Info("WebServer start!")
	defer func() {
		log.Info("WebServer stop!\n")
	}()

	router.Run()
}
