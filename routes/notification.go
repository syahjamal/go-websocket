package routes

import (
	"github.com/gin-gonic/gin"
	"github.com/syahjamal/go-websocket/config"
	"github.com/syahjamal/go-websocket/models"
)

//func GetNotif
func GetNotif(c *gin.Context) {
	message := []models.Notification{}
	config.DB.Find(&message)

	c.JSON(200, gin.H{
		"status": "berhasil mendapatkan pesan",
		"data":   message,
	})
}

//GetNotifId
func GetNotifId(c *gin.Context) {
	id := c.Param("id")

	var item models.Notification

	config.DB.Model(&item).Where("id = ?", id).Find(&item)
	c.JSON(200, gin.H{
		"status": "berhasil mendapatkan data",
		"data":   item,
	})
}
