package models

import "gorm.io/gorm"

type User struct {
	gorm.Model
	// Id       int64  `gorm:"primaryKey" json:"id"`
	Username string  `gorm:"type:varchar(255)" json:"username"`
	Email    string  `gorm:"type:varchar(255)" json:"email"`
	Password string  `gorm:"type:varchar(255)" json:"password"`
	Photos   []Photo `json:"photos"`
	// CreatedAt time.Time `gorm:"type:DateTime" json:"createdAt"`
	// UpdateAt  time.Time `gorm:"type:DateTime" json:"updatedAt"`
}
