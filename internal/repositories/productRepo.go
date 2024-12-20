// product database operations.
package repositories

import (
	"context"
	"time"

	"github.com/actanonvebra/honeyshop/internal/db"
	"github.com/actanonvebra/honeyshop/internal/models"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
)

type ProductRepository interface {
	GetAllProducts() ([]models.Product, error)
}
type MongoProductRepo struct {
	Collection *mongo.Collection
}

func NewMongoProductRepo(database, collection string) *MongoProductRepo {
	return &MongoProductRepo{
		Collection: db.GetCollection(database, collection),
	}
}

func (repo *MongoProductRepo) GetAllProducts() ([]models.Product, error) {
	var products []models.Product
	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()

	cursor, err := repo.Collection.Find(ctx, bson.M{})
	if err != nil {
		return nil, err
	}
	defer cursor.Close(ctx)

	for cursor.Next(ctx) {
		var product models.Product
		if err := cursor.Decode(&product); err != nil {
			return nil, err
		}
		products = append(products, product)
	}
	return products, nil

}
