package model

// UMKMGallery stores images for an UMKM
type UMKMGallery struct {
	ID     uint   `gorm:"primaryKey" json:"id"`
	UMKMID uint   `gorm:"index" json:"umkm_id"`
	URL    string `gorm:"size:255" json:"url"`
}

func (UMKMGallery) TableName() string { return "umkm_gallerys" }
