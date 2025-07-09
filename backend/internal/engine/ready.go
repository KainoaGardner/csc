package engine

import (
	"fmt"
	"github.com/KainoaGardner/csc/internal/types"
	"time"
)

func ReadyPlayer(ready bool, turn int, game *types.Game) error {
	err := checkGameState(types.PlaceState, game.State)
	if err != nil {
		return err
	}

	if !ready {
		if !game.Ready[turn] {
			return fmt.Errorf("Already not ready")
		}

		game.Ready[turn] = false

		return nil
	}

	err = checkHasKing(turn, *game)
	if err != nil {
		return err
	}

	game.Ready[turn] = true

	if checkBothReady(*game) {
		game.State = types.MoveState
		currTime := time.Now()
		game.LastMoveTime = currTime.UTC()
		err = checkGameOverAtStart(game)
		if err != nil {
			return err
		}
	}

	return nil
}

func checkGameOverAtStart(game *types.Game) error {
	checkMate := [2]bool{false, false}
	if GetInCheckmate(*game) {
		checkMate[game.Turn] = true
	}

	game.Turn = getEnemyTurnInt(*game)
	if GetInCheckmate(*game) {
		checkMate[game.Turn] = true
	}

	game.Turn = getEnemyTurnInt(*game)

	white := types.White
	black := types.Black
	tie := types.Tie

	if checkMate[0] && checkMate[1] {
		game.Winner = &tie
		game.Reason = "Double checkmate"
		game.State = types.OverState
	} else if checkMate[0] {
		game.Winner = &black
		game.Reason = "Checkmate"
		game.State = types.OverState
	} else if checkMate[1] {
		game.Winner = &white
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

func checkHasKing(turn int, game types.Game) error {
	for i := 0; i < game.Board.Height; i++ {
		for j := 0; j < game.Board.Width; j++ {
			space := game.Board.Board[i][j]
			if space != nil && space.Owner == turn && (space.Type == types.King || space.Type == types.Ou) {
				return nil
			}
		}
	}

	return fmt.Errorf("Need at least one king or ou to ready up")
}

func checkBothReady(game types.Game) bool {
	return game.Ready[0] && game.Ready[1]
}
