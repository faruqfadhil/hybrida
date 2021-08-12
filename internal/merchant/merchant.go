package merchant

import (
	"context"
)

type Merchant struct {
	ID          string `json:"id"`
	Name        string `json:"name"`
	Description string `json:"description"`
	Address     string `json:"address"`
	PhoneNumber string `json:"phone_number"`
}

type MerchantService interface {
	GetMerchantByID(ctx context.Context, id string) (*Merchant, error)
	GetMerchantByIDs(ctx context.Context, ids []string) ([]*Merchant, error)
}

type merchantService struct {
	repo Repository
}

func NewMerchantService(repo Repository) MerchantService {
	return &merchantService{
		repo: repo,
	}
}

func (s *merchantService) GetMerchantByID(ctx context.Context, id string) (*Merchant, error) {
	return s.repo.FindByID(ctx, id)
}

func (s *merchantService) GetMerchantByIDs(ctx context.Context, ids []string) ([]*Merchant, error) {
	return s.repo.FindByIDs(ctx, ids)
}
