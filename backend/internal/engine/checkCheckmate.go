package engine

import (
	"fmt"
	"github.com/KainoaGardner/csc/internal/types"
)

func GetInCheck(game Game) bool {
	kings := getKingPos(game)
	for _, king := range kings {
		if checkUnderAttack(king, game) {
			return true
		}
	}

	return false
}

func getKingPos(game Game) []types.Vec2 {
	var result []types.Vec2

	for i := 0; i < game.Board.Height; i++ {
		for j := 0; j < game.Board.Width; j++ {
			space := game.Board.Board[i][j]

			if space != nil && space.Owner == game.Turn && (space.Type == King || space.Type == Ou) {
				result = append(result, types.Vec2{X: j, Y: i})
			}
		}
	}

	return result
}

func checkUnderAttack(pos types.Vec2, game Game) bool {
	attackSpace := map[types.Vec2]bool{}

	game.Turn = getEnemyTurnInt(game)
	for i := 0; i < game.Board.Height; i++ {
		for j := 0; j < game.Board.Width; j++ {
			space := game.Board.Board[i][j]
			if space != nil && space.Owner == game.Turn {
				pos := types.Vec2{X: j, Y: i}
				dir := getMoveDirection(game)
				possibleMoves := getPieceMoves(pos, *space, game, dir)
				for _, move := range possibleMoves {
					attackSpace[move] = true
				}
			}
		}
	}

	_, ok := attackSpace[pos]
	return ok
}

func GetInCheckmate(game Game) int {
	if !GetInCheck(game) {
		possibleMoves := getAllPossibleMovesCheckmate(game)

		fmt.Println(possibleMoves)
		if len(possibleMoves) > 0 {
			return 0
		} else {
			possibleDrops := getAllPossibleDrops(game)
			if len(possibleDrops) > 0 {
				return 0
			} else {
				return 2
			}
		}
	}

	possibleMoves := getAllPossibleMovesCheckmate(game)
	if len(possibleMoves) > 0 {
		return 0
	}

	possibleDrops := getAllPossibleDrops(game)
	for i := 0; i < len(possibleDrops); i++ {
		movePos := possibleDrops[i]
		move := Move{}
		move.End = movePos
		piece := Piece{}
		piece.Owner = game.Turn
		piece.Type = Fu
		gameCopy := copyGame(game)
		gameCopy.Board.Board[movePos.Y][movePos.X] = &piece
		if !GetInCheck(*gameCopy) {
			return 0
		}
	}

	return 1
}

func getValidPieceMovesForCheckmate(pos types.Vec2, piece Piece, game Game) []types.Vec2 {
	dir := getMoveDirection(game)
	possibleMoves := []types.Vec2{}
	if piece.Type == King {
		possibleMoves = getKingMoves(pos, piece, game)
	} else {
		possibleMoves = getPieceMoves(pos, piece, game, dir)
	}

	filterPossibleMoves(pos, &possibleMoves, game)

	return possibleMoves
}

func getAllPossibleMovesCheckmate(game Game) []types.Vec2 {
	var possibleMoves []types.Vec2
	//check normal moves
	for i := 0; i < game.Board.Height; i++ {
		for j := 0; j < game.Board.Width; j++ {
			space := game.Board.Board[i][j]
			if space != nil && space.Owner == game.Turn {
				possibleMoves = append(possibleMoves, getValidPieceMovesForCheckmate(types.Vec2{X: j, Y: i}, *space, game)...)
			}
		}
	}

	return possibleMoves
}

func getAllPossibleDrops(game Game) []types.Vec2 {
	var possibleDrops []types.Vec2

	for i := 0; i < game.Board.Height; i++ {
		for j := 0; j < game.Board.Width; j++ {
			for k := 0; k < 7; k++ {
				move := Move{}
				move.End.X = j
				move.End.Y = i
				move.Drop = &k
				piece := Piece{}
				piece.Owner = game.Turn
				piece.Type = Fu + k
				err := checkValidDrop(move, piece, game)
				if err == nil {
					possibleDrops = append(possibleDrops, move.End)
				}
			}
		}
	}

	return possibleDrops
}
