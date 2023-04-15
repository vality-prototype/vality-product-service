package usecases

import (
	"errors"

	"github.com/vality-prototype/vality-product-service/pkg/entities/models"
	"github.com/vality-prototype/vality-product-service/pkg/usecases/repository"
	"github.com/vality-prototype/vality-utility-service/pkg"
)

type Product interface {
	CreateProduct(prod *models.Product) (*models.Product, error)
}

type productUsecase struct {
	productRepo repository.ProductRepository
	appCtx      pkg.AppContext
}

func NewCreateUser(appCtx pkg.AppContext, productRepo repository.ProductRepository) *productUsecase {
	return &productUsecase{
		productRepo: productRepo,
		appCtx:      appCtx,
	}
}

func (c *productUsecase) CreateProduct(prod *models.Product) (*models.Product, error) {
	prodExists, err := c.productRepo.Exists(prod.Name)
	if err != nil {
		return nil, err
	}
	if prodExists {
		return nil, errors.New("product already exists")
	}
	return c.productRepo.Create(prod)
}
