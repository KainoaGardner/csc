package db

import (
	"context"
	"github.com/KainoaGardner/csc/internal/config"
	"github.com/KainoaGardner/csc/internal/types"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
)

func CreateUserStats(client *mongo.Client, db config.DB, userStats *types.UserStats) (string, error) {
	collection := client.Database(db.Name).Collection(db.Collections.UserStats)

	userStats.ID = primitive.NewObjectID()
	_, err := collection.InsertOne(context.Background(), userStats)
	if err != nil {
		return "", err
	}

	return userStats.ID.Hex(), nil
}

func ListAllUserStats(client *mongo.Client, db config.DB) ([]types.UserStats, error) {
	var userStats []types.UserStats

	collection := client.Database(db.Name).Collection(db.Collections.UserStats)

	cursor, err := collection.Find(context.Background(), bson.M{})
	if err != nil {
	}

	err = cursor.All(context.Background(), &userStats)
	if err != nil {
		return nil, err
	}

	return userStats, nil
}

func DeleteAllUserStats(client *mongo.Client, db config.DB) (int, error) {
	collection := client.Database(db.Name).Collection(db.Collections.UserStats)
	result, err := collection.DeleteMany(context.Background(), bson.M{}, nil)
	if err != nil {
		return 0, err
	}

	return int(result.DeletedCount), nil
}

func FindUserStatsFromUserID(client *mongo.Client, db config.DB, userID string) (*types.UserStats, error) {
	var result types.UserStats

	id, err := primitive.ObjectIDFromHex(userID)
	if err != nil {
		return nil, err
	}

	filter := bson.M{"userID": id}

	collection := client.Database(db.Name).Collection(db.Collections.UserStats)
	err = collection.FindOne(context.Background(), filter).Decode(&result)
	if err != nil {
		return nil, err
	}

	return &result, nil
}

func UpdateUserStats(client *mongo.Client, db config.DB, userID string, userStatsUpdate types.UpdateUserStats) error {
	id, err := primitive.ObjectIDFromHex(userID)
	if err != nil {
		return err
	}

	filter := bson.M{"userID": id}
	update := bson.M{
		"$push": bson.M{
			"gameLogs": userStatsUpdate.GameLog,
		},
		"$set": bson.M{
			"gamesWon":    userStatsUpdate.GamesWon,
			"gamesPlayed": userStatsUpdate.GamesPlayed,
		},
	}

	collection := client.Database(db.Name).Collection(db.Collections.UserStats)
	_, err = collection.UpdateOne(context.Background(), filter, update)
	if err != nil {
		return err
	}

	return nil
}
