package database

import (
	"context"
	"fmt"
	"github.com/airchains-network/tracks-espresso-client/config"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

// DB struct encapsulates the MongoDB database and collection
type DB struct {
	*mongo.Database
}

// InitConnection initializes MongoDB connection and sets up the collection
func InitConnection() (*DB, error) {
	clientOptions := options.Client().ApplyURI(config.MongoUrl)
	client, err := mongo.Connect(context.TODO(), clientOptions)
	if err != nil {
		return nil, fmt.Errorf("database connection failed: %s", err)
	}

	err = client.Ping(context.TODO(), nil)
	if err != nil {
		return nil, fmt.Errorf("database connection failed: %s", err)
	}

	db := client.Database("espresso")

	return &DB{
		Database: db,
	}, nil
}
