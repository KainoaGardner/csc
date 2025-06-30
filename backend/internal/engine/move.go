package engine

import (
	"fmt"
	"github.com/KainoaGardner/csc/internal/types"
	"github.com/KainoaGardner/csc/internal/utils"
)

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
		game.EnPassant = &types.Vec2{X: move.Start.X, Y: move.Start.Y - dir}
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
