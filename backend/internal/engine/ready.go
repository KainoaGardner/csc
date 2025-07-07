package engine

import (
	"fmt"
	"github.com/KainoaGardner/csc/internal/types"
	"time"
)

func ReadyPlayer(turn int, game *types.Game) error {
	err := checkGameState(types.PlaceState, game.State)
	if err != nil {
		return err
	}

	err = checkHasKing(turn, *game)
	if err != nil {
		return err
	}

	game.Ready[turn] = true

	if checkBothReady(*game) {
		game.State = 2
		game.LastMoveTime = time.Now()
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
