package model

// UMKMRating represents a single user rating for an UMKM (value 1-5)
// Aggregation (count, average) should be computed at query/usecase layer.
type UMKMRating struct {
	ID     uint `gorm:"primaryKey" json:"id"`
	UMKMID uint `gorm:"index" json:"umkm_id"`
	// UserID uint `gorm:"index" json:"user_id"`
	Value uint8 `gorm:"type:tinyint unsigned;not null;check:value >= 1 and value <= 5" json:"value"`
}

func (UMKMRating) TableName() string { return "umkm_ratings" }
