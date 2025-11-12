package model

// UMKMWorkHour represents a single day's opening hours; one-to-many per UMKM
type UMKMWorkHour struct {
	ID     uint   `gorm:"primaryKey" json:"id"`
	UMKMID uint   `gorm:"index:idx_umkm_day,unique" json:"umkm_id"`
	Day    string `gorm:"type:ENUM('monday','tuesday','wednesday','thursday','friday','saturday','sunday');not null;index:idx_umkm_day,unique" json:"day"`
	Hours  string `gorm:"size:50;not null" json:"hours"`
}

func (UMKMWorkHour) TableName() string { return "umkm_work_hours" }
