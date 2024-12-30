// product listing.
package services

import (
	"github.com/actanonvebra/honeyshop/internal/models"
	"github.com/actanonvebra/honeyshop/internal/repositories"
)

type ProductService interface {
	FetchAllProducts() ([]models.Product, error)
	SearchProducts(keyword string) ([]models.Product, error)
	AddProduct(product models.Product) error
	GetProductByID(productID string) (*models.Product, error)
}

type DefaultProductService struct {
	Repo repositories.ProductRepository
}

func NewProductService(repo repositories.ProductRepository) *DefaultProductService {
	return &DefaultProductService{Repo: repo}
}

func (s *DefaultProductService) FetchAllProducts() ([]models.Product, error) {
	return s.Repo.GetAllProducts()
}

func (s *DefaultProductService) SearchProducts(keyword string) ([]models.Product, error) {
	return s.Repo.SearchProducts(keyword)
}

func (s *DefaultProductService) AddProduct(product models.Product) error {
	return s.Repo.AddProduct(product)
}

func (s *DefaultProductService) GetProductByID(productID string) (*models.Product, error) {
	return s.Repo.FindProductByID(productID)
}
