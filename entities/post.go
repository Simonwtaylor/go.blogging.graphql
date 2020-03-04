package entities

import (
	"time"

	"github.com/jinzhu/gorm"
)

type Post struct {
	gorm.Model
	Content    string
	DatePosted time.Time
	User       User `gorm:"foreignkey:UserID"`
	UserID     uint
}
