package product

import (
	"context"
	"strings"

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
	resp := []*product.Product{}
	for _, p := range r.products {
		if strings.Contains(strings.ToLower(p.Name), strings.ToLower(name)) {
			resp = append(resp, p)
		}
	}

	return resp, nil
}
func (r *productRepository) FindByID(ctx context.Context, id string) (*product.Product, error) {
	return nil, nil
}
