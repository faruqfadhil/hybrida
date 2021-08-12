package product

import (
	"context"
	"net/http"

	"github.com/faruqfadhil/hybrida/internal/product"
	"github.com/faruqfadhil/hybrida/pkg/response"
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

func (h *productHandler) CreateProduct(c *gin.Context) {}
func (h *productHandler) GetListProduct(c *gin.Context) {
	var name string
	name, ok := c.GetQuery("name")
	if !ok {
		response.Error(c, http.StatusBadRequest, "name required")
		return
	}
	products, err := h.svc.GetProductList(context.Background(), name)
	if err != nil {
		response.Error(c, http.StatusInternalServerError, err.Error())
		return
	}

	response.Success(c, products, "success")
}
func (h *productHandler) GetProductDetail(c *gin.Context) {}
