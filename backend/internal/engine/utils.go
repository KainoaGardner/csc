package engine

import (
	"github.com/KainoaGardner/csc/internal/types"
)

func copyGame(game types.Game) *types.Game {
	gameCopy := game

	boardCopy := types.Board{}
	boardCopy.Width = game.Board.Width
	boardCopy.Height = game.Board.Height

	boardCopy.Board = make([][]*types.Piece, game.Board.Height)
	for i := range game.Board.Board {
		boardCopy.Board[i] = make([]*types.Piece, game.Board.Width)
	}

	for i := range game.Board.Height {
		for j := range game.Board.Width {
			piece := game.Board.Board[i][j]
			if piece != nil {
				pieceCopy := types.Piece{}
				pieceCopy.Type = piece.Type
				pieceCopy.Owner = piece.Owner
				pieceCopy.Moved = piece.Moved
				boardCopy.Board[i][j] = &pieceCopy
			} else {
				boardCopy.Board[i][j] = nil
			}

		}
	}

	gameCopy.Board = boardCopy

	if game.EnPassant != nil {
		enPassantPos := types.Vec2{X: game.EnPassant.X, Y: game.EnPassant.Y}
		gameCopy.EnPassant = &enPassantPos
	} else {
		gameCopy.EnPassant = nil
	}

	if game.CheckerJump != nil {
		checkerJump := types.Vec2{X: game.CheckerJump.X, Y: game.CheckerJump.Y}
		gameCopy.CheckerJump = &checkerJump
	} else {
		gameCopy.CheckerJump = nil
	}

	return &gameCopy
}
