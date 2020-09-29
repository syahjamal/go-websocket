package main

import (
	"log"
	"net/http"

	"github.com/syahjamal/go-websocket/config"
	"github.com/syahjamal/go-websocket/routes"

	"github.com/gin-gonic/gin"
	"github.com/gorilla/websocket"
	"github.com/syahjamal/go-websocket/models"
)

var clients = make(map[*websocket.Conn]bool)

var upgrader = websocket.Upgrader{
	CheckOrigin: func(r *http.Request) bool {
		return true
	},
}

func main() {
	config.InitDB()
	defer config.DB.Close()

	r := gin.Default()

	r.POST("/post", routes.PostNotif)
	r.StaticFS("/", http.Dir("./frontend/src"))

	// Configure websocket route
	http.HandleFunc("/ws", handleConnections)
	// r.GET("/getnotif", routes.GetMessage)

	go handleMessages()

	// log.Println("http server started on :8000")
	// err := http.ListenAndServe(":8000", nil)
	// if err != nil {
	// 	log.Fatal("ListenAndServe: ", err)
	// }
	r.Run()

}

func handleConnections(w http.ResponseWriter, r *http.Request) {
	// Upgrade initial GET request to a websocket
	// ws, err := upgrader.Upgrade(c.Writer, c.Request, nil)
	ws, err := upgrader.Upgrade(w, r, nil)
	if err != nil {
		log.Fatal(err)
	}
	// Make sure we close the connection when the function returns
	defer ws.Close()

	// Register our new client
	clients[ws] = true
	for {
		// var ctx *gin.Context
		// routes.GetMessage(ctx)

		var notif models.Notification

		err := ws.ReadJSON(&notif)
		if err != nil {
			log.Printf("error: %v", err)
			delete(clients, ws)
			break
		}

		models.Broadcast <- notif

	}

}

func handleMessages() {
	message := <-models.Broadcast
	for client := range clients {
		err := client.WriteJSON(message)
		if err != nil {
			log.Printf("error: %v", err)
			client.Close()
			delete(clients, client)
		}
	}
}
