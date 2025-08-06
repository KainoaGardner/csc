package engine

import (
	"fmt"
	"github.com/KainoaGardner/csc/internal/types"
	"github.com/KainoaGardner/csc/internal/utils"
)

func MovePiece(move types.Move, game *types.Game) error {
	err := checkGameState(types.MoveState, game.State)
	if err != nil {
		return err
	}

	err = checkGameOver(*game)
	if err != nil {
		return err
	}

	err = checkValidMove(move, *game)
	if err != nil {
		return err
	}

	piece, err := getPiece(move, *game)
	if err != nil {
		return err
	}

	updateMoveTime(game)
	if checkTimeLoss(*game) {
		moveTurn := getEnemyTurnInt(*game)
		game.Winner = &moveTurn
		game.Reason = "Time"
		game.State = types.OverState
		return nil
	}

	takePiece := game.Board.Board[move.End.Y][move.End.X]
	// takePiece := getTakePiece()
	validCastle := takePiece != nil && piece.Type == types.King && takePiece.Type == types.Rook && takePiece.Owner == piece.Owner

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

	if checkCheckerNextJumps(move.Start, move.End, *piece, *game) {
		game.CheckerJump = &move.End
	} else {
		updateHalfMoveCount(piece, takePiece, game)
		updateMoveCount(game)
		game.Turn = getEnemyTurnInt(*game)
		err = checkCheckmateOrDraw(game)
		if err != nil {
			return err
		}
	}

	return nil
}

func checkCheckmateOrDraw(game *types.Game) error {
	if GetInCheckmate(*game) {
		moveTurn := getEnemyTurnInt(*game)
		game.Winner = &moveTurn
		game.Reason = "Checkmate"
		game.State = types.OverState
	} else {
		result, err := GetDraw(game)
		if err != nil {
			return err
		}
		if result {
			tie := types.Tie
			game.Winner = &tie
			game.Reason = "Draw"
			game.State = types.OverState
		}
	}
	return nil
}

func checkValidMove(move types.Move, game types.Game) error {
	err := checkMoveInBounds(move, game)
	if err != nil {
		return err
	}

	piece, err := getPiece(move, game)
	if err != nil {
		return err
	}

	err = checkCheckerJumpMove(move, game)
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

func checkMoveInBounds(move types.Move, game types.Game) error {
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

func getPiece(move types.Move, game types.Game) (*types.Piece, error) {
	var piece *types.Piece
	if move.Drop != nil {
		var dropPiece types.Piece
		mochigoma := *move.Drop
		koma, ok := types.ShogiMochiPieceToDropPiece[mochigoma]
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

func checkPositionInbounds(pos types.Vec2, game types.Game) bool {
	if pos.X < 0 || pos.X >= game.Board.Width {
		return false
	}
	if pos.Y < 0 || pos.Y >= game.Board.Height {
		return false
	}

	return true
}

func getEnemyTurnInt(game types.Game) int {
	if game.Turn == 0 {
		return 1
	} else {
		return 0
	}
}

func checkCheckerJumpMove(move types.Move, game types.Game) error {
	if game.CheckerJump == nil {
		return nil
	}

	if !utils.CheckVec2Equal(move.Start, *game.CheckerJump) {
		return fmt.Errorf("Most complete all checker jumps")
	}

	return nil
}

func checkGameOver(game types.Game) error {
	if game.Winner != nil {
		return fmt.Errorf("Game is over")
	}

	return nil
}

func updateHalfMoveCount(piece *types.Piece, takePiece *types.Piece, game *types.Game) {
	if checkHalfMoveReset(piece, takePiece) {
		game.HalfMoveCount = 0
	} else {
		game.HalfMoveCount++
	}
}

func updateMoveCount(game *types.Game) {
	if game.Turn == 1 {
		game.MoveCount++
	}
}

func updateEnPassantPosition(piece *types.Piece, move types.Move, game *types.Game, dir int) {
	//update enPassant position
	if piece.Type == types.Pawn && utils.AbsoluteValueInt(move.Start.Y-move.End.Y) == 2 {
		game.EnPassant = &types.Vec2{X: move.Start.X, Y: move.Start.Y - dir}
	} else {
		game.EnPassant = nil
	}
}

func updateMochigoma(takePiece *types.Piece, game *types.Game, offset int) error {
	if takePiece != nil && takePiece.Type >= types.Fu && takePiece.Type <= types.Ryuu {
		mochigoma, ok := types.ShogiDropPieceToMochiPiece[takePiece.Type]
		if !ok {
			return fmt.Errorf("Error converting taken piece to mochigoma")
		}
		game.Mochigoma[mochigoma+offset]++
	}
	return nil
}

func updateRemoveStartPosition(move types.Move, game *types.Game, offset int, validCastle bool) {
	if move.Drop != nil {
		game.Mochigoma[*move.Drop+offset]--
	} else if !validCastle {
		game.Board.Board[move.Start.Y][move.Start.X] = nil
	}
}

func updateEnPassantTakePosition(move types.Move, game *types.Game, dir int) {
	if game.EnPassant != nil && move.End == *game.EnPassant {
		game.Board.Board[move.End.Y+dir][move.End.X] = nil
	}
}

func updateEndPosition(move types.Move, game *types.Game, piece *types.Piece, takePiece *types.Piece, dir int, validCastle bool) {
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

func checkHalfMoveReset(piece *types.Piece, takePiece *types.Piece) bool {
	if takePiece != nil {
		return true
	}

	if piece != nil && (piece.Type == types.Pawn || piece.Type == types.Fu) {
		return true
	}

	return false
}
