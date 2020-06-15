package cmd

import (
	"5conf/router"
	"github.com/gin-gonic/gin"
)

func HttpServer() {
	server := gin.New()

	server.Use(gin.Recovery())

	router.InitRouter(server)

	server.Run()
}