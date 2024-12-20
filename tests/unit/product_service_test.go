package unit

import (
	"testing"

	"github.com/actanonvebra/honeyshop/internal/models"
	"github.com/actanonvebra/honeyshop/internal/services"
	"github.com/stretchr/testify/assert"
)

type MockRepository struct{}

func (repo *MockRepository) GetAllProducts() ([]models.Product, error) {
	return []models.Product{
		{
			ID:          "1",
			Name:        "Wireless Mouse",
			Description: "A high-precision wireless mouse with ergonomic design.",
			Price:       29.99,
			Stock:       150,
			Category:    "Electronics",
		},
		{
			ID:          "2",
			Name:        "Mechanical Keyboard",
			Description: "Durable mechanical keyboard with customizable RGB lighting.",
			Price:       69.99,
			Stock:       85,
			Category:    "Electronics",
		},
	}, nil
}

func TestFetchAllProducts(t *testing.T) {
	mockRepo := &MockRepository{}
	service := services.DefaultProductService{Repo: mockRepo}
	products, err := service.FetchAllProducts()

	assert.NoError(t, err, "There should be no error")
	assert.Equal(t, 2, len(products), "There should be 2 products")
	assert.Equal(t, "Wireless Mouse", products[0].Name, "The first product should be'Wireless Mouse")
	assert.Equal(t, 29.99, products[0].Price, "The price should be 29.99")
}
