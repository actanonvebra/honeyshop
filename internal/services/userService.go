// user login and resgistration.
package services

import (
	"errors"

	"github.com/actanonvebra/honeyshop/internal/models"
	"github.com/actanonvebra/honeyshop/internal/repositories"
)

type UserService interface {
	Login(username, password string) (models.User, error)
}

type DefaultUserService struct {
	Repo repositories.UserRepository
}

func (s *DefaultUserService) Login(username, password string) (models.User, error) {
	user, err := s.Repo.GetUserByUserName(username)
	if err != nil {
		return models.User{}, errors.New("User not found")
	}

	if user.Password != password {
		return models.User{}, errors.New("Invalid password")
	}
	return user, nil
}
