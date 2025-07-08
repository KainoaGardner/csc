package api

import (
	"github.com/KainoaGardner/csc/internal/db"
	"github.com/KainoaGardner/csc/internal/engine"
	"github.com/KainoaGardner/csc/internal/types"
	"github.com/KainoaGardner/csc/internal/utils"
	"github.com/go-chi/chi/v5"

	"fmt"
	"net/http"

	"strconv"
)

func (h *Handler) registerGameRoutes(r chi.Router) {
	r.Get("/game", h.getAllGames)
	r.Post("/game", h.postCreateGame)

	r.Delete("/game", h.deleteAllGames)

	r.Get("/game/{gameID}", h.getBoard)
	r.Post("/game/{gameID}/move", h.postMovePiece)
	r.Post("/game/{gameID}/place", h.postPlacePiece)
	r.Delete("/game/{gameID}/place", h.deletePlacePiece)

	r.Post("/game/{gameID}/state", h.postState)

	r.Post("/game/{gameID}/ready", h.postReady)
	r.Post("/game/{gameID}/unready", h.postUnready)
}

func (h *Handler) getAllGames(w http.ResponseWriter, r *http.Request) {
	games, err := db.ListAllGames(h.client, h.dbConfig)
	if err != nil {
		utils.WriteError(w, http.StatusBadRequest, err)
		return
	}

	ids := []string{}
	for _, game := range games {
		idString := game.ID.Hex()
		ids = append(ids, idString)
	}

	utils.WriteResponse(w, http.StatusOK, fmt.Sprintf("%d games found", len(ids)), ids)
}

func (h *Handler) deleteAllGames(w http.ResponseWriter, r *http.Request) {
	amount, err := db.DeleteAllGames(h.client, h.dbConfig)
	if err != nil {
		utils.WriteError(w, http.StatusBadRequest, err)
		return
	}

	data := map[string]interface{}{"count": amount}

	utils.WriteResponse(w, http.StatusOK, "Games deleted", data)
}

func (h *Handler) postCreateGame(w http.ResponseWriter, r *http.Request) {
	var postGame types.PostGame
	err := utils.ParseJSON(r, &postGame)
	if err != nil {
		utils.WriteError(w, http.StatusBadRequest, err)
		return
	}

	game, err := engine.SetupNewGame(postGame)
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
	}

	utils.WriteResponse(w, http.StatusOK, fmt.Sprintf("Game created"), data)
}

func (h *Handler) getBoard(w http.ResponseWriter, r *http.Request) {
	gameID := chi.URLParam(r, "gameID")
	game, err := db.FindGame(h.client, h.dbConfig, gameID)
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
		"fen": result,
	}

	utils.WriteResponse(w, http.StatusOK, "Board", data)
}

func (h *Handler) postMovePiece(w http.ResponseWriter, r *http.Request) {
	var postMove types.PostMove
	err := utils.ParseJSON(r, &postMove)
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

	err = engine.CheckTurn(postMove.Turn, game.Turn)
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
	}

	data := types.PostMoveResponse{
		FEN:  fen,
		Move: postMove.Move,
	}
	utils.WriteResponse(w, http.StatusOK, "Piece moved", data)
}

func (h *Handler) postPlacePiece(w http.ResponseWriter, r *http.Request) {
	var postPlace types.PostPlace
	err := utils.ParseJSON(r, &postPlace)
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

	place, err := engine.SetupPlace(postPlace, strconv.Itoa(postPlace.Turn), *game)
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
		FEN:      fen,
		Position: postPlace.Position,
		Type:     postPlace.Type,
		Cost:     place.Cost,
		Money:    game.Money,
	}

	utils.WriteResponse(w, http.StatusOK, "Piece placed", data)
}

func (h *Handler) deletePlacePiece(w http.ResponseWriter, r *http.Request) {
	var deletePlace types.DeletePlace
	err := utils.ParseJSON(r, &deletePlace)
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

	place, err := engine.SetupDeletePlace(deletePlace, strconv.Itoa(deletePlace.Turn), *game)
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
		FEN:      fen,
		Position: deletePlace.Position,
		Type:     place.Type,
		Cost:     -place.Cost,
		Money:    game.Money,
	}

	utils.WriteResponse(w, http.StatusOK, "Piece deleted", data)
}

func (h *Handler) postState(w http.ResponseWriter, r *http.Request) {
	var postState types.PostState
	err := utils.ParseJSON(r, &postState)
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

func (h *Handler) postReady(w http.ResponseWriter, r *http.Request) {
	var postReady types.PostReady
	err := utils.ParseJSON(r, &postReady)
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

	err = engine.ReadyPlayer(postReady.Turn, game)
	if err != nil {
		utils.WriteError(w, http.StatusBadRequest, err)
		return
	}

	err = db.GameReadyUpdate(h.client, h.dbConfig, gameID, *game)
	if err != nil {
		utils.WriteError(w, http.StatusBadRequest, err)
		return
	}

	data := map[string]interface{}{
		"state": game.State,
		"ready": game.Ready,
	}

	if game.State == types.MoveState {
		gameLog := engine.SetupGameLog(*game)
		gameLogID, err := db.CreateGameLog(h.client, h.dbConfig, gameLog)
		if err != nil {
			utils.WriteError(w, http.StatusBadRequest, err)
			return
		}
		data = map[string]interface{}{
			"state":     game.State,
			"ready":     game.Ready,
			"gameLogID": gameLogID,
		}

	}

	utils.WriteResponse(w, http.StatusOK, "Ready", data)
}

func (h *Handler) postUnready(w http.ResponseWriter, r *http.Request) {
	var postReady types.PostReady
	err := utils.ParseJSON(r, &postReady)
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

	err = engine.UnreadyPlayer(postReady.Turn, game)
	if err != nil {
		utils.WriteError(w, http.StatusBadRequest, err)
		return
	}

	err = db.GameReadyUpdate(h.client, h.dbConfig, gameID, *game)
	if err != nil {
		utils.WriteError(w, http.StatusBadRequest, err)
		return
	}

	data := map[string]interface{}{
		"state": game.State,
		"ready": game.Ready,
	}

	utils.WriteResponse(w, http.StatusOK, "Unready", data)
}
