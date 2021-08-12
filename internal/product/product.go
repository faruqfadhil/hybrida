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
	ID           string             `json:"id"`
	Name         string             `json:"name"`
	ThumbnailURL string             `json:"thumbnail_url"`
	Description  string             `json:"description"`
	Price        int                `json:"price"`
	MerchantInfo *merchant.Merchant `json:"merchant"`
}

type ProductList struct {
	Products []*Product `json:"products"`
}

type ProductDetail struct {
	Product *Product `json:"product"`
}

type ProductService interface {
	CreateProduct(ctx context.Context, req *CreateProductRequest) (*Product, error)
	GetProductList(ctx context.Context, name string) (*ProductList, error)
	GetProductDetail(ctx context.Context, id string) (*ProductDetail, error)
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
func (s *productService) GetProductList(ctx context.Context, name string) (*ProductList, error) {
	products, err := s.repo.FindByName(ctx, name)
	if err != nil {
		return nil, err
	}
	return &ProductList{
		Products: products,
	}, nil
}
func (s *productService) GetProductDetail(ctx context.Context, id string) (*ProductDetail, error) {
	return nil, nil
}
