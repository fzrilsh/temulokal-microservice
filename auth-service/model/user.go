package model

type User struct {
	ID       uint   `gorm:"primaryKey" json:"id"`
	FullName string `gorm:"size:100;not null" json:"full_name"`
	Email    string `gorm:"uniqueIndex;size:100;not null" json:"email"`
	Password string `gorm:"size:255;not null" json:"password"`
}
