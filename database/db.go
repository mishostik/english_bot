package database

import (
	"context"
	"english_bot/config"
	"fmt"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
	"log"
)

//type Database interface {} ?????????

type Database struct {
	client *mongo.Client
	dbase  *mongo.Database
}

func Connect(ctx context.Context, dbName string, cfg *config.Config) (*Database, error) {

	var (
		URL = fmt.Sprintf("mongodb://%s:%s@%s:%s", cfg.Mongo.User, cfg.Mongo.Password, cfg.Mongo.Host,
			cfg.Mongo.Port)
	)
	client, err := mongo.Connect(ctx, options.Client().ApplyURI(URL))
	if err != nil {
		log.Fatal("Error while connection to MongoDB")
	}
	db := client.Database(dbName)
	return &Database{
		client: client,
		dbase:  db,
	}, nil
}

func (db *Database) Collection(collectionName string) (*mongo.Collection, error) {
	return db.dbase.Collection(collectionName), nil
}

func (db *Database) Disconnect() error {
	return db.client.Disconnect(context.Background())
}
