package product

import "context"

type Repository interface {
	Create(ctx context.Context, product *Product) error
	FindByName(ctx context.Context, name string) ([]*Product, error)
	FindByID(ctx context.Context, id string) (*Product, error)
}
