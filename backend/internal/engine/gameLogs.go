package engine

import (
	"github.com/KainoaGardner/csc/internal/types"
	"time"
)

func SetupGameLog(game types.Game) *types.GameLog {
	var result types.GameLog

	result.ID = game.ID
	result.WhiteID = game.WhiteID
	result.BlackID = game.BlackID

	result.Moves = []string{}
	result.BoardStates = []string{}

	localTime := time.Now()
	utcTime := localTime.UTC()
	result.Date = utcTime

	return &result
}

func SetupFinalGameLog(game types.Game, gameLog *types.GameLog) {
	gameLog.GameID = game.ID.Hex()
	gameLog.MoveCount = game.MoveCount
	gameLog.Winner = game.Winner
	gameLog.Reason = game.Reason
}
