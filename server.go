package main

import (
	"log"
	"net/http"

	"github.com/gin-contrib/static"
	"github.com/syahjamal/go-websocket/config"
	"github.com/syahjamal/go-websocket/routes"

	"github.com/gin-gonic/gin"
	"github.com/gorilla/websocket"
	"github.com/syahjamal/go-websocket/models"
)

//Clients Variable
var Clients = make(map[*websocket.Conn]bool)

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
	// r.StaticFS("/", http.Dir("./frontend/src"))
	r.Use(static.Serve("/", static.LocalFile("./frontend/src", true)))

	// Configure websocket route
	// http.HandleFunc("/ws", handleConnections)
	r.GET("/ws", handleConnections)
	// r.GET("/getnotif", routes.GetMessage)
	// go GetMsg()

	go handleMessages()

	r.Run(":8080")

}

//GetMsg function
func GetMsg() {
	r := gin.Default()
	r.Use(routes.GetMessage)
}

func handleConnections(c *gin.Context) {
	// Upgrade initial GET request to a websocket
	// ws, err := upgrader.Upgrade(c.Writer, c.Request, nil)
	ws, err := upgrader.Upgrade(c.Writer, c.Request, nil)
	if err != nil {
		log.Fatal(err)
	}
	// Make sure we close the connection when the function returns
	defer ws.Close()

	// Register our new client
	Clients[ws] = true

	createdAt := c.Param("created_at")
	// log.Printf("createdAt: %v", createdAt)

	var notif models.Notification
	config.DB.Where("created_at > ?", createdAt).Find(&notif)
	defer config.DB.Close()
	// log.Printf("notif: %v", notif.Username)
	for {
		// ? check
		err := ws.WriteJSON(notif.Message)

		// log.Printf("message: %v", err)
		if err != nil {
			// log.Printf("error: %v", err)
			delete(Clients, ws)
			break
		}

		models.Broadcast <- notif
	}

}

func handleMessages() {
	message := <-models.Broadcast

	for client := range Clients {
		// log.Printf("client: %v", client)

		err := client.ReadJSON(message.Message)
		if err != nil {
			log.Printf("error: %v", err)
			client.Close()
			delete(Clients, client)
		}
	}

}
