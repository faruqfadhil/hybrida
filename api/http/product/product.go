package product

import (
	"github.com/faruqfadhil/hybrida/internal/product"
	"github.com/gin-gonic/gin"
)

type ProductHandler interface {
	CreateProduct(c *gin.Context)
	GetListProduct(c *gin.Context)
	GetProductDetail(c *gin.Context)
}

type productHandler struct {
	svc product.ProductService
}

func NewProductHandler(svc product.ProductService) ProductHandler {
	return &productHandler{
		svc: svc,
	}
}

func (h *productHandler) CreateProduct(c *gin.Context)    {}
func (h *productHandler) GetListProduct(c *gin.Context)   {}
func (h *productHandler) GetProductDetail(c *gin.Context) {}
