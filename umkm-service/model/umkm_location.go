package model

// UMKMLocation stores the location data for an UMKM
type UMKMLocation struct {
	ID        uint    `gorm:"primaryKey" json:"id"`
	UMKMID    uint    `gorm:"index" json:"umkm_id"`
	URL       string  `gorm:"size:255" json:"url"`
	Text      string  `gorm:"size:255" json:"text"`
	ShortText string  `gorm:"size:150" json:"short_text"`
	Longitude float64 `json:"longitude"`
	Latitude  float64 `json:"latitude"`
}

func (UMKMLocation) TableName() string { return "umkm_locations" }
