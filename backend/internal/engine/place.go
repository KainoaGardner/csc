package engine

import (
	"fmt"
	"github.com/KainoaGardner/csc/internal/types"
)

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

func updatePlacePiece(place types.Place, game *types.Game) {
	piece := types.Piece{}
	piece.Owner = place.Turn
	piece.Type = place.Type

	game.Board.Board[place.Pos.Y][place.Pos.X] = &piece

	//set to actual cost
	game.Money[place.Turn] -= place.Cost
}
