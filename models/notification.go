package models

import (
	"github.com/jinzhu/gorm"
)

// Notification struct
type Notification struct {
	gorm.Model
	UserIDTo string
	Username string
	Message  string `json:"message"`
}
