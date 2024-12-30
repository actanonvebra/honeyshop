// product database operations.
package repositories

import (
	"context"
	"log"
	"time"

	"github.com/actanonvebra/honeyshop/internal/db"
	"github.com/actanonvebra/honeyshop/internal/models"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
)

type ProductRepository interface {
	GetAllProducts() ([]models.Product, error)
	SearchProducts(keyword string) ([]models.Product, error)
	AddProduct(product models.Product) error
	FindProductByID(productID string) (*models.Product, error)
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

func (repo *MongoProductRepo) SearchProducts(keyword string) ([]models.Product, error) {
	var products []models.Product
	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()

	filter := bson.M{"name": bson.M{"$regex": keyword, "$options": "i"}}
	cursor, err := repo.Collection.Find(ctx, filter)
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

func (repo *MongoProductRepo) AddProduct(product models.Product) error {
	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()

	_, err := repo.Collection.InsertOne(ctx, product)
	if err != nil {
		log.Print("Added is not possible")
		return err
	}
	return nil
}

func (r *MongoProductRepo) FindProductByID(productID string) (*models.Product, error) {
	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()

	var product models.Product
	filter := bson.M{"id": productID}
	err := r.Collection.FindOne(ctx, filter).Decode(&product)
	if err != nil {
		return nil, err
	}
	return &product, nil
}
