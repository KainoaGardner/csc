package engine

import "fmt"

func checkValidDrop(move Move, piece Piece, game Game) error {
	err := checkHaveDropPiece(move, game)
	if err != nil {
		return err
	}

	err = checkEmptySpace(move, game)
	if err != nil {
		return err
	}

	err = checkPromoteDrop(move)
	if err != nil {
		return err
	}
	err = checkNifu(move, piece, game)
	if err != nil {
		return err
	}
	err = checkIkidokoronoNaiKoma(move, piece, game)
	if err != nil {
		return err
	}
	err = checkUtifudume(move, piece, game) //NOT DONE
	if err != nil {
		return err
	}

	return nil
}

func checkHaveDropPiece(move Move, game Game) error {
	offset := getMochigomaOffset(game)

	if move.Drop == nil {
		return fmt.Errorf("Drop not set")
	}

	komaCount := game.Mochigoma[*move.Drop+offset]
	if komaCount <= 0 {
		return fmt.Errorf("Not enough mochigoma")
	}

	return nil
}

func checkEmptySpace(move Move, game Game) error {
	space := game.Board.Board[move.End.Y][move.End.X]
	if space != nil {
		return fmt.Errorf("Cant drop on non empty space")
	}

	return nil
}

func checkPromoteDrop(move Move) error {
	if move.Drop != nil && move.Promote != nil {
		return fmt.Errorf("Cant promote when dropping")
	}
	return nil
}

func checkNifu(move Move, piece Piece, game Game) error {
	if piece.Type != Fu {
		return nil
	}

	for i := 0; i < game.Board.Height; i++ {
		space := game.Board.Board[i][move.End.X]
		if space != nil && space.Type == Fu && space.Owner == piece.Owner {
			return fmt.Errorf("Cant place Fu in row with Fu. Nifu")
		}
	}

	return nil
}

func checkIkidokoronoNaiKoma(move Move, piece Piece, game Game) error {
	var row0 int
	var row1 int
	if piece.Owner == 0 {
		row0 = 0
		row1 = 1
	} else {
		row0 = game.Board.Height - 1
		row1 = game.Board.Height - 2
	}

	if piece.Type == Fu || piece.Type == Kyou {
		if move.End.Y == row0 {
			return fmt.Errorf("Can drop piece with no move")
		}
	} else if piece.Type == Kei {
		if move.End.Y == row0 || move.End.Y == row1 {
			return fmt.Errorf("Can drop piece with no move")
		}
	}

	return nil
}

func checkUtifudume(move Move, piece Piece, game Game) error {
	if piece.Type != Fu {
		return nil
	}

	//CHANGE DO if move makes a checkmate
	if 0 == 1 {
		return fmt.Errorf("Cant checkmate with fu drop")
	}

	return nil
}

func getMochigomaOffset(game Game) int {
	offset := 0
	if game.Turn == 1 {
		offset = 7
	}

	return offset
}
