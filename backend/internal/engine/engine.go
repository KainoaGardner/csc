package engine

import "fmt"

// in bounds
// moving your piece
// correct piece movement
// if promotiion check correct promotion
// if drop check drop
// if in check cant move unless not in check after

func CheckValidMove(moves *[]Move, game Game) error {
	moveCount := len(*moves)
	if moveCount == 1 {
		move := (*moves)[0]
		err := checkMoveInBounds(move, game)
		if err != nil {
			return err
		}

		piece, err := getPiece(move, game)
		if err != nil {
			return err
		}

		err = checkValidPieceMoves(move, *piece, game)
		if err != nil {
			return err
		}
	}

	return nil
}

func checkMoveInBounds(move Move, game Game) error {
	if move.Start.X < 0 || move.Start.X >= game.Board.Width {
		return fmt.Errorf("Start x out of board bounds")
	}
	if move.Start.Y < 0 || move.Start.Y >= game.Board.Height {
		return fmt.Errorf("Start y out of board bounds")
	}
	if move.End.X < 0 || move.End.X >= game.Board.Width {
		return fmt.Errorf("End x out of board bounds")
	}
	if move.End.Y < 0 || move.End.Y >= game.Board.Height {
		return fmt.Errorf("End y out of board bounds")
	}

	return nil
}

func getPiece(move Move, game Game) (*Piece, error) {
	var piece *Piece
	if move.Drop != nil {
		var dropPiece Piece
		dropPiece.Type = *move.Drop
		dropPiece.Owner = game.Turn
		piece = &dropPiece

	} else {
		piece = game.Board.Board[move.Start.Y][move.Start.X]
		if piece == nil {
			return piece, fmt.Errorf("Cant move empty piece")
		}

		if piece.Owner != game.Turn {
			return piece, fmt.Errorf("Cant move other players piece")
		}
	}

	return piece, nil
}

func checkValidPieceMoves(move Move, piece Piece, game Game) error {
	//check valid drop

	//check valid normal move
	//in normal moves if move has promotion check that

	var result error
	dir := getMoveDirection(game)

	possibleMoves := []Vec2{}

	if move.Drop != nil {
		//drop check
		err := checkEmptySpace(move, game)
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
		err = checkUtifudume(move, piece, game)
		if err != nil {
			return err
		}

	}

	switch piece.Type {
	case Pawn: //add en passant
		possibleMoves = getPawnMoves(move, piece, game, dir)
		result = checkEndPosInPossibleMoves(possibleMoves, move)
	case Knight:
		possibleMoves = getKnightMoves(move, piece, game)
		result = checkEndPosInPossibleMoves(possibleMoves, move)
	case Bishop:
		possibleMoves = getBishopMoves(move, piece, game)
		result = checkEndPosInPossibleMoves(possibleMoves, move)
	case Rook:
		possibleMoves = getRookMoves(move, piece, game)
		result = checkEndPosInPossibleMoves(possibleMoves, move)
	case Queen:
		possibleMoves = getQueenMoves(move, piece, game)
		result = checkEndPosInPossibleMoves(possibleMoves, move)
	case King:
		possibleMoves = getKingMoves(move, piece, game)
		result = checkEndPosInPossibleMoves(possibleMoves, move)
	case Fu:
		possibleMoves = getFuMoves(move, piece, game, dir)
		result = checkEndPosInPossibleMoves(possibleMoves, move)
	case Kyou:
		possibleMoves = getKyouMoves(move, piece, game, dir)
		result = checkEndPosInPossibleMoves(possibleMoves, move)
	case Kei:
		possibleMoves = getKeiMoves(move, piece, game, dir)
		result = checkEndPosInPossibleMoves(possibleMoves, move)
	case Gin:
		possibleMoves = getGinMoves(move, piece, game, dir)
		result = checkEndPosInPossibleMoves(possibleMoves, move)
	case Kin:
		possibleMoves = getKinMoves(move, piece, game, dir)
		result = checkEndPosInPossibleMoves(possibleMoves, move)
	case Kaku:
		possibleMoves = getBishopMoves(move, piece, game)
		result = checkEndPosInPossibleMoves(possibleMoves, move)
	case Hi:
		possibleMoves = getRookMoves(move, piece, game)
		result = checkEndPosInPossibleMoves(possibleMoves, move)
	case Ou:
		possibleMoves = getKingMoves(move, piece, game)
		result = checkEndPosInPossibleMoves(possibleMoves, move)
	case To:
		possibleMoves = getKinMoves(move, piece, game, dir)
		result = checkEndPosInPossibleMoves(possibleMoves, move)
	case NariKyou:
		possibleMoves = getKinMoves(move, piece, game, dir)
		result = checkEndPosInPossibleMoves(possibleMoves, move)
	case NariKei:
		possibleMoves = getKinMoves(move, piece, game, dir)
		result = checkEndPosInPossibleMoves(possibleMoves, move)
	case NariGin:
		possibleMoves = getKinMoves(move, piece, game, dir)
		result = checkEndPosInPossibleMoves(possibleMoves, move)
	case Uma:
		possibleMoves = getUmaMoves(move, piece, game)
		result = checkEndPosInPossibleMoves(possibleMoves, move)
	case Ryuu:
		possibleMoves = getRyuuMoves(move, piece, game)
		result = checkEndPosInPossibleMoves(possibleMoves, move)
	case Checker:
	}

	if move.Promote != nil {
		//promote check

		switch piece.Type {
		case Pawn:
			err := checkPawnCheckerPromote(move, piece, game)
			if err != nil {
				return err
			}
		case Checker:
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
	}

	for i := 0; i < len(possibleMoves); i++ {
		possibleMove := possibleMoves[i]
		fmt.Println(possibleMove)
	}

	return result
}

func checkPositionInbounds(pos Vec2, game Game) bool {
	if pos.X < 0 || pos.X >= game.Board.Width {
		return false
	}
	if pos.Y < 0 || pos.Y >= game.Board.Height {
		return false
	}

	return true
}

func getEnemyTurnInt(game Game) int {
	if game.Turn == 0 {
		return 1
	} else {
		return 0
	}
}

func checkEndPosInPossibleMoves(possibleMoves []Vec2, move Move) error {
	for i := 0; i < len(possibleMoves); i++ {
		possibleMove := possibleMoves[i]

		if move.End == possibleMove {
			return nil
		}
	}

	return fmt.Errorf("No valid moves")
}

func getMoveDirection(game Game) int {
	direction := 1
	if game.Turn == 1 {
		direction = -1
	}
	return direction
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

func checkEmptySpace(move Move, game Game) error {
	space := game.Board.Board[move.End.Y][move.End.X]
	if space != nil {
		return fmt.Errorf("Cant drop on non empty space")
	}

	return nil
}

func checkPawnCheckerPromote(move Move, piece Piece, game Game) error {
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

func checkShogiPromote(move Move, piece Piece, game Game) error {
	if piece.Owner == 0 {

	} else {

	}

	return nil
}
