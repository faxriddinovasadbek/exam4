package db

import (
	"context"
	"fmt"
	"post-service/config"

	"github.com/jmoiron/sqlx"
	_ "github.com/lib/pq" //postgres drivers
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

func ConnectToDB(cfg config.Config) (*sqlx.DB, error, func()) {
	psqlString := fmt.Sprintf("host=%s port=%d user=%s password=%s dbname=%s sslmode=disable",
		cfg.PostgresHost,
		cfg.PostgresPort,
		cfg.PostgresUser,
		cfg.PostgresPassword,
		cfg.PostgresDatabase,
	)

	connDb, err := sqlx.Connect("postgres", psqlString)
	if err != nil {
		return nil, err, nil
	}

	cleanUpFunc := func() {
		connDb.Close()
	}

	return connDb, nil, cleanUpFunc
}

func ConnectToMongoDB(cfg config.Config) (*mongo.Collection, error) {
	clientOptions := options.Client().ApplyURI(fmt.Sprintf("mongodb://%s:%d", cfg.MongoHost, cfg.MongoPort))
	client, err := mongo.Connect(context.Background(), clientOptions)
	if err != nil {
		return nil, err
	}
	collection := client.Database(cfg.MongoDatabase).Collection("posts")
	return collection, nil
}

func ConnectToDBForSuite(cfg config.Config) (*sqlx.DB, func()) {
	psqlString := fmt.Sprintf("host=%s port=%d user=%s password=%s dbname=%s sslmode=disable",
		cfg.PostgresHost,
		cfg.PostgresPort,
		cfg.PostgresUser,
		cfg.PostgresPassword,
		cfg.PostgresDatabase,
	)
	connDB, err := sqlx.Connect("postgres", psqlString)
	if err != nil {
		panic(err)
	}
	cleanUpFunc := func() {
		connDB.Close()
	}
	return connDB, cleanUpFunc
}

func ConnectToMongoDBForSuite(cfg config.Config) (*mongo.Collection, func()) {
	clientOptions := options.Client().ApplyURI(fmt.Sprintf("mongodb://%s:%d", cfg.MongoHost, cfg.MongoPort))
	client, err := mongo.Connect(context.Background(), clientOptions)
	if err != nil {
		return nil, nil
	}
	collection := client.Database(cfg.MongoDatabase).Collection("posts")
	cleanUpFunc := func() {
		collection.Clone()
	}
	return collection, cleanUpFunc
}
