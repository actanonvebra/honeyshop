// payment verification and transaction management.
package services

import (
	"github.com/actanonvebra/honeyshop/internal/models"
	"github.com/actanonvebra/honeyshop/internal/repositories"
)

type CheckoutService interface {
	ProcessCheckout(checkout models.Checkout) error
	CreateCheckout(checkout models.Checkout) error
}

type DefaultCheckoutService struct {
	Repo repositories.CheckoutRepository
}

func NewCheckoutService(repo repositories.CheckoutRepository) CheckoutService {
	return &DefaultCheckoutService{Repo: repo}
}
func (s *DefaultCheckoutService) ProcessCheckout(checkout models.Checkout) error {
	return s.Repo.SaveCheckout(checkout)
}

func (s *DefaultCheckoutService) CreateCheckout(checkout models.Checkout) error {
	return s.Repo.SaveCheckout(checkout)
}
