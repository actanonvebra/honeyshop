package repositories

import (
	"context"
	"log"
	"time"

	"github.com/actanonvebra/honeyshop/internal/db"
	"github.com/actanonvebra/honeyshop/internal/models"
	"go.mongodb.org/mongo-driver/mongo"
)

type LogRepository interface {
	LogAttack(attackType, details, ip string) error
}

type MongoLogRepo struct {
	Collection *mongo.Collection
}

func NewMongoLogRepo(database, collection string) *MongoLogRepo {
	return &MongoLogRepo{Collection: db.GetCollection(database, collection)}
}

func (repo *MongoLogRepo) LogAttack(attackType, details, ip string) error {
	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()

	logEntry := models.AttackLog{
		Type:      attackType,
		Details:   details,
		IP:        ip,
		Timestamp: time.Now().Format(time.RFC3339),
	}

	_, err := repo.Collection.InsertOne(ctx, logEntry)
	if err != nil {
		log.Printf("Failed to log attack: %v", err)
		return err
	}
	log.Printf("Attack logged: %+v", logEntry)
	return nil
}
