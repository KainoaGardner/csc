package db

import (
	"context"
	"github.com/KainoaGardner/csc/internal/config"
	"github.com/KainoaGardner/csc/internal/types"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
)

func CreateGameLog(client *mongo.Client, db config.DB, gameLog *types.GameLog) (string, error) {
	collection := client.Database(db.Name).Collection(db.Collections.GameLogs)

	gameLog.ID = primitive.NewObjectID()
	_, err := collection.InsertOne(context.Background(), gameLog)
	if err != nil {
		return "", err
	}

	return gameLog.ID.Hex(), nil
}

func ListAllGameLogs(client *mongo.Client, db config.DB) ([]types.GameLog, error) {
	var gameLogs []types.GameLog

	collection := client.Database(db.Name).Collection(db.Collections.GameLogs)

	cursor, err := collection.Find(context.Background(), bson.M{})
	if err != nil {
	}

	err = cursor.All(context.Background(), &gameLogs)
	if err != nil {
		return nil, err
	}

	return gameLogs, nil
}

func DeleteAllGameLogs(client *mongo.Client, db config.DB) (int, error) {
	collection := client.Database(db.Name).Collection(db.Collections.GameLogs)
	result, err := collection.DeleteMany(context.Background(), bson.M{}, nil)
	if err != nil {
		return 0, err
	}

	return int(result.DeletedCount), nil
}

func FindGameLog(client *mongo.Client, db config.DB, gameLogID string) (*types.GameLog, error) {
	var result types.GameLog

	id, err := primitive.ObjectIDFromHex(gameLogID)
	if err != nil {
		return nil, err
	}

	filter := bson.M{"_id": id}

	collection := client.Database(db.Name).Collection(db.Collections.GameLogs)
	err = collection.FindOne(context.Background(), filter).Decode(&result)
	if err != nil {
		return nil, err
	}

	return &result, nil
}

func FindGameLogFromGameID(client *mongo.Client, db config.DB, gameID string) (*types.GameLog, error) {
	var result types.GameLog

	id, err := primitive.ObjectIDFromHex(gameID)
	if err != nil {
		return nil, err
	}

	filter := bson.M{"gameID": id}

	collection := client.Database(db.Name).Collection(db.Collections.GameLogs)
	err = collection.FindOne(context.Background(), filter).Decode(&result)
	if err != nil {
		return nil, err
	}

	return &result, nil
}

func GameLogUpdate(client *mongo.Client, db config.DB, gameID string, moveString string, fenString string) error {
	id, err := primitive.ObjectIDFromHex(gameID)
	if err != nil {
		return err
	}

	filter := bson.M{"gameID": id}
	update := bson.M{"$push": bson.M{"moves": moveString, "boardStates": fenString}}

	collection := client.Database(db.Name).Collection(db.Collections.GameLogs)
	_, err = collection.UpdateOne(context.Background(), filter, update)
	if err != nil {
		return err
	}

	return nil
}

func GameLogFinalUpdate(client *mongo.Client, db config.DB, gameID string, gameLog types.GameLog) error {
	id, err := primitive.ObjectIDFromHex(gameID)
	if err != nil {
		return err
	}

	filter := bson.M{"_id": id}
	update := bson.M{"$set": gameLog}

	collection := client.Database(db.Name).Collection(db.Collections.GameLogs)
	_, err = collection.UpdateOne(context.Background(), filter, update)
	if err != nil {
		return err
	}

	return nil
}
