package engine

import (
	"fmt"
	"github.com/KainoaGardner/csc/internal/types"
	"github.com/KainoaGardner/csc/internal/utils"
)

func TestGame(game types.Game) int {
	game.Board.Board[0][0] = &types.Piece{Type: types.Ryuu, Owner: 1, Moved: true}

	return game.Board.Board[0][0].Type
}

// if in check cant move unless not in check after

// ADD other parts of move
func MovePiece(move types.Move, game *types.Game) error {
	err := checkGameOver(*game)
	if err != nil {
		return err
	}

	err = CheckValidMove(move, *game)
	if err != nil {
		return err
	}

	piece, err := getPiece(move, *game)
	if err != nil {
		return err
	}

	takePiece := game.Board.Board[move.End.Y][move.End.X]
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

		checkmate := GetInCheckmate(*game)
		if checkmate == 1 {
			moveTurn := getEnemyTurnInt(*game)
			game.Winner = &moveTurn
		} else if checkmate == 2 {
			tie := types.Tie
			game.Winner = &tie
		}
	}

	return nil
}

func CheckValidMove(move types.Move, game types.Game) error {
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
