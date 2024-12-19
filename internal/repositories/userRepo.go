// user database operations. (login, registration, etc.)
package repositories

import (
	"context"
	"log"
	"time"

	"github.com/actanonvebra/honeyshop/internal/db"
	"github.com/actanonvebra/honeyshop/internal/models"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
)

type MongoUserRepo struct {
	Collection *mongo.Collection
}

type UserRepository interface {
	GetUserByUserName(username string) (models.User, error)
	CreateUser(user models.User) error
}

func NewMongoUserRepo(database, collection string) *MongoUserRepo {
	return &MongoUserRepo{
		Collection: db.GetCollection(database, collection),
	}
}

func (repo *MongoUserRepo) GetUserByUserName(username string) (models.User, error) {
	var user models.User
	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()
	filter := bson.M{"username": username}
	err := repo.Collection.FindOne(ctx, filter).Decode(&user)
	return user, err
}

// func (repo *MongoUserRepo) CreateUser(user models.User) error {
// 	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
// 	defer cancel()
// 	_, err := repo.Collection.InsertOne(ctx, user)
// 	return err
// }

func (repo *MongoUserRepo) CreateUser(user models.User) error {
	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()

	//automatic creation ObjectID
	if user.ID.IsZero() {
		user.ID = primitive.NewObjectID()
	}

	// add mongodb user collection
	_, err := repo.Collection.InsertOne(ctx, user)
	if err != nil {
		log.Printf("Erorr instering user into MongoDB: %v", err)
		return err
	}

	log.Println("User added to MongoDB: ", user.Username)
	return nil
}
