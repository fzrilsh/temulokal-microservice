package model

// UMKM main entity
type UMKM struct {
	ID          uint    `gorm:"primaryKey" json:"id"`
	Name        string  `gorm:"size:150;not null" json:"name"`
	About       string  `gorm:"type:text" json:"about"`
	Description string  `gorm:"type:text" json:"description"`
	Icon        string  `gorm:"size:255" json:"icon"`
	Slug        string  `gorm:"size:150;uniqueIndex" json:"slug"`
	Type        string  `gorm:"size:100" json:"type"`
	RatingCount int     `gorm:"default:0" json:"-"`
	RatingStar  float32 `gorm:"default:0" json:"-"`

	OwnerID   uint           `json:"-"`
	Owner     UMKMOwner      `gorm:"foreignKey:OwnerID" json:"-"`
	Gallery   []UMKMGallery  `gorm:"foreignKey:UMKMID" json:"-"`
	Location  UMKMLocation   `gorm:"foreignKey:UMKMID" json:"-"`
	WorkHours []UMKMWorkHour `gorm:"foreignKey:UMKMID" json:"-"`
}

func (UMKM) TableName() string { return "umkms" }

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

// UMKMGallery stores images for an UMKM
type UMKMGallery struct {
	ID     uint   `gorm:"primaryKey" json:"id"`
	UMKMID uint   `gorm:"index" json:"umkm_id"`
	URL    string `gorm:"size:255" json:"url"`
}

func (UMKMGallery) TableName() string { return "umkm_gallerys" }

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

// UMKMWorkHour represents a single day's opening hours; one-to-many per UMKM
type UMKMWorkHour struct {
	ID     uint   `gorm:"primaryKey" json:"id"`
	UMKMID uint   `gorm:"index:idx_umkm_day,unique" json:"umkm_id"`
	Day    string `gorm:"type:ENUM('monday','tuesday','wednesday','thursday','friday','saturday','sunday');not null;index:idx_umkm_day,unique" json:"day"`
	Hours  string `gorm:"size:50;not null" json:"hours"`
}

func (UMKMWorkHour) TableName() string { return "umkm_work_hours" }
