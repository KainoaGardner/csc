package engine

import (
	"github.com/KainoaGardner/csc/internal/types"
	"time"
)

func checkTimeLoss(game types.Game) bool {
	if game.Time[game.Turn] < 0 {
		return true
	}

	return false
}

func updateMoveTime(game *types.Game) {
	currTime := time.Now().UTC()
	dt := currTime.Sub(game.LastMoveTime).Milliseconds()
	buffer := (time.Duration(2) * time.Second).Milliseconds()

	if dt <= buffer {
		return
	}
	game.Time[game.Turn] -= dt
	game.LastMoveTime = time.Now().UTC()
}
