package api

import (
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

	r.Delete("/game", h.deleteAllGames)

	r.Get("/game/{gameID}/board", h.getBoard)
	r.Post("/game/{gameID}/move", h.postMovePiece)
	r.Post("/game/{gameID}/place", h.postPlacePiece)
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
	var gameConfig types.PostGame
	err := utils.ParseJSON(r, &gameConfig)
	if err != nil {
		utils.WriteError(w, http.StatusBadRequest, err)
		return
	}

	game, err := engine.SetupNewGame(gameConfig)
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
	// game, err := db.FindGame(h.client, h.dbConfig, gameID)
	// if err != nil {
	// 	utils.WriteError(w, http.StatusBadRequest, err)
	// 	return
	// }
	//
	// utils.WriteJSON(w, http.StatusOK, result)
	//
}

func (h *Handler) postPlacePiece(w http.ResponseWriter, r *http.Request) {
	var placeConfig types.PostPlace
	err := utils.ParseJSON(r, &placeConfig)
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

	place, err := engine.SetupPlace(placeConfig, "", *game)
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
