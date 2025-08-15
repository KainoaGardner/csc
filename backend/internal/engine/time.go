package engine

import (
	"time"

	"fmt"
	"github.com/KainoaGardner/csc/internal/config"
	"github.com/KainoaGardner/csc/internal/db"
	"github.com/KainoaGardner/csc/internal/types"
	"go.mongodb.org/mongo-driver/mongo"
	"log"
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
	game.Time[game.Turn] -= dt - buffer
	game.LastMoveTime = time.Now().UTC()
}

func StartGlobalTimeCheck(
	interval time.Duration,
	client *mongo.Client,
	config config.Config,
	gameOver func(*types.Game, string, string, *mongo.Client, config.Config) bool,
) {
	ticker := time.NewTicker(interval)

	go func() {
		for range ticker.C {

			games, err := db.ListAllGames(client, config.DB)
			if err != nil {
				log.Println(err)
				continue
			}

			currTime := time.Now().UTC()
			for _, game := range games {
				if game.State != types.MoveState {
					continue
				}

				dt := currTime.Sub(game.LastMoveTime).Milliseconds()
				buffer := (time.Duration(2) * time.Second).Milliseconds()

				if dt <= buffer {
					continue
				}

				remainingTime := game.Time[game.Turn] - dt + buffer
				if remainingTime < 0 {
					playerID := ""
					if game.Turn == types.White {
						playerID = game.WhiteID
					} else {
						playerID = game.BlackID
					}

					moveTurn := getEnemyTurnInt(*game)
					game.Winner = &moveTurn
					game.Reason = "Time"
					game.State = types.OverState

					game.Time[game.Turn] = 0
					gameOver(game, game.ID.Hex(), playerID, client, config)
					fmt.Println("TIme up")
				}
			}
		}
	}()

}
