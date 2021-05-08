package entity

import (
	"gorm.io/gorm"
	"time"
)

type PasswordReset struct {
	ID 					uint64		`gorm:"primary_key:auto_increment" json:"id"`

	Email				string		`gorm:"uniqueIndex;type:varchar(255)" json:"email"`
	Token    			string  	`gorm:"-" json:"token,omitempty"`

	CreatedAt 			time.Time
	UpdatedAt 			time.Time
	DeletedAt 			gorm.DeletedAt	`gorm:"index"`
}