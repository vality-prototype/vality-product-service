package controller

import (
	"context"

	"github.com/vality-prototype/vality-product-service/pkg/entities/models"
	"github.com/vality-prototype/vality-product-service/pkg/usecases"
	"github.com/vality-prototype/vality-utility-service/configs"
	"github.com/vality-prototype/vality-utility-service/pkg"
)

type ProductController interface {
	CreateProduct(ctx context.Context, prod *models.Product) (*configs.Response, error)
}

type productController struct {
	appCtx         pkg.AppContext
	productUsecase usecases.Product
}

func NewProductController(appCtx pkg.AppContext, productUsecase usecases.Product) ProductController {
	return &productController{
		appCtx:         appCtx,
		productUsecase: productUsecase,
	}
}

func (p *productController) CreateProduct(ctx context.Context, prod *models.Product) (resp *configs.Response, err error) {
	var createdProd *models.Product
	createdProd, err = p.productUsecase.CreateProduct(prod)
	if err != nil {
		return
	}
	resp.SimpleResponse(createdProd)
	return
}
