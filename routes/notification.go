package routes

import (
	"github.com/gin-gonic/gin"
	"github.com/syahjamal/go-websocket/config"
	"github.com/syahjamal/go-websocket/models"
)

// GetNotif function
func GetNotif(c *gin.Context) {
	message := []models.Notification{}
	config.DB.Find(&message)

	c.JSON(200, gin.H{
		"status": "berhasil mendapatkan pesan",
		"data":   message,
	})
}

//GetNotifID function
func GetNotifID(c *gin.Context) {
	id := c.Param("id")

	var item models.Notification

	config.DB.Model(&item).Where("id = ?", id).Find(&item)
	c.JSON(200, gin.H{
		"status": "berhasil mendapatkan data",
		"data":   item,
	})
}

// GetMessage function
func GetMessage(c *gin.Context) {
	createdAt := c.Param("created_at")

	var item models.Notification
	config.DB.Where("created_at > ?", createdAt).Find(&item)
	c.JSON(200, gin.H{
		"status": "berhasil mendapatkan data",
		"data":   item.Message,
	})
}

//PostNotif function
func PostNotif(c *gin.Context) {
	item := models.Notification{
		UserIDTo: c.PostForm("user_id_to"),
		Username: c.PostForm("username"),
		Message:  c.PostForm("message"),
	}

	config.DB.Create(&item)
	c.JSON(200, gin.H{
		"status": "berhasil post",
		"data":   item,
	})
}
