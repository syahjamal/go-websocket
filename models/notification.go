package models

import "github.com/jinzhu/gorm"

//Notification struct
type Notification struct {
	gorm.Model
	UserID   string
	Username string
	Message  string
}
