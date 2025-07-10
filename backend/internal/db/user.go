package db

import (
	"context"
	"github.com/KainoaGardner/csc/internal/config"
	"github.com/KainoaGardner/csc/internal/types"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
)

func FindUserLogin(client *mongo.Client, db config.DB, user types.User) error {
	var result bson.M

	filter := bson.M{
		"$or": []bson.M{
			{"username": user.Username},
			{"email": user.Email},
		},
	}

	collection := client.Database(db.Name).Collection(db.Collections.Users)
	err := collection.FindOne(context.Background(), filter).Decode(&result)
	if err != nil {
		return err
	}

	return nil
}

func CreateUser(client *mongo.Client, db config.DB, newUser *types.User) (string, error) {
	collection := client.Database(db.Name).Collection(db.Collections.Users)

	newUser.ID = primitive.NewObjectID()
	_, err := collection.InsertOne(context.Background(), newUser)
	if err != nil {
		return "", err
	}

	return newUser.ID.Hex(), nil
}

func ListAllUsers(client *mongo.Client, db config.DB) ([]types.User, error) {
	var users []types.User

	collection := client.Database(db.Name).Collection(db.Collections.Users)

	cursor, err := collection.Find(context.Background(), bson.M{})
	if err != nil {
	}

	err = cursor.All(context.Background(), &users)
	if err != nil {
		return nil, err
	}

	return users, nil
}

func DeleteAllUsers(client *mongo.Client, db config.DB) (int, error) {
	collection := client.Database(db.Name).Collection(db.Collections.Users)
	result, err := collection.DeleteMany(context.Background(), bson.M{}, nil)
	if err != nil {
		return 0, err
	}

	return int(result.DeletedCount), nil
}
