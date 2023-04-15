package repository

import (
	"github.com/vality-prototype/vality-product-service/pkg/entities/models"
	"github.com/vality-prototype/vality-product-service/pkg/usecases/repository"
	"github.com/vality-prototype/vality-utility-service/pkg"
	"gorm.io/gorm"
)

type productRepository struct {
	appCtx pkg.AppContext
}

func NewProductRepo(appCtx pkg.AppContext) repository.ProductRepository {
	return &productRepository{
		appCtx: appCtx,
	}
}

func (r *productRepository) Exists(name string) (bool, error) {
	var (
		db   = r.appCtx.GetMainDBConnection()
		prod = models.Product{}
	)
	if err := db.Where("name = ?", name).First(&prod).Error; err != nil {
		if err != gorm.ErrRecordNotFound {
			return false, err
		}
	}
	return prod.ID != 0, nil
}

func (r *productRepository) Create(prod *models.Product) (*models.Product, error) {
	var (
		db = r.appCtx.GetMainDBConnection()
	)
	if err := db.Create(prod).Error; err != nil {
		return nil, err
	}
	return prod, nil
}
