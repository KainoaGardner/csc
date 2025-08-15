package engine

import (
	"fmt"
	"github.com/KainoaGardner/csc/internal/types"
)

func SetupNewGame(gameConfig types.PostGame, userID string) (*types.Game, error) {
	err := checkSetupConfig(gameConfig)
	if err != nil {
		return nil, err
	}

	game := types.Game{}

	// game.WhiteID = userID
	game.Time = gameConfig.StartTime
	game.Time[0] *= 1000
	game.Time[1] *= 1000
	game.Money = gameConfig.Money

	game.PositionHistory = map[string]int{}

	game.Board.Width = gameConfig.Width
	game.Board.Height = gameConfig.Height
	game.Board.Board = make([][]*types.Piece, game.Board.Height)
	for i := range game.Board.Board {
		game.Board.Board[i] = make([]*types.Piece, game.Board.Width)
	}

	game.Board.PlaceLine = gameConfig.PlaceLine

	game.State = 0
	game.Public = gameConfig.Public

	return &game, nil
}

func checkSetupConfig(gameConfig types.PostGame) error {
	if gameConfig.StartTime[0] < 0 || gameConfig.StartTime[1] < 0 {
		return fmt.Errorf("Cannot have negative start time")
	}

	if gameConfig.StartTime[0] > 100000 || gameConfig.StartTime[1] > 100000 {
		return fmt.Errorf("Starttime limit 100000")
	}

	if gameConfig.Money[0] < 50 || gameConfig.Money[1] < 50 {
		return fmt.Errorf("Need at least 50 money")
	}

	if gameConfig.Money[0] > 100000000 || gameConfig.Money[1] > 100000000 {
		return fmt.Errorf("Money limit 100000 seconds")
	}

	if gameConfig.Width <= 0 || gameConfig.Height <= 0 {
		return fmt.Errorf("Cannot have negative or 0 Width or Height")
	}

	if gameConfig.Width > 20 || gameConfig.Height > 20 {
		return fmt.Errorf("Cannot have width or height bigger than 20")
	}

	if gameConfig.PlaceLine >= gameConfig.Height || gameConfig.PlaceLine <= 0 {
		return fmt.Errorf("PlaceLine must be within the board")
	}

	return nil
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

func UpdateStartGame(game *types.Game, userID string) error {
	err := checkGameState(types.ConnectState, game.State)
	if err != nil {
		return err
	}

	if userID == game.WhiteID {
		return fmt.Errorf("Cant join your own game")
	}

	game.BlackID = userID
	game.State = types.PlaceState
	return nil
}

func SetupJoinGame(game *types.Game, userID string) error {
	err := checkGameState(types.ConnectState, game.State)
	if err != nil {
		return err
	}

	if game.WhiteID == userID || game.BlackID == userID {
		return fmt.Errorf("Already joined game")
	}

	if game.WhiteID == "" {
		game.WhiteID = userID

	} else if game.BlackID == "" {
		game.BlackID = userID
	} else {
		return fmt.Errorf("Game full")
	}

	if game.WhiteID != "" && game.BlackID != "" {
		game.State = types.PlaceState
	}

	return nil
}

func GetTurnFromID(game types.Game, userID string) (int, error) {
	if game.WhiteID == userID {
		return types.White, nil
	}
	if game.BlackID == userID {
		return types.Black, nil
	}

	return -1, fmt.Errorf("Player not in game")
}

func SetupResignGame(game *types.Game, turn int) {
	var otherTurn int
	if turn == 0 {
		otherTurn = 1
	} else {
		otherTurn = 0
	}

	game.Winner = &otherTurn
	game.Reason = "Resignation"
	game.State = types.OverState
}
