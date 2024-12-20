// product listing.
package services

import (
	"github.com/actanonvebra/honeyshop/internal/models"
	"github.com/actanonvebra/honeyshop/internal/repositories"
)

type ProductService interface {
	FetchAllProducts() ([]models.Product, error)
}

type DefaultProductService struct {
	Repo repositories.ProductRepository
}

func (s *DefaultProductService) FetchAllProducts() ([]models.Product, error) {
	return s.Repo.GetAllProducts()
}
