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
	repo        Repository
	merchantSvc merchant.MerchantService
}

func NewProductService(repo Repository, merchantSvc merchant.MerchantService) ProductService {
	return &productService{
		repo:        repo,
		merchantSvc: merchantSvc,
	}
}

func (s *productService) CreateProduct(ctx context.Context, req *CreateProductRequest) (*Product, error) {
	// return s.repo.Create(ctx, req)
	return nil, nil
}
func (s *productService) GetProductList(ctx context.Context, name string) (*ProductList, error) {
	products, err := s.repo.FindByName(ctx, name)
	if err != nil {
		return nil, err
	}

	merchantIDs := []string{}
	productsMapByMerchantID := map[string][]*Product{}
	for _, p := range products {
		merchantIDs = append(merchantIDs, p.MerchantInfo.ID)
		productsMapByMerchantID[p.MerchantInfo.ID] = append(productsMapByMerchantID[p.MerchantInfo.ID], p)
	}

	if len(merchantIDs) > 0 {
		merchants, err := s.merchantSvc.GetMerchantByIDs(ctx, merchantIDs)
		if err != nil {
			return nil, err
		}

		for _, m := range merchants {
			if _, ok := productsMapByMerchantID[m.ID]; !ok {
				continue
			}

			for _, p := range productsMapByMerchantID[m.ID] {
				p.MerchantInfo = m
			}
		}
	}

	return &ProductList{
		Products: products,
	}, nil
}
func (s *productService) GetProductDetail(ctx context.Context, id string) (*ProductDetail, error) {
	product, err := s.repo.FindByID(ctx, id)
	if err != nil {
		return nil, err
	}

	merchantInfo, err := s.merchantSvc.GetMerchantByID(ctx, product.MerchantInfo.ID)
	if err != nil {
		return nil, err
	}
	product.MerchantInfo = merchantInfo

	return &ProductDetail{
		Product: product,
	}, nil
}
