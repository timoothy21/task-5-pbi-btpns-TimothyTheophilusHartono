package models

import "time"

type User struct {
	Id        int64     `gorm:"primaryKey" json:"id"`
	Username  string    `gorm:"type:varchar(255)" json:"username"`
	Email     string    `gorm:"type:varchar(255)" json:"email"`
	Password  string    `gorm:"type:varchar(255)" json:"password"`
	CreatedAt time.Time `gorm:"type:DateTime" json:"createdAt"`
	UpdateAt  time.Time `gorm:"type:DateTime" json:"updateAt"`
}
