package engine

import (
	"fmt"
	"github.com/KainoaGardner/csc/internal/utils"
)

// if drop check drop
// if in check cant move unless not in check after

// ADD other parts of move
func MovePiece(move Move, game *Game) error {
	err := CheckValidMove(move, *game)
	if err != nil {
		return err
	}

	piece, err := getPiece(move, *game)
	if err != nil {
		return err
	}

	takePiece := game.Board.Board[move.End.Y][move.End.X]
	validCastle := takePiece != nil && piece.Type == King && takePiece.Type == Rook && takePiece.Owner == piece.Owner

	dir := getMoveDirection(*game)

	updateEndPosition(move, game, piece, takePiece, dir, validCastle)

	updateEnPassantTakePosition(move, game, dir)
	updateEnPassantPosition(piece, move, game, dir)

	offset := getMochigomaOffset(*game)
	updateRemoveStartPosition(move, game, offset, validCastle)

	err = updateMochigoma(takePiece, game, offset)
	if err != nil {
		return err
	}

	if checkCheckerNextJumps(move, *piece, *game) {
		game.CheckerJump = true
	} else {
		updateHalfMoveCount(piece, takePiece, game)
		updateMoveCount(game)
		game.Turn = getEnemyTurnInt(*game)
	}

	return nil
}

func updateHalfMoveCount(piece *Piece, takePiece *Piece, game *Game) {
	if checkHalfMoveReset(piece, takePiece) {
		game.HalfMoveCount = 0
	} else {
		game.HalfMoveCount++
	}
}

func updateMoveCount(game *Game) {
	if game.Turn == 1 {
		game.MoveCount++
	}
}

func updateEnPassantPosition(piece *Piece, move Move, game *Game, dir int) {
	//update enPassant position
	if piece.Type == Pawn && utils.AbsoluteValueInt(move.Start.Y-move.End.Y) == 2 {
		game.EnPassant = &Vec2{X: move.Start.X, Y: move.Start.Y - dir}
	} else {
		game.EnPassant = nil
	}
}

func updateMochigoma(takePiece *Piece, game *Game, offset int) error {
	if takePiece != nil && takePiece.Type >= Fu && takePiece.Type <= Ryuu {
		mochigoma, ok := shogiDropPieceToMochiPiece[takePiece.Type]
		if !ok {
			return fmt.Errorf("Error converting taken piece to mochigoma")
		}
		game.Mochigoma[mochigoma+offset]++
	}
	return nil
}

func updateRemoveStartPosition(move Move, game *Game, offset int, validCastle bool) {
	if move.Drop != nil {
		game.Mochigoma[*move.Drop+offset]--
	} else if !validCastle {
		game.Board.Board[move.Start.Y][move.Start.X] = nil
	}
}

func updateEnPassantTakePosition(move Move, game *Game, dir int) {
	if game.EnPassant != nil && move.End == *game.EnPassant {
		game.Board.Board[move.End.Y-dir][move.End.X] = nil
	}
}

func updateEndPosition(move Move, game *Game, piece *Piece, takePiece *Piece, dir int, validCastle bool) {
	if validCastle {
		dx := utils.AbsoluteValueInt(move.End.X-move.Start.X) - 1
		kingX := (dx/2 + 1) * dir
		rookX := (dx / 2) * dir

		game.Board.Board[move.End.Y][kingX] = piece
		game.Board.Board[move.End.Y][rookX] = takePiece
		takePiece.Moved = true
	} else {
		game.Board.Board[move.End.Y][move.End.X] = piece
	}
	piece.Moved = true
}

func UpdateMoveHistory(move string, game *Game) {
	game.Moves = append(game.Moves, move)
}

func checkHalfMoveReset(piece *Piece, takePiece *Piece) bool {
	if takePiece != nil {
		return true
	}

	if piece != nil && (piece.Type == Pawn || piece.Type == Fu) {
		return true
	}

	return false
}

func CheckValidMove(move Move, game Game) error {
	err := checkMoveInBounds(move, game)
	if err != nil {
		return err
	}

	piece, err := getPiece(move, game)
	if err != nil {
		return err
	}

	if move.Drop != nil { //drop
		err = checkValidDrop(move, *piece, game)
		if err != nil {
			return err
		}
	} else {
		err = checkValidPieceMoves(move, *piece, game) //normal
		if err != nil {
			return err
		}
	}

	if move.Promote != nil {
		err = checkValidPromote(move, *piece, game)
		if err != nil {
			return err
		}
	} else {
		err = checkMustPromote(move, *piece, game) //pawn checker last row must promote
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
		mochigoma := *move.Drop
		koma, ok := shogiMochiPieceToDropPiece[mochigoma]
		if !ok {
			return piece, fmt.Errorf("Could not fight correct piece from drop mochigoma")
		}

		dropPiece.Type = koma
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
