package db

import (
	"context"
	"github.com/KainoaGardner/csc/internal/config"
	"github.com/KainoaGardner/csc/internal/types"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"

	"fmt"
)

func CreateGame(client *mongo.Client, db config.DB, game *types.Game) (string, error) {
	collection := client.Database(db.Name).Collection(db.Collections.Games)

	game.ID = primitive.NewObjectID()
	_, err := collection.InsertOne(context.Background(), game)
	if err != nil {
		return "", err
	}

	return game.ID.Hex(), nil
}

func ListAllGames(client *mongo.Client, db config.DB) ([]types.Game, error) {
	var games []types.Game

	collection := client.Database(db.Name).Collection(db.Collections.Games)

	cursor, err := collection.Find(context.Background(), bson.M{})
	if err != nil {
	}

	err = cursor.All(context.Background(), &games)
	if err != nil {
		return nil, err
	}

	return games, nil
}

func DeleteAllGames(client *mongo.Client, db config.DB) (int, error) {
	collection := client.Database(db.Name).Collection(db.Collections.Games)
	result, err := collection.DeleteMany(context.Background(), bson.M{}, nil)
	if err != nil {
		return 0, err
	}

	return int(result.DeletedCount), nil
}

func FindGame(client *mongo.Client, db config.DB, gameID string) (*types.Game, error) {
	var result types.Game

	id, err := primitive.ObjectIDFromHex(gameID)
	if err != nil {
		return nil, err
	}

	filter := bson.M{"_id": id}

	collection := client.Database(db.Name).Collection(db.Collections.Games)
	err = collection.FindOne(context.Background(), filter).Decode(&result)
	if err != nil {
		return nil, err
	}

	return &result, nil
}

func GamePlaceUpdate(client *mongo.Client, db config.DB, gameID string, place types.Place, game types.Game) error {
	id, err := primitive.ObjectIDFromHex(gameID)
	if err != nil {
		return err
	}

	replaceString := fmt.Sprintf("board.board.%d.%d", place.Pos.Y, place.Pos.X)

	filter := bson.M{"_id": id}
	update := bson.M{"$set": bson.M{replaceString: game.Board.Board[place.Pos.Y][place.Pos.X], "money": game.Money}}

	collection := client.Database(db.Name).Collection(db.Collections.Games)
	_, err = collection.UpdateOne(context.Background(), filter, update)
	if err != nil {
		return err
	}

	return nil
}

func GameMoveUpdate(client *mongo.Client, db config.DB, gameID string, game types.Game) error {
	id, err := primitive.ObjectIDFromHex(gameID)
	if err != nil {
		return err
	}

	filter := bson.M{"_id": id}
	update := bson.M{"$set": game}

	collection := client.Database(db.Name).Collection(db.Collections.Games)
	_, err = collection.UpdateOne(context.Background(), filter, update)
	if err != nil {
		return err
	}

	return nil
}

func GameStateUpdate(client *mongo.Client, db config.DB, gameID string, game types.Game) error {
	id, err := primitive.ObjectIDFromHex(gameID)
	if err != nil {
		return err
	}

	filter := bson.M{"_id": id}
	update := bson.M{"$set": bson.M{"state": game.State}}

	collection := client.Database(db.Name).Collection(db.Collections.Games)
	_, err = collection.UpdateOne(context.Background(), filter, update)
	if err != nil {
		return err
	}

	return nil
}

func GameReadyUpdate(client *mongo.Client, db config.DB, gameID string, game types.Game) error {
	id, err := primitive.ObjectIDFromHex(gameID)
	if err != nil {
		return err
	}

	filter := bson.M{"_id": id}
	update := bson.M{"$set": bson.M{"state": game.State, "ready": game.Ready, "lastMoveTime": game.LastMoveTime}}

	collection := client.Database(db.Name).Collection(db.Collections.Games)
	_, err = collection.UpdateOne(context.Background(), filter, update)
	if err != nil {
		return err
	}

	return nil
}

func GameDrawUpdate(client *mongo.Client, db config.DB, gameID string, game types.Game) error {
	id, err := primitive.ObjectIDFromHex(gameID)
	if err != nil {
		return err
	}

	filter := bson.M{"_id": id}
	update := bson.M{"$set": bson.M{"draw": game.Draw}}

	collection := client.Database(db.Name).Collection(db.Collections.Games)
	_, err = collection.UpdateOne(context.Background(), filter, update)
	if err != nil {
		return err
	}

	return nil
}
