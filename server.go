package main

import (
	"log"
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/gorilla/websocket"
	"github.com/syahjamal/go-websocket/config"
	"github.com/syahjamal/go-websocket/models"
)

var clients = make(map[*websocket.Conn]bool)

var upgrader = websocket.Upgrader{
	CheckOrigin: func(r *http.Request) bool {
		return true
	},
}

func main() {
	r := gin.Default()
	// Create a simple file server
	fs := http.FileServer(http.Dir("../public"))
	http.Handle("/", fs)

	// Configure websocket route
	// http.HandleFunc("/ws", handleConnections)

	// http.HandleFunc("/ws", handleConnections)
	r.GET("/ws", handleConnections)

	go handleMessages()

	log.Println("http server started on :8000")
	err := http.ListenAndServe(":8000", nil)
	if err != nil {
		log.Fatal("ListenAndServe: ", err)
	}

}

func handleConnections(c *gin.Context) {
	// Upgrade initial GET request to a websocket
	ws, err := upgrader.Upgrade(c.Writer, c.Request, nil)
	if err != nil {
		log.Fatal(err)
	}
	// Make sure we close the connection when the function returns
	defer ws.Close()

	// Register our new client
	clients[ws] = true

	for {
		// Read in a new message as JSON and map it to a Message object
		createdAt := c.Param("created_at")

		var message models.Notification
		config.DB.Where("created_at > ?", createdAt).Find(&message)
		c.JSON(200, gin.H{
			"status": "berhasil mendapatkan data",
			"data":   message.Message,
		})
		// err := ws.ReadJSON(&message)
		// if err != nil {
		// 	log.Printf("error: %v", err)
		// 	delete(clients, ws)
		// 	break
		// }
		// Send the newly received message to the broadcast channel
		models.Broadcast <- message
	}
}

func handleMessages() {
	for {
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
}
