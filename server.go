package main

import (
	"github.com/gin-gonic/gin"
	"github.com/syahjamal/go-websocket/config"
	"github.com/syahjamal/go-websocket/routes"
)

func main() {
	config.InitDB()
	defer config.DB.Close()

	router := gin.Default()

	notif := router.Group("/test")
	{
		notif.GET("/", routes.GetNotif)
		notif.GET("/:id", routes.GetNotifId)
	}
	router.Run()
}
