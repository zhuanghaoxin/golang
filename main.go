// @Author zhangxinmin 2023/12/19 22:40:00
package main

import (
	log "github.com/sirupsen/logrus"
	"user-center/src/router"
)

func main() {
	log.Info("The project is start!")
	Start()
}

func Start() {
	router.InitRouterAndServer()
}
