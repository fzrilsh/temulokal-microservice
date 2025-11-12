package model

// UMKMOwner represents owner data including socials
type UMKMOwner struct {
	ID        uint    `gorm:"primaryKey" json:"id"`
	Name      string  `gorm:"size:150;not null" json:"name"`
	Image     string  `gorm:"size:255" json:"image"`
	Phone     string  `gorm:"size:50" json:"phone"`
	Email     string  `gorm:"size:150" json:"email"`
	Website   *string `gorm:"size:255" json:"website"`
	Facebook  *string `gorm:"size:255" json:"facebook"`
	Twitter   *string `gorm:"size:255" json:"twitter"`
	Instagram *string `gorm:"size:255" json:"instagram"`
	Whatsapp  *string `gorm:"size:255" json:"whatsapp"`
}

func (UMKMOwner) TableName() string { return "umkm_owners" }
