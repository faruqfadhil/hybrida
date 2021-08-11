package product

import (
	"context"

	"github.com/faruqfadhil/hybrida/internal/merchant"
)

type CreateProductRequest struct {
	Name         string
	ThumbnailURL string
	Description  string
	Price        int
	MerchantID   string
}

type Product struct {
	ID           string
	Name         string
	ThumbnailURL string
	Description  string
	Price        int
	MerchantInfo *merchant.Merchant
}

type ProductService interface {
	CreateProduct(ctx context.Context, req *CreateProductRequest) (*Product, error)
	GetProductList(ctx context.Context, name string) ([]*Product, error)
	GetProductDetail(ctx context.Context, id string) (*Product, error)
}

type productService struct {
	repo Repository
}

func NewProductService(repo Repository) ProductService {
	return &productService{
		repo: repo,
	}
}

func (s *productService) CreateProduct(ctx context.Context, req *CreateProductRequest) (*Product, error) {
	return nil, nil
}
func (s *productService) GetProductList(ctx context.Context, name string) ([]*Product, error) {
	return nil, nil
}
func (s *productService) GetProductDetail(ctx context.Context, id string) (*Product, error) {
	return nil, nil
}
