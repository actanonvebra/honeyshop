// cart.
package services

import (
	"github.com/actanonvebra/honeyshop/internal/models"
	"github.com/actanonvebra/honeyshop/internal/repositories"
)

type CartService interface {
	GetCartByUserID(userID string) (*models.Cart, error)
	ClearCart(userID string) error
}

type DefaultCartService struct {
	Repo repositories.CartRepository
}

func NewCartService(repo repositories.CartRepository) CartService {
	return &DefaultCartService{Repo: repo}
}

func (s *DefaultCartService) GetCartByUserID(userID string) (*models.Cart, error) {
	return s.Repo.FindCartByUserID(userID)
}

func (s *DefaultCartService) ClearCart(userID string) error {
	return s.Repo.DeleteCartByUserID(userID)
}
