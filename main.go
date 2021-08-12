package main

import (
	productHandler "github.com/faruqfadhil/hybrida/api/http/product"
	"github.com/faruqfadhil/hybrida/internal/merchant"
	"github.com/faruqfadhil/hybrida/internal/product"
	productRepo "github.com/faruqfadhil/hybrida/internal/repository/inmem/product"
	"github.com/gin-gonic/gin"
)

func main() {
	initProducts := []*product.Product{
		{
			ID:           "PROD-1",
			Name:         "kemeja",
			ThumbnailURL: "https://abcde.com/image",
			Description:  "ini kemeja",
			Price:        10000,
			MerchantInfo: &merchant.Merchant{
				ID:          "MERCH-1",
				Name:        "Matahari",
				Description: "Matahari mall menjual segalanya",
				Address:     "Jl. cipedak 1, Jaksel",
				PhoneNumber: "085859943222",
			},
		},
	}

	productRepository := productRepo.NewProductRepository(initProducts)
	productSvc := product.NewProductService(productRepository)
	productHandler := productHandler.NewProductHandler(productSvc)

	// set root router gor gin
	router := gin.New()
	initRouterV1(router, productHandler)

	router.Run(":8080")
}

func initRouterV1(r *gin.Engine, productHandler productHandler.ProductHandler) {
	v1Router := r.Group("api/v1")
	v1Router.GET("/products", productHandler.GetListProduct)
	v1Router.GET("/product", productHandler.GetProductDetail)
}
