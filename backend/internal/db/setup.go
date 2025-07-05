package db

import (
	"context"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
	"log"

	"github.com/KainoaGardner/csc/internal/config"
)

func ConnectToDB(context context.Context, db config.DB) (*mongo.Client, error) {
	client, err := mongo.Connect(context, options.Client().ApplyURI(db.Uri))
	if err != nil {
		return nil, err
	}

	err = client.Ping(context, nil)
	if err != nil {
		return nil, err
	}

	log.Println("Connected to MongoDB")
	return client, nil
}
