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

type ProductEntity struct {
	ID           string
	Name         string
	ThumbnailURL string
	Description  string
	Price        int
	MerchantInfo *merchant.MerchantEntity
}

type Product interface {
	CreateProduct(ctx context.Context, req *CreateProductRequest) (*ProductEntity, error)
	GetProductList(ctx context.Context, name string) ([]*ProductEntity, error)
	GetProductDetail(ctx context.Context, id string) (*ProductEntity, error)
}
