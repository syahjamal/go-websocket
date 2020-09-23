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

	var bc models.Notification
	config.DB.Where("created_at > ?", createdAt).Find(&bc)
	c.JSON(200, gin.H{
		"status": "berhasil mendapatkan data",
		"data":   bc.Message,
	})

	return
}

//PostNotif function
func PostNotif(c *gin.Context) {
	item := models.Notification{
		UserIDTo: c.PostForm("userid"),
		Username: c.PostForm("nama"),
		Message:  c.PostForm("msg"),
	}
	config.DB.Create(&item)
	c.JSON(200, gin.H{
		"status": "berhasil post",
		"data":   item,
	})
}
