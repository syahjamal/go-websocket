package models

import (
	"github.com/jinzhu/gorm"
)

//Broadcast Variable
var Broadcast = make(chan Notification)

// Notification struct
type Notification struct {
	gorm.Model
	UserIDTo string `json:"user_id_to"`
	Username string `json:"username"`
	Message  string `json:"message"`
}
