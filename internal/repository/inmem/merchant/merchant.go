package merchant

import (
	"context"
	"fmt"

	"github.com/faruqfadhil/hybrida/internal/merchant"
)

type merchantRepository struct {
	merchants []*merchant.Merchant
}

func NewMerchantRepository(merchants []*merchant.Merchant) merchant.Repository {
	return &merchantRepository{
		merchants: merchants,
	}
}

func (r *merchantRepository) FindByID(ctx context.Context, id string) (*merchant.Merchant, error) {
	for _, m := range r.merchants {
		if m.ID == id {
			return m, nil
		}
	}

	return nil, fmt.Errorf("merchant id %s not found", id)
}

func (r *merchantRepository) FindByIDs(ctx context.Context, ids []string) ([]*merchant.Merchant, error) {
	resp := []*merchant.Merchant{}
	for _, m := range r.merchants {
		for _, id := range ids {
			if id == m.ID {
				resp = append(resp, m)
			}
		}
	}

	return resp, nil
}
