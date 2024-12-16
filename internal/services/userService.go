// user login and resgistration.
package services

import (
	"errors"

	"github.com/actanonvebra/honeyshop/internal/models"
	"github.com/actanonvebra/honeyshop/internal/repositories"
)

type UserService interface {
	Login(username, password string) (models.User, error)
	Register(user models.User) error
}

type DefaultUserService struct {
	Repo repositories.UserRepository
}

func (s *DefaultUserService) Login(username, password string) (models.User, error) {
	user, err := s.Repo.GetUserByUserName(username)
	if err != nil || user.Password != password {
		return models.User{}, errors.New("invalid username or password")
	}
	return user, nil
}

func (s *DefaultUserService) Register(user models.User) error {
	return s.Repo.CreateUser(user)
}
