package usecase

import "temulokal-microservice/umkm-service/repository"

type UMKMResponse struct {
	ID      uint   `json:"id"`
	Name    string `json:"name"`
	Email   string `json:"email"`
	Address string `json:"address"`
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
		resp = append(resp, UMKMResponse{
			ID:      i.ID,
			Name:    i.Name,
			Email:   i.Email,
			Address: i.Address,
		})
	}
	return resp, nil
}
