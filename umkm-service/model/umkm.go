package model

// UMKM main entity (split associated models into separate files)
type UMKM struct {
	ID          uint   `gorm:"primaryKey" json:"id"`
	Name        string `gorm:"size:150;not null" json:"name"`
	About       string `gorm:"type:text" json:"about"`
	Description string `gorm:"type:text" json:"description"`
	Icon        string `gorm:"size:255" json:"icon"`
	Slug        string `gorm:"size:150;uniqueIndex" json:"slug"`
	Type        string `gorm:"size:100" json:"type"`

	OwnerID   uint           `json:"-"`
	Owner     UMKMOwner      `gorm:"foreignKey:OwnerID" json:"-"`
	Gallery   []UMKMGallery  `gorm:"foreignKey:UMKMID" json:"-"`
	Location  UMKMLocation   `gorm:"foreignKey:UMKMID" json:"-"`
	WorkHours []UMKMWorkHour `gorm:"foreignKey:UMKMID" json:"-"`
	Ratings   []UMKMRating   `gorm:"foreignKey:UMKMID" json:"-"`
}

func (UMKM) TableName() string { return "umkms" }
