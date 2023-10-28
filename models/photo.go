package models

import "gorm.io/gorm"

type Photo struct {
	gorm.Model
	// Id       int64  `gorm:"primaryKey" json:"id"`
	Title    string `gorm:"type:varchar(255)" json:"title"`
	Caption  string `gorm:"type:varchar(255)" json:"caption"`
	PhotoURL string `gorm:"type:varchar(255)" json:"photoURL"`
	UserID   int    `json:"userID`
	User     User   `json:"user`
}
