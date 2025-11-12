package model

type UMKM struct {
	ID      uint   `gorm:"primaryKey" json:"id"`
	Name    string `gorm:"size:150;not null" json:"name"`
	Email   string `gorm:"size:150;not null" json:"email"`
	Address string `gorm:"size:255" json:"address"`
}
