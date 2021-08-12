package merchant

import "context"

type Repository interface {
	FindByID(ctx context.Context, id string) (*Merchant, error)
	FindByIDs(ctx context.Context, ids []string) ([]*Merchant, error)
}
