// payment information processing queries.
package repositories

import (
	"context"
	"time"

	"github.com/actanonvebra/honeyshop/internal/db"
	"github.com/actanonvebra/honeyshop/internal/models"
	"go.mongodb.org/mongo-driver/mongo"
)

type CheckoutRepository interface {
	SaveCheckout(checkout models.Checkout) error
}

type MongoCheckoutRepo struct {
	Collection *mongo.Collection
}

func NewMongoCheckoutRepo(database, collection string) *MongoCheckoutRepo {
	return &MongoCheckoutRepo{
		Collection: db.GetCollection(database, collection),
	}
}

func (repo *MongoCheckoutRepo) SaveCheckout(checkout models.Checkout) error {
	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()

	_, err := repo.Collection.InsertOne(ctx, checkout)
	return err
}
