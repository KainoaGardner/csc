package engine

import (
	"github.com/KainoaGardner/csc/internal/db"
	"github.com/KainoaGardner/csc/internal/types"
	// "github.com/KainoaGardner/csc/internal/utils"

	"github.com/KainoaGardner/csc/internal/config"
	"go.mongodb.org/mongo-driver/mongo"
)

func JoinGameCase(gameID string, userID string, client *mongo.Client, config config.Config) (*types.Game, error) {
	game, err := db.FindGame(client, config.DB, gameID)
	if err != nil {
		return nil, err
	}

	err = SetupJoinGame(game, userID)
	if err != nil {
		return nil, err
	}

	err = db.GameMoveUpdate(client, config.DB, gameID, *game)
	if err != nil {
		return nil, err
	}

	return game, nil
}

func MoveCase(gameID string, userID string, postMove types.PostMove, client *mongo.Client, config config.Config) (*types.Game, string, error) {
	game, err := db.FindGame(client, config.DB, gameID)
	if err != nil {
		return nil, "", err
	}

	turn, err := GetTurnFromID(*game, userID)
	if err != nil {
		return nil, "", err
	}

	err = CheckTurn(turn, game.Turn)
	if err != nil {
		return nil, "", err
	}

	move, err := ConvertStringToMove(postMove.Move, *game)
	if err != nil {
		return nil, "", err
	}

	err = MovePiece(move, game)
	if err != nil {
		return nil, "", err
	}

	err = db.GameMoveUpdate(client, config.DB, gameID, *game)
	if err != nil {
		return nil, "", err
	}

	fen, err := ConvertBoardToString(*game)
	if err != nil {
		return nil, "", err
	}

	err = db.GameLogUpdate(client, config.DB, gameID, postMove.Move, fen)
	if err != nil {
		return nil, "", err
	}

	return game, fen, nil
}

func PlaceCase(gameID string, userID string, postPlace types.PostPlace, client *mongo.Client, config config.Config) (*types.Game, types.PlaceResponse, error) {
	var result types.PlaceResponse

	game, err := db.FindGame(client, config.DB, gameID)
	if err != nil {
		return nil, result, err
	}

	turn, err := GetTurnFromID(*game, userID)
	if err != nil {
		return nil, result, err
	}

	var place types.Place
	//if place else delete placed piece
	if postPlace.Place {
		place, err = SetupPlace(postPlace, turn, *game)
		if err != nil {
			return nil, result, err
		}

		err = PlacePiece(place, game)
		if err != nil {
			return nil, result, err
		}
	} else {

		place, err = SetupDeletePlace(postPlace, turn, *game)
		if err != nil {
			return nil, result, err
		}
		err = PlacePieceDelete(&place, game)
		if err != nil {
			return nil, result, err
		}

	}

	err = db.GamePlaceUpdate(client, config.DB, gameID, place, *game)
	if err != nil {
		return nil, result, err
	}

	fen, err := ConvertBoardToString(*game)
	if err != nil {
		return nil, result, err
	}

	var cost int
	if postPlace.Place {
		cost = place.Cost
	} else {
		cost = place.Cost * -1
	}

	result = types.PlaceResponse{
		ID:       game.ID,
		FEN:      fen,
		Position: postPlace.Position,
		Type:     postPlace.Type,
		Cost:     cost,
		Money:    game.Money,
	}

	return game, result, nil
}

func ReadyCase(gameID string, userID string, postReady types.PostReady, client *mongo.Client, config config.Config) (*types.Game, error) {
	game, err := db.FindGame(client, config.DB, gameID)
	if err != nil {
		return nil, err
	}

	turn, err := GetTurnFromID(*game, userID)
	if err != nil {
		return nil, err
	}

	err = ReadyPlayer(postReady.Ready, turn, game)
	if err != nil {
		return nil, err
	}

	err = db.GameReadyUpdate(client, config.DB, gameID, *game)
	if err != nil {
		return nil, err
	}

	return game, nil
}

func DrawCase(gameID string, userID string, postDraw types.PostDrawRequest, client *mongo.Client, config config.Config) (*types.Game, error) {
	game, err := db.FindGame(client, config.DB, gameID)
	if err != nil {
		return nil, err
	}

	turn, err := GetTurnFromID(*game, userID)
	if err != nil {
		return nil, err
	}

	err = DrawRequest(postDraw.Draw, turn, game)
	if err != nil {
		return nil, err
	}

	err = db.GameDrawUpdate(client, config.DB, gameID, *game)
	if err != nil {
		return nil, err
	}

	return game, nil
}

func ResignCase(gameID string, userID string, client *mongo.Client, config config.Config) (*types.Game, error) {
	game, err := db.FindGame(client, config.DB, gameID)
	if err != nil {
		return nil, err
	}

	turn, err := GetTurnFromID(*game, userID)
	if err != nil {
		return nil, err
	}

	SetupResignGame(game, turn)

	err = db.GameMoveUpdate(client, config.DB, gameID, *game)
	if err != nil {
		return nil, err
	}

	return game, nil
}
