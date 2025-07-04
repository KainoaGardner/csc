package engine

import (
	"fmt"
	"github.com/KainoaGardner/csc/internal/types"
)

func checkValidPromote(move types.Move, piece types.Piece, game types.Game) error {
	switch piece.Type {
	case types.Pawn:
		err := checkPawnCheckerPromote(move, piece, game)
		if err != nil {
			return err
		}
	case types.Checker:
		err := checkPawnCheckerPromote(move, piece, game)
		if err != nil {
			return err
		}
	default:
		err := checkShogiPromote(move, piece, game)
		if err != nil {
			return err
		}
	}

	return nil
}

func checkPawnCheckerPromote(move types.Move, piece types.Piece, game types.Game) error {
	var row int
	if piece.Owner == 0 {
		row = 0
	} else {
		row = game.Board.Height - 1
	}

	if move.End.Y != row {
		return fmt.Errorf("Can only promote on last row")
	}

	return nil
}

func checkShogiPromote(move types.Move, piece types.Piece, game types.Game) error {
	var rowStart, rowEnd int
	if piece.Owner == 0 {
		rowStart = 0
		rowEnd = 2
	} else {
		rowStart = game.Board.Height - 3
		rowEnd = game.Board.Height - 1
	}

	if move.Start.Y >= rowStart && move.Start.Y <= rowEnd {
		return nil
	}
	if move.End.Y >= rowStart && move.End.Y <= rowEnd {
		return nil
	}

	return fmt.Errorf("Must Move in promotion zone to promote")
}

func checkMustPromote(move types.Move, piece types.Piece, game types.Game) error {
	if piece.Type == types.Pawn || piece.Type == types.Checker || piece.Type == types.Fu || piece.Type == types.Kyou {
		err := checkMustPromoteLastLine(move, piece, game)
		if err != nil {
			return err
		}
	} else if piece.Type == types.Kei {
		err := checkMustPromoteLastTwoLines(move, piece, game)
		if err != nil {
			return err
		}
	}
	return nil
}

func checkMustPromoteLastLine(move types.Move, piece types.Piece, game types.Game) error {
	var row0 int
	if piece.Owner == 0 {
		row0 = 0
	} else {
		row0 = game.Board.Height - 1
	}

	if move.End.Y == row0 {
		return fmt.Errorf("Must promote and last row")
	}

	return nil
}

func checkMustPromoteLastTwoLines(move types.Move, piece types.Piece, game types.Game) error {
	var row0 int
	var row1 int
	if piece.Owner == 0 {
		row0 = 0
		row1 = 1
	} else {
		row0 = game.Board.Height - 1
		row1 = game.Board.Height - 2
	}

	if move.End.Y == row0 && move.End.Y == row1 {
		return fmt.Errorf("Must promote and last row")
	}

	return nil
}
