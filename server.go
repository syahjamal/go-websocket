package main

import (
	"log"
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/gorilla/websocket"
	"github.com/syahjamal/go-websocket/models"
	"github.com/syahjamal/go-websocket/routes"
)

// musti gimana gw bgung ga konek mulu depan ma blkng nya//error dimana
// Hasil ngetestnya dimana
// koneksi depannya dimana
var clients = make(map[*websocket.Conn]bool)

var upgrader = websocket.Upgrader{
	CheckOrigin: func(r *http.Request) bool {
		return true
	},
}

func main() {
	r := gin.Default()
	// Create a simple file server
	// fs := http.FileServer(http.Dir("./frontend"))
	// http.Handle("/", fs)
	r.StaticFS("/", http.Dir("./frontend/src"))

	// Configure websocket route
	// http.HandleFunc("/ws", handleConnections)

	http.HandleFunc("/ws", handleConnections)
	// r.GET("/ws", handleConnections)
	// r.GET("/ws", gin.Wrap(handleConnections))

	go handleMessages()

	log.Println("http server started on :8000")
	err := http.ListenAndServe(":8000", nil)
	if err != nil {
		log.Fatal("ListenAndServe: ", err)
	}

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
		var ctx *gin.Context
		routes.GetMessage(ctx)
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
