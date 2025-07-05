package engine

import (
	"fmt"
	"github.com/KainoaGardner/csc/internal/types"
)

func PlacePiece(place types.Place, game *types.Game) error {
	err := checkGameState(types.PlaceState, game.State)
	if err != nil {
		return err
	}

	err = checkValidPlace(place, *game)
	if err != nil {
		return err
	}

	updatePlacePiece(place, game)
	return nil
}

func PlacePieceDelete(place types.Place, game *types.Game) error {
	err := checkGameState(types.PlaceState, game.State)
	if err != nil {
		return err
	}

	err = checkValidPlaceDelete(place, *game)
	if err != nil {
		return err
	}

	updateDeletePlacePiece(place, game)

	return nil
}

func updatePlacePiece(place types.Place, game *types.Game) {
	piece := types.Piece{}
	piece.Owner = place.Turn
	piece.Type = place.Type

	game.Board.Board[place.Pos.Y][place.Pos.X] = &piece

	//set to actual cost
	game.Money[place.Turn] -= place.Cost
}

func checkValidPlace(place types.Place, game types.Game) error {
	err := checkEnoughMoney(place, game)
	if err != nil {
		return err
	}

	err = checkPlaceInBounds(place, game)
	if err != nil {
		return err
	}

	err = checkPlaceOnYourSide(place, game)
	if err != nil {
		return err
	}

	err = checkEmptyPlaceSpace(place, game)
	if err != nil {
		return err
	}

	return nil
}

func updateDeletePlacePiece(place types.Place, game *types.Game) error {
	piece := game.Board.Board[place.Pos.Y][place.Pos.X]
	if piece == nil {
		return fmt.Errorf("Cannot delete empty piece")
	}

	place.Type = piece.Type

	cost, ok := types.PieceToCost[place.Type]
	if !ok {
		return fmt.Errorf("Could not get cost of piece")
	}

	place.Cost = cost

	game.Board.Board[place.Pos.Y][place.Pos.X] = nil
	game.Money[place.Turn] += place.Cost
	return nil
}

func checkValidPlaceDelete(place types.Place, game types.Game) error {
	err := checkPlaceInBounds(place, game)
	if err != nil {
		return err
	}

	err = checkPlayerPiece(place, game)
	if err != nil {
		return err
	}

	return nil
}

func checkPlayerPiece(place types.Place, game types.Game) error {
	piece := game.Board.Board[place.Pos.Y][place.Pos.X]
	if piece == nil {
		return fmt.Errorf("Cannot delete empty piece")
	}

	if piece.Owner != place.Turn {
		return fmt.Errorf("Cannot delete opponents piece")
	}

	return nil
}

func checkEnoughMoney(place types.Place, game types.Game) error {
	if game.Money[place.Turn]-place.Cost < 0 {
		return fmt.Errorf("Not enough money")
	}

	return nil
}

func checkPlaceInBounds(place types.Place, game types.Game) error {
	if place.Pos.X < 0 || place.Pos.X >= game.Board.Width {
		return fmt.Errorf("Place x out of board bounds")
	}
	if place.Pos.Y < 0 || place.Pos.Y >= game.Board.Height {
		return fmt.Errorf("Place y out of board bounds")
	}

	return nil
}

func checkPlaceOnYourSide(place types.Place, game types.Game) error {
	if place.Turn == 0 && place.Pos.Y < game.Board.PlaceLine {
		return fmt.Errorf("Cannot place on opponents side")
	}
	if place.Turn == 1 && place.Pos.Y >= game.Board.PlaceLine {
		return fmt.Errorf("Cannot place on opponents side")
	}

	return nil
}

func checkEmptyPlaceSpace(place types.Place, game types.Game) error {
	if game.Board.Board[place.Pos.Y][place.Pos.X] != nil {
		return fmt.Errorf("Cannot place piece on another piece")
	}

	return nil
}

func checkValidPlaceType(place types.PostPlace) error {
	if place.Type >= types.Pawn && place.Type <= types.King {
		return nil
	}
	if place.Type >= types.Fu && place.Type <= types.Ou {
		return nil
	}
	if place.Type >= types.Checker && place.Type <= types.CheckerKing {
		return nil
	}

	return fmt.Errorf("Invalid place type")
}

func SetupPlace(placeConfig types.PostPlace, userID string, game types.Game) (types.Place, error) {
	var result types.Place

	turn, err := GetTurnFromID(userID, game) //USE id
	if err != nil {
		return result, fmt.Errorf("Could not get turn from player ID")
	}
	result.Turn = turn

	err = checkValidPlaceType(placeConfig)
	if err != nil {
		return result, err
	}
	result.Type = placeConfig.Type

	position, err := convertStringToPosition(placeConfig.Position, game.Board.Height)
	if err != nil {
		return result, err
	}
	result.Pos = position

	cost, ok := types.PieceToCost[result.Type]
	if !ok {
		return result, fmt.Errorf("Could not get cost of piece")
	}
	result.Cost = cost

	return result, nil
}

func SetupDeletePlace(placeConfig types.DeletePlace, userID string, game types.Game) (types.Place, error) {
	var result types.Place

	turn, err := GetTurnFromID(userID, game) //USE id
	if err != nil {
		return result, fmt.Errorf("Could not get turn from player ID")
	}
	result.Turn = turn

	position, err := convertStringToPosition(placeConfig.Position, game.Board.Height)
	if err != nil {
		return result, err
	}
	result.Pos = position

	return result, nil
}
