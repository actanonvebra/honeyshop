// user login and resgistration.
package services

import (
	"errors"
	"log"
	"time"

	"github.com/actanonvebra/honeyshop/internal/models"
	"github.com/actanonvebra/honeyshop/internal/repositories"
	"go.mongodb.org/mongo-driver/bson/primitive"
)

type UserService interface {
	Login(username, password string) (models.User, error)
	Register(username, password, email string) (models.User, error)
}

type DefaultUserService struct {
	Repo repositories.UserRepository
}

// düzeltilmesi gerek buranın imza uyumsuzluğu vs var.
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

func (s *DefaultUserService) Register(user models.User) error {
	// verify users models.
	if user.Username == "" || user.Password == "" || user.Email == "" {
		return errors.New("username, password, and email are required")
	}

	newUser := models.User{
		ID:        primitive.NewObjectID(),
		Username:  username,
		Password:  password,
		Email:     email,
		CreatedAt: time.Now().Format(time.RFC3339),
		UpdateAt:  time.Now().Format(time.RFC3339),
	}

	err := s.Repo.CreateUser(user)
	if err != nil {
		log.Printf("Error creating user: %v", err)
		return models.User{}, err

	}
	log.Println("User registered successfully: ", newUser.Username)
	return nil
}
