package entity

import (
	"gorm.io/gorm"
	"time"
)

type User struct {
	ID 					uint64		`gorm:"primary_key:auto_increment" json:"id"`
	Phone				string		`gorm:"uniqueIndex;type:varchar(255)" json:"phone"`
	Password			string		`gorm:"->;<-;not null" json:"password"`
	Email				string		`gorm:"uniqueIndex;type:varchar(255)" json:"email"`
	IsVerifiedEmail     bool        `grom:"type:varchar(255)" json:"is_verified_email"`
	FirstName			string		`grom:"type:varchar(255)" json:"first_name"`
	LastName			string		`grom:"type:varchar(255)" json:"last_name"`
	DataOFBirth			string		`grom:"type:varchar(255)" json:"data_of_birth"`
	BVN					string		`grom:"uniqueIndex;type:varchar(255); json:"bvn"`
	Citizenship			string		`grom:"type:varchar(255)" json:"citizenship"`
	IsVerifiedKYC		bool		`grom:"type:varchar(255)" json:"is_verified_kyc"`

	Token    			string  	`gorm:"-" json:"token,omitempty"`
	CreatedAt 			time.Time
	UpdatedAt 			time.Time
	DeletedAt 			gorm.DeletedAt	`gorm:"index"`
}

