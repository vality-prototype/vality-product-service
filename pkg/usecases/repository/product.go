package repository

import "github.com/vality-prototype/vality-product-service/pkg/entities/models"

type ProductRepository interface {
	Create(prod *models.Product) (*models.Product, error)
	Exists(name string) (bool, error)
}
