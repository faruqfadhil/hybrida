package product

import (
	"context"

	"github.com/faruqfadhil/hybrida/internal/product"
)

type productRepository struct {
	products []*product.Product
}

func NewProductRepository(products []*product.Product) product.Repository {
	return &productRepository{
		products: products,
	}
}

func (r *productRepository) Create(ctx context.Context, product *product.Product) error {
	return nil
}
func (r *productRepository) FindByName(ctx context.Context, name string) ([]*product.Product, error) {
	return nil, nil
}
func (r *productRepository) FindByID(ctx context.Context, id string) (*product.Product, error) {
	return nil, nil
}
