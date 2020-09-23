package config

import (
	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/mysql"
	"github.com/syahjamal/go-websocket/models"
)

//DB variable
var DB *gorm.DB

// InitDB function
func InitDB() {
	var err error

	DB, err = gorm.Open("mysql", "root:root@/timesheet?charset=utf8&parseTime=True&loc=Local")
	if err != nil {
		panic("failed to connect database")
	}

	DB.AutoMigrate(&models.Notification{})
}
