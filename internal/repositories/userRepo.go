// user database operations. (login, registration, etc.)
package repositories

import "github.com/actanonvebra/honeyshop/internal/models"

type UserRepository interface {
	GetUserByUserName(username string) (models.User, error)
	CreateUser(user models.User) error
}

type MockUserRepo struct{}

func (repo *MockUserRepo) GetUserByUserName(username string) (models.User, error) {

	if username == "admin" {
		return models.User{ID: 1, Username: "admin", Password: "1234"}, nil
	}
	return models.User{}, nil
}

func (repo *MockUserRepo) CreateUser(user models.User) error {
	return nil
}
