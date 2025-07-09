package engine

import (
	"fmt"
	"github.com/KainoaGardner/csc/internal/types"
	"strconv"
)

func SetupNewGame(gameConfig types.PostGame) (*types.Game, error) {
	err := checkSetupConfig(gameConfig)
	if err != nil {
		return nil, err
	}

	game := types.Game{}
	game.Time = gameConfig.StartTime
	game.Money = gameConfig.Money

	game.PositionHistory = map[string]int{}

	game.Board.Width = gameConfig.Width
	game.Board.Height = gameConfig.Height
	game.Board.Board = make([][]*types.Piece, game.Board.Height)
	for i := range game.Board.Board {
		game.Board.Board[i] = make([]*types.Piece, game.Board.Width)
	}

	game.Board.PlaceLine = gameConfig.PlaceLine

	//REMOVE LATER
	game.State = 0

	return &game, nil
}

func checkSetupConfig(gameConfig types.PostGame) error {
	if gameConfig.StartTime[0] < 0 || gameConfig.StartTime[1] < 0 {
		return fmt.Errorf("Cannot have negative start time")
	}

	if gameConfig.Money[0] < 0 || gameConfig.Money[1] < 0 {
		return fmt.Errorf("Cannot have negative money")
	}

	if gameConfig.Width <= 0 || gameConfig.Height <= 0 {
		return fmt.Errorf("Cannot have negative or 0 Width or Height")
	}

	return nil
}

func GetTurnFromID(id string, game types.Game) (int, error) {
	var result int

	// if id == game.WhiteID {
	// 	result = 0
	// } else if id == game.BlackID {
	// 	result = 1
	// } else {
	// 	return 0, fmt.Errorf("Invalid ID")
	// }

	result, _ = strconv.Atoi(id) //TEMP

	return result, nil
}

func checkGameState(mode int, state int) error {
	if state != mode {
		return fmt.Errorf("Incorrect state")
	}

	return nil
}

func CheckTurn(turn int, gameTurn int) error {
	if turn != gameTurn {
		return fmt.Errorf("Not your turn")
	}

	return nil
}

func UpdateStartGame(game *types.Game) error {
	err := checkGameState(types.ConnectState, game.State)
	if err != nil {
		return err
	}
	game.State = types.PlaceState
	return nil

	//set black ID
}
