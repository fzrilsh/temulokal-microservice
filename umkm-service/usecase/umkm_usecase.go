package usecase

import (
	"temulokal-microservice/umkm-service/model"
	"temulokal-microservice/umkm-service/repository"
)

type UMKMResponse struct {
	ID           uint              `json:"id"`
	Name         string            `json:"name"`
	About        string            `json:"about"`
	Description  string            `json:"description"`
	Icon         string            `json:"icon"`
	Slug         string            `json:"slug"`
	Type         string            `json:"type"`
	Rating       UMKMRating        `json:"rating"`
	Owner        UMKMOwnerResponse `json:"owner"`
	Gallery      []UMKMGalleryItem `json:"gallery"`
	Location     UMKMLocationResp  `json:"location"`
	OpeningHours OpeningHoursResp  `json:"opening_hours"`
}

type UMKMRating struct {
	Count int     `json:"count"`
	Star  float32 `json:"star"`
}

type UMKMOwnerResponse struct {
	Name    string          `json:"name"`
	Image   string          `json:"image"`
	Phone   string          `json:"phone"`
	Email   string          `json:"email"`
	Website *string         `json:"website"`
	Socials OwnerSocialResp `json:"socials"`
}

type OwnerSocialResp struct {
	Facebook  *string `json:"facebook"`
	Twitter   *string `json:"twitter"`
	Instagram *string `json:"instagram"`
	Whatsapp  *string `json:"whatsapp"`
}

type UMKMGalleryItem struct {
	URL string `json:"url"`
}

type UMKMLocationResp struct {
	URL       string  `json:"url"`
	Text      string  `json:"text"`
	ShortText string  `json:"short_text"`
	Longitude float64 `json:"logitude"`
	Latitude  float64 `json:"latitude"`
}

type OpeningHoursResp struct {
	Monday    string `json:"monday"`
	Tuesday   string `json:"tuesday"`
	Wednesday string `json:"wednesday"`
	Thursday  string `json:"thursday"`
	Friday    string `json:"friday"`
	Saturday  string `json:"saturday"`
	Sunday    string `json:"sunday"`
}

type UMKMUsecase struct {
	repo repository.UMKMRepository
}

func NewUMKMUsecase(repo repository.UMKMRepository) *UMKMUsecase {
	return &UMKMUsecase{repo: repo}
}

func (u *UMKMUsecase) List() ([]UMKMResponse, error) {
	items, err := u.repo.FindAll()
	if err != nil {
		return nil, err
	}
	resp := make([]UMKMResponse, 0, len(items))
	for _, i := range items {
		resp = append(resp, mapUMKMToResponse(i))
	}
	return resp, nil
}

func mapUMKMToResponse(m model.UMKM) UMKMResponse {
	// Map gallery
	gallery := make([]UMKMGalleryItem, 0, len(m.Gallery))
	for _, g := range m.Gallery {
		gallery = append(gallery, UMKMGalleryItem{URL: g.URL})
	}

	// Map owner
	owner := UMKMOwnerResponse{
		Name:    m.Owner.Name,
		Image:   m.Owner.Image,
		Phone:   m.Owner.Phone,
		Email:   m.Owner.Email,
		Website: m.Owner.Website,
		Socials: OwnerSocialResp{
			Facebook:  m.Owner.Facebook,
			Twitter:   m.Owner.Twitter,
			Instagram: m.Owner.Instagram,
			Whatsapp:  m.Owner.Whatsapp,
		},
	}

	// Map location
	location := UMKMLocationResp{
		URL:       m.Location.URL,
		Text:      m.Location.Text,
		ShortText: m.Location.ShortText,
		Longitude: m.Location.Longitude,
		Latitude:  m.Location.Latitude,
	}

	// Map opening hours
	// Map work hours slice into structured days
	opening := OpeningHoursResp{}
	for _, wh := range m.WorkHours {
		switch wh.Day {
		case "monday":
			opening.Monday = wh.Hours
		case "tuesday":
			opening.Tuesday = wh.Hours
		case "wednesday":
			opening.Wednesday = wh.Hours
		case "thursday":
			opening.Thursday = wh.Hours
		case "friday":
			opening.Friday = wh.Hours
		case "saturday":
			opening.Saturday = wh.Hours
		case "sunday":
			opening.Sunday = wh.Hours
		}
	}

	return UMKMResponse{
		ID:          m.ID,
		Name:        m.Name,
		About:       m.About,
		Description: m.Description,
		Icon:        m.Icon,
		Slug:        m.Slug,
		Type:        m.Type,
		Rating: UMKMRating{
			Count: m.RatingCount,
			Star:  m.RatingStar,
		},
		Owner:        owner,
		Gallery:      gallery,
		Location:     location,
		OpeningHours: opening,
	}
}
