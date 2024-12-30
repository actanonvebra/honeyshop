// database queries for cart operations.
package repositories

import (
	"context"
	"time"

	"github.com/actanonvebra/honeyshop/internal/db"
	"github.com/actanonvebra/honeyshop/internal/models"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
)

type CartRepository interface {
	FindCartByUserID(userID string) (*models.Cart, error)
	DeleteCartByUserID(userID string) error
}

type MongoCartRepository struct {
	Collection *mongo.Collection
}

func NewMongoCartRepository(database, collection string) *MongoCartRepository {
	return &MongoCartRepository{
		Collection: db.GetCollection(database, collection),
	}
}

func (r *MongoCartRepository) FindCartByUserID(userID string) (*models.Cart, error) {
	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()

	var cart models.Cart
	filter := bson.M{"UserId": userID}
	err := r.Collection.FindOne(ctx, filter).Decode(&cart)
	if err != nil {
		return nil, err
	}
	return &cart, nil
}

func (r *MongoCartRepository) DeleteCartByUserID(userID string) error {
	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()

	filter := bson.M{"UserId": userID}
	_, err := r.Collection.DeleteOne(ctx, filter)
	return err
}
