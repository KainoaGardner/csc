package engine

import (
	"github.com/KainoaGardner/csc/internal/types"
)

func GetDraw(game *types.Game) (bool, error) {
	if checkStalemate(*game) {
		return true, nil
	}

	boardString, err := ConvertBoardToStringPositionKey(*game)
	if err != nil {
		return false, err
	}
	updateMoveHistory(boardString, game)
	if checkThreefoldRepetition(boardString, *game) {
		return true, nil
	}

	if checkFiftyMoveRule(*game) {
		return true, nil
	}

	if checkInsufficientMaterial(*game) {
		return true, nil
	}

	return false, nil
}

func checkStalemate(game types.Game) bool {
	if GetInCheck(game) {
		return false
	}

	possibleMoves := getAllPossibleMovesCheckmate(game)
	if len(possibleMoves) > 0 {
		return false
	}

	possibleDrops := getAllPossibleDrops(game)
	if len(possibleDrops) > 0 {
		return false
	}

	return true
}

func updateMoveHistory(boardString string, game *types.Game) {
	_, ok := game.PositionHistory[boardString]
	if !ok {
		game.PositionHistory[boardString] = 0
	}
	game.PositionHistory[boardString]++
}

func checkThreefoldRepetition(boardString string, game types.Game) bool {
	count, ok := game.PositionHistory[boardString]
	if ok && count >= 3 {
		return true
	}
	return false
}

func checkFiftyMoveRule(game types.Game) bool {
	if game.HalfMoveCount >= 100 {
		return true
	}
	return false
}

func checkInsufficientMaterial(game types.Game) bool {
	for i := 0; i < types.MochigomaSize; i++ {
		if game.Mochigoma[i] > 0 {
			return false
		}
	}

	pieceCounts := [2]map[int]int{
		make(map[int]int),
		make(map[int]int),
	}

	for i := 0; i < game.Board.Height; i++ {
		for j := 0; j < game.Board.Width; j++ {
			piece := game.Board.Board[i][j]
			if piece != nil {
				if piece.Type >= types.Fu && piece.Type <= types.Ryuu { //if any shogi pieces no stalemate
					return false
				}

				pieceCounts[piece.Owner][piece.Type]++
			}
		}
	}

	if checkAbleToMate(pieceCounts[0]) {
		return false
	}
	if checkAbleToMate(pieceCounts[1]) {
		return false
	}

	return true
}

func checkAbleToMate(pieceCounts map[int]int) bool {
	if pieceCounts[types.Queen] > 0 || pieceCounts[types.Rook] > 0 || pieceCounts[types.Pawn] > 0 {
		return true
	}

	if pieceCounts[types.Knight] >= 2 || pieceCounts[types.Bishop] >= 2 {
		return true
	}

	if pieceCounts[types.Knight] >= 1 && pieceCounts[types.Bishop] >= 1 {
		return true
	}

	checkerTotal := pieceCounts[types.Checker] + pieceCounts[types.CheckerKing]
	if checkerTotal >= 2 {
		return true
	}

	if checkerTotal >= 1 && (pieceCounts[types.Knight] >= 1 || pieceCounts[types.Bishop] >= 1) {
		return true
	}

	return false
}

func DrawRequest(draw bool, turn int, game *types.Game) error {
	err := checkGameState(types.MoveState, game.State)
	if err != nil {
		return err
	}

	if !draw {
		game.Draw[0] = false
		game.Draw[1] = false
		return nil
	}

	game.Draw[turn] = true
	if game.Draw[0] && game.Draw[1] {
		tie := types.Tie
		game.Winner = &tie
		game.Reason = "Draw"
		game.State = types.OverState
	}

	return nil
}
