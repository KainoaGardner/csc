package api

import (
	"github.com/KainoaGardner/csc/internal/auth"
	"github.com/KainoaGardner/csc/internal/db"
	"github.com/KainoaGardner/csc/internal/engine"
	"github.com/KainoaGardner/csc/internal/types"
	"github.com/KainoaGardner/csc/internal/utils"
	"github.com/go-chi/chi/v5"

	"fmt"
	"net/http"
)

func (h *Handler) registerGameRoutes(r chi.Router) {
	r.Get("/game", h.getAllGames)
	r.Post("/game", h.postCreateGame)
	r.Post("/game/{gameID}/join", h.postJoinGame)

	r.Delete("/game", h.deleteAllGames)

	r.Get("/game/{gameID}", h.getBoard)
	r.Post("/game/{gameID}/move", h.postMovePiece)
	r.Post("/game/{gameID}/place", h.postPlacePiece)
	r.Delete("/game/{gameID}/place", h.deletePlacePiece)

	r.Post("/game/{gameID}/state", h.postState)

	r.Post("/game/{gameID}/ready", h.postReady)
	r.Post("/game/{gameID}/draw", h.postDraw)
}

// admin
func (h *Handler) getAllGames(w http.ResponseWriter, r *http.Request) {
	statusCode, err := auth.CheckAdminRequest(h.client, h.dbConfig, h.jwt.AccessKey, r)
	if err != nil {
		utils.WriteError(w, statusCode, err)
		return
	}

	games, err := db.ListAllGames(h.client, h.dbConfig)
	if err != nil {
		utils.WriteError(w, http.StatusBadRequest, err)
		return
	}

	result := []types.GetGameResponse{}
	for _, game := range games {
		gameResponse := types.GetGameResponse{}
		gameResponse.ID = game.ID
		gameResponse.WhiteID = game.WhiteID
		gameResponse.BlackID = game.BlackID
		gameResponse.Turn = game.Turn
		gameResponse.MoveCount = game.MoveCount
		gameResponse.HalfMoveCount = game.HalfMoveCount
		gameResponse.Winner = game.Winner
		gameResponse.Reason = game.Reason
		gameResponse.State = game.State
		gameResponse.Time = game.Time
		gameResponse.LastMoveTime = game.LastMoveTime
		gameResponse.Money = game.Money
		gameResponse.Ready = game.Ready
		gameResponse.Draw = game.Draw

		result = append(result, gameResponse)
	}

	utils.WriteResponse(w, http.StatusOK, fmt.Sprintf("%d games found", len(result)), result)
}

// admin
func (h *Handler) deleteAllGames(w http.ResponseWriter, r *http.Request) {
	statusCode, err := auth.CheckAdminRequest(h.client, h.dbConfig, h.jwt.AccessKey, r)
	if err != nil {
		utils.WriteError(w, statusCode, err)
		return
	}

	amount, err := db.DeleteAllGames(h.client, h.dbConfig)
	if err != nil {
		utils.WriteError(w, http.StatusBadRequest, err)
		return
	}

	data := map[string]interface{}{"count": amount}

	utils.WriteResponse(w, http.StatusOK, "Games deleted", data)
}

// auth
func (h *Handler) postCreateGame(w http.ResponseWriter, r *http.Request) {
	claims, statusCode, err := auth.CheckValidAuth(h.client, h.dbConfig, h.jwt.AccessKey, r)
	if err != nil {
		utils.WriteError(w, statusCode, err)
		return
	}

	var postGame types.PostGame
	err = utils.ParseJSON(r, &postGame)
	if err != nil {
		utils.WriteError(w, http.StatusBadRequest, err)
		return
	}

	game, err := engine.SetupNewGame(postGame, claims.UserID)
	if err != nil {
		utils.WriteError(w, http.StatusBadRequest, err)
		return
	}

	gameID, err := db.CreateGame(h.client, h.dbConfig, game)
	if err != nil {
		utils.WriteError(w, http.StatusBadRequest, err)
		return
	}

	data := types.PostGameResponse{
		ID:        gameID,
		WhiteID:   game.WhiteID,
		BlackID:   game.BlackID,
		Color:     "w",
		Width:     game.Board.Width,
		Height:    game.Board.Height,
		Money:     game.Money,
		PlaceLine: game.Board.PlaceLine,
		State:     game.State,
	}

	utils.WriteResponse(w, http.StatusOK, fmt.Sprintf("Game created"), data)
}

// auth
func (h *Handler) postJoinGame(w http.ResponseWriter, r *http.Request) {
	claims, statusCode, err := auth.CheckValidAuth(h.client, h.dbConfig, h.jwt.AccessKey, r)
	if err != nil {
		utils.WriteError(w, statusCode, err)
		return
	}

	gameID := chi.URLParam(r, "gameID")
	game, err := db.FindGame(h.client, h.dbConfig, gameID)
	if err != nil {
		utils.WriteError(w, http.StatusBadRequest, err)
		return
	}

	err = engine.UpdateStartGame(game, claims.UserID)
	if err != nil {
		utils.WriteError(w, http.StatusBadRequest, err)
		return
	}

	err = db.GameMoveUpdate(h.client, h.dbConfig, gameID, *game)
	if err != nil {
		utils.WriteError(w, http.StatusBadRequest, err)
		return
	}

	data := types.PostGameResponse{
		ID:        gameID,
		WhiteID:   game.WhiteID,
		BlackID:   game.BlackID,
		Color:     "w",
		Width:     game.Board.Width,
		Height:    game.Board.Height,
		Money:     game.Money,
		PlaceLine: game.Board.PlaceLine,
		State:     game.State,
	}

	utils.WriteResponse(w, http.StatusOK, fmt.Sprintf("Joined"), data)
}

// auth either player
func (h *Handler) getBoard(w http.ResponseWriter, r *http.Request) {
	claims, statusCode, err := auth.CheckValidAuth(h.client, h.dbConfig, h.jwt.AccessKey, r)
	if err != nil {
		utils.WriteError(w, statusCode, err)
		return
	}

	gameID := chi.URLParam(r, "gameID")
	game, err := db.FindGame(h.client, h.dbConfig, gameID)
	if err != nil {
		utils.WriteError(w, http.StatusBadRequest, err)
		return
	}

	_, err = engine.GetTurnFromID(*game, claims.UserID)
	if err != nil {
		utils.WriteError(w, http.StatusBadRequest, err)
		return
	}

	result, err := engine.ConvertBoardToString(*game)
	if err != nil {
		utils.WriteError(w, http.StatusBadRequest, err)
		return
	}

	data := map[string]interface{}{
		"_id": game.ID,
		"fen": result,
	}

	utils.WriteResponse(w, http.StatusOK, "Board", data)
}

// auth either player
func (h *Handler) postMovePiece(w http.ResponseWriter, r *http.Request) {
	claims, statusCode, err := auth.CheckValidAuth(h.client, h.dbConfig, h.jwt.AccessKey, r)
	if err != nil {
		utils.WriteError(w, statusCode, err)
		return
	}

	var postMove types.PostMove
	err = utils.ParseJSON(r, &postMove)
	if err != nil {
		utils.WriteError(w, http.StatusBadRequest, err)
		return
	}

	gameID := chi.URLParam(r, "gameID")
	game, err := db.FindGame(h.client, h.dbConfig, gameID)
	if err != nil {
		utils.WriteError(w, http.StatusBadRequest, err)
		return
	}

	turn, err := engine.GetTurnFromID(*game, claims.UserID)
	if err != nil {
		utils.WriteError(w, http.StatusBadRequest, err)
		return
	}

	err = engine.CheckTurn(turn, game.Turn)
	if err != nil {
		utils.WriteError(w, http.StatusBadRequest, err)
		return
	}

	move, err := engine.ConvertStringToMove(postMove.Move, *game)
	if err != nil {
		utils.WriteError(w, http.StatusBadRequest, err)
		return
	}

	err = engine.MovePiece(move, game)
	if err != nil {
		utils.WriteError(w, http.StatusBadRequest, err)
		return
	}

	err = db.GameMoveUpdate(h.client, h.dbConfig, gameID, *game)
	if err != nil {
		utils.WriteError(w, http.StatusBadRequest, err)
		return
	}

	fen, err := engine.ConvertBoardToString(*game)
	if err != nil {
		utils.WriteError(w, http.StatusBadRequest, err)
		return
	}

	err = db.GameLogUpdate(h.client, h.dbConfig, gameID, postMove.Move, fen)
	if err != nil {
		utils.WriteError(w, http.StatusBadRequest, err)
		return
	}

	if game.State == types.OverState {
		gameLog, err := db.FindGameLogFromGameID(h.client, h.dbConfig, gameID)
		if err != nil {
			utils.WriteError(w, http.StatusBadRequest, err)
			return
		}
		engine.SetupFinalGameLog(*game, gameLog)
		err = db.GameLogFinalUpdate(h.client, h.dbConfig, gameID, *gameLog)
		if err != nil {
			utils.WriteError(w, http.StatusBadRequest, err)
			return
		}

		_, err = db.DeleteGame(h.client, h.dbConfig, gameID)
		if err != nil {
			utils.WriteError(w, http.StatusBadRequest, err)
			return
		}

		data := map[string]interface{}{
			"_id":           game.ID,
			"whiteID":       game.WhiteID,
			"blackID":       game.BlackID,
			"moveCount":     game.MoveCount,
			"halfMoveCount": game.HalfMoveCount,
			"winner":        game.Winner,
			"reason":        game.Reason,
			"state":         game.State,
			"lastMoveTime":  game.LastMoveTime,
		}

		utils.WriteResponse(w, http.StatusOK, "Game Over", data)
	} else {
		data := types.PostMoveResponse{
			ID:   game.ID,
			FEN:  fen,
			Move: postMove.Move,
		}
		utils.WriteResponse(w, http.StatusOK, "Piece moved", data)
	}

}

// auth either player
func (h *Handler) postPlacePiece(w http.ResponseWriter, r *http.Request) {
	claims, statusCode, err := auth.CheckValidAuth(h.client, h.dbConfig, h.jwt.AccessKey, r)
	if err != nil {
		utils.WriteError(w, statusCode, err)
		return
	}

	var postPlace types.PostPlace
	err = utils.ParseJSON(r, &postPlace)
	if err != nil {
		utils.WriteError(w, http.StatusBadRequest, err)
		return
	}

	gameID := chi.URLParam(r, "gameID")
	game, err := db.FindGame(h.client, h.dbConfig, gameID)
	if err != nil {
		utils.WriteError(w, http.StatusBadRequest, err)
		return
	}

	turn, err := engine.GetTurnFromID(*game, claims.UserID)
	if err != nil {
		utils.WriteError(w, http.StatusBadRequest, err)
		return
	}

	place, err := engine.SetupPlace(postPlace, turn, *game)
	if err != nil {
		utils.WriteError(w, http.StatusBadRequest, err)
		return
	}

	err = engine.PlacePiece(place, game)
	if err != nil {
		utils.WriteError(w, http.StatusBadRequest, err)
		return
	}

	err = db.GamePlaceUpdate(h.client, h.dbConfig, gameID, place, *game)
	if err != nil {
		utils.WriteError(w, http.StatusBadRequest, err)
		return
	}

	fen, err := engine.ConvertBoardToString(*game)
	if err != nil {
		utils.WriteError(w, http.StatusBadRequest, err)
		return
	}

	data := types.PlaceResponse{
		ID:       game.ID,
		FEN:      fen,
		Position: postPlace.Position,
		Type:     postPlace.Type,
		Cost:     place.Cost,
		Money:    game.Money,
	}

	utils.WriteResponse(w, http.StatusOK, "Piece placed", data)
}

// auth either player
func (h *Handler) deletePlacePiece(w http.ResponseWriter, r *http.Request) {
	claims, statusCode, err := auth.CheckValidAuth(h.client, h.dbConfig, h.jwt.AccessKey, r)
	if err != nil {
		utils.WriteError(w, statusCode, err)
		return
	}

	var deletePlace types.DeletePlace
	err = utils.ParseJSON(r, &deletePlace)
	if err != nil {
		utils.WriteError(w, http.StatusBadRequest, err)
		return
	}

	gameID := chi.URLParam(r, "gameID")
	game, err := db.FindGame(h.client, h.dbConfig, gameID)
	if err != nil {
		utils.WriteError(w, http.StatusBadRequest, err)
		return
	}

	turn, err := engine.GetTurnFromID(*game, claims.UserID)
	if err != nil {
		utils.WriteError(w, http.StatusBadRequest, err)
		return
	}

	place, err := engine.SetupDeletePlace(deletePlace, turn, *game)
	if err != nil {
		utils.WriteError(w, http.StatusBadRequest, err)
		return
	}

	err = engine.PlacePieceDelete(&place, game)
	if err != nil {
		utils.WriteError(w, http.StatusBadRequest, err)
		return
	}

	err = db.GamePlaceUpdate(h.client, h.dbConfig, gameID, place, *game)
	if err != nil {
		utils.WriteError(w, http.StatusBadRequest, err)
		return
	}

	fen, err := engine.ConvertBoardToString(*game)
	if err != nil {
		utils.WriteError(w, http.StatusBadRequest, err)
		return
	}

	data := types.PlaceResponse{
		ID:       game.ID,
		FEN:      fen,
		Position: deletePlace.Position,
		Type:     place.Type,
		Cost:     -place.Cost,
		Money:    game.Money,
	}

	utils.WriteResponse(w, http.StatusOK, "Piece deleted", data)
}

// admin
func (h *Handler) postState(w http.ResponseWriter, r *http.Request) {
	statusCode, err := auth.CheckAdminRequest(h.client, h.dbConfig, h.jwt.AccessKey, r)
	if err != nil {
		utils.WriteError(w, statusCode, err)
		return
	}

	var postState types.PostState
	err = utils.ParseJSON(r, &postState)
	if err != nil {
		utils.WriteError(w, http.StatusBadRequest, err)
		return
	}

	gameID := chi.URLParam(r, "gameID")
	game, err := db.FindGame(h.client, h.dbConfig, gameID)
	if err != nil {
		utils.WriteError(w, http.StatusBadRequest, err)
		return
	}

	game.State = postState.State

	err = db.GameStateUpdate(h.client, h.dbConfig, gameID, *game)
	if err != nil {
		utils.WriteError(w, http.StatusBadRequest, err)
		return
	}

	data := map[string]interface{}{
		"state": game.State,
	}

	utils.WriteResponse(w, http.StatusOK, "State changed", data)
}

// auth either player
func (h *Handler) postReady(w http.ResponseWriter, r *http.Request) {
	claims, statusCode, err := auth.CheckValidAuth(h.client, h.dbConfig, h.jwt.AccessKey, r)
	if err != nil {
		utils.WriteError(w, statusCode, err)
		return
	}

	var postReady types.PostReady
	err = utils.ParseJSON(r, &postReady)
	if err != nil {
		utils.WriteError(w, http.StatusBadRequest, err)
		return
	}

	gameID := chi.URLParam(r, "gameID")
	game, err := db.FindGame(h.client, h.dbConfig, gameID)
	if err != nil {
		utils.WriteError(w, http.StatusBadRequest, err)
		return
	}

	turn, err := engine.GetTurnFromID(*game, claims.UserID)
	if err != nil {
		utils.WriteError(w, http.StatusBadRequest, err)
		return
	}

	err = engine.ReadyPlayer(postReady.Ready, turn, game)
	if err != nil {
		utils.WriteError(w, http.StatusBadRequest, err)
		return
	}

	err = db.GameReadyUpdate(h.client, h.dbConfig, gameID, *game)
	if err != nil {
		utils.WriteError(w, http.StatusBadRequest, err)
		return
	}

	if game.State == types.MoveState {
		gameLog := engine.SetupGameLog(*game)
		gameLogID, err := db.CreateGameLog(h.client, h.dbConfig, gameLog)
		if err != nil {
			utils.WriteError(w, http.StatusBadRequest, err)
			return
		}
		data := map[string]interface{}{
			"_id":       game.ID,
			"state":     game.State,
			"ready":     game.Ready,
			"gameLogID": gameLogID,
		}
		utils.WriteResponse(w, http.StatusOK, "Game Start", data)
	} else if game.State == types.OverState {
		data := map[string]interface{}{
			"_id":           game.ID,
			"whiteID":       game.WhiteID,
			"blackID":       game.BlackID,
			"moveCount":     game.MoveCount,
			"halfMoveCount": game.HalfMoveCount,
			"winner":        game.Winner,
			"reason":        game.Reason,
			"state":         game.State,
			"lastMoveTime":  game.LastMoveTime,
		}

		utils.WriteResponse(w, http.StatusOK, "Game Over", data)
	} else {
		data := map[string]interface{}{
			"_id":   game.ID,
			"state": game.State,
			"ready": game.Ready,
		}

		utils.WriteResponse(w, http.StatusOK, "Ready", data)
	}

}

// auth either player
func (h *Handler) postDraw(w http.ResponseWriter, r *http.Request) {
	claims, statusCode, err := auth.CheckValidAuth(h.client, h.dbConfig, h.jwt.AccessKey, r)
	if err != nil {
		utils.WriteError(w, statusCode, err)
		return
	}

	var postDraw types.PostDrawRequest
	err = utils.ParseJSON(r, &postDraw)
	if err != nil {
		utils.WriteError(w, http.StatusBadRequest, err)
		return
	}

	gameID := chi.URLParam(r, "gameID")
	game, err := db.FindGame(h.client, h.dbConfig, gameID)
	if err != nil {
		utils.WriteError(w, http.StatusBadRequest, err)
		return
	}

	turn, err := engine.GetTurnFromID(*game, claims.UserID)
	if err != nil {
		utils.WriteError(w, http.StatusBadRequest, err)
		return
	}

	err = engine.DrawRequest(postDraw.Draw, turn, game)
	if err != nil {
		utils.WriteError(w, http.StatusBadRequest, err)
		return
	}

	err = db.GameDrawUpdate(h.client, h.dbConfig, gameID, *game)
	if err != nil {
		utils.WriteError(w, http.StatusBadRequest, err)
		return
	}

	if game.State == types.OverState {
		gameLog, err := db.FindGameLogFromGameID(h.client, h.dbConfig, gameID)
		if err != nil {
			utils.WriteError(w, http.StatusBadRequest, err)
			return
		}
		engine.SetupFinalGameLog(*game, gameLog)
		err = db.GameLogFinalUpdate(h.client, h.dbConfig, gameID, *gameLog)
		if err != nil {
			utils.WriteError(w, http.StatusBadRequest, err)
			return
		}

		data := map[string]interface{}{
			"_id":           game.ID,
			"whiteID":       game.WhiteID,
			"blackID":       game.BlackID,
			"moveCount":     game.MoveCount,
			"halfMoveCount": game.HalfMoveCount,
			"winner":        game.Winner,
			"reason":        game.Reason,
			"state":         game.State,
			"draw":          game.Draw,
			"lastMoveTime":  game.LastMoveTime,
		}

		utils.WriteResponse(w, http.StatusOK, "Game Over", data)
	} else {
		data := map[string]interface{}{
			"_id":  game.ID,
			"draw": game.Draw,
		}

		utils.WriteResponse(w, http.StatusOK, "Draw", data)
	}
}
