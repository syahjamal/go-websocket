package main

import (
	"fmt"
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/syahjamal/go-websocket/routes"

	"github.com/syahjamal/go-websocket/config"
	"github.com/syahjamal/go-websocket/websocket"
)

func serveWs(pool *websocket.Pool, w http.ResponseWriter, r *http.Request) {
	fmt.Println("WebSocket Endpoint Hit")
	conn, err := websocket.Upgrade(w, r)
	if err != nil {
		fmt.Fprintf(w, "%+v\n", err)
	}

	client := &websocket.Client{
		Conn: conn,
		Pool: pool,
	}

	client.Read()

}

func setupRoutes() {
	config.InitDB()
	defer config.DB.Close()

	pool := websocket.NewPool()
	go pool.Start()

	http.HandleFunc("/ws", func(w http.ResponseWriter, r *http.Request) {
		serveWs(pool, w, r)
	})

}

func main() {

	config.InitDB()
	defer config.DB.Close()

	notif := gin.Default()

	notif.GET("/", routes.GetNotif)
	// notif.GET("/:id", routes.GetNotifID)
	notif.GET("/message", routes.GetMessage)
	notif.POST("/post", routes.PostNotif)
	notif.Run()
}
