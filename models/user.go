package models

type User struct {
	Id        int64  `gorm:"primaryKey" json:"id"`
	Username  string `gorm:"type:varchar(255)" json:"username"`
	Email     string `gorm:"type:varchar(255)" json:"email"`
	Password  string `gorm:"type:varchar(255)" json:"password"`
	CreatedAt string `gorm:"type:DateTime" json:"createdAt"`
	UpdateAt  string `gorm:"type:DateTime" json:"updateAt"`
}
