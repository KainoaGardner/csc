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

	r.Get("/game/{gameID}/board", h.getBoard)
	r.Post("/game/{gameID}/move", h.postMovePiece)
	r.Post("/game/{gameID}/place", h.postPlacePiece)
	r.Delete("/game/{gameID}/place", h.deletePlacePiece)

	r.Post("/game/{gameID}/state", h.postState)
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

	utils.WriteJSON(w, http.StatusOK, ids)
}

func (h *Handler) deleteAllGames(w http.ResponseWriter, r *http.Request) {
	amount, err := db.DeleteAllGames(h.client, h.dbConfig)
	if err != nil {
		utils.WriteError(w, http.StatusBadRequest, err)
		return
	}

	utils.WriteJSON(w, http.StatusOK, fmt.Sprintf("%d games deleted", amount))
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

	utils.WriteJSON(w, http.StatusOK, gameID)
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

	fmt.Println(result)
	utils.WriteJSON(w, http.StatusOK, result)
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

	err = db.MoveUpdate(h.client, h.dbConfig, gameID, *game)
	if err != nil {
		utils.WriteError(w, http.StatusBadRequest, err)
		return
	}

	fen, err := engine.ConvertBoardToString(*game)
	if err != nil {
		utils.WriteError(w, http.StatusBadRequest, err)
		return
	}

	data := types.PostMoveResponse{
		FEN:  fen,
		Move: postMove.Move,
	}

	utils.WriteResponse(w, http.StatusOK, "Piece Moved", data)

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

	err = db.PlaceUpdate(h.client, h.dbConfig, gameID, place, *game)
	if err != nil {
		utils.WriteError(w, http.StatusBadRequest, err)
		return
	}

	utils.WriteJSON(w, http.StatusOK, "Piece Placed")
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

	err = engine.PlacePieceDelete(place, game)
	if err != nil {
		utils.WriteError(w, http.StatusBadRequest, err)
		return
	}

	err = db.PlaceUpdate(h.client, h.dbConfig, gameID, place, *game)
	if err != nil {
		utils.WriteError(w, http.StatusBadRequest, err)
		return
	}

	utils.WriteJSON(w, http.StatusOK, "Piece Deleted")
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

	err = db.StateUpdate(h.client, h.dbConfig, gameID, *game)
	if err != nil {
		utils.WriteError(w, http.StatusBadRequest, err)
		return
	}

	utils.WriteJSON(w, http.StatusOK, game.State)
}
