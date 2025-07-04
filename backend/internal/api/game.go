package api

import (
	"fmt"
	"github.com/KainoaGardner/csc/internal/engine"
	"github.com/KainoaGardner/csc/internal/types"
	"github.com/KainoaGardner/csc/internal/utils"
	"github.com/go-chi/chi/v5"
	"net/http"
)

func (h *Handler) registerGameRoutes(r chi.Router) {
	r.Get("/test/move", h.moveTest)
	r.Get("/game/board", h.getBoard)
	r.Post("/game/move", h.movePiece)
	r.Get("/game/check", h.getCheck)
	r.Get("/game/checkmate", h.getCheckmate)
	r.Post("/game/validMove", h.validMove)

	//use ID for each game access
	// r.Post("/game/id/create", h.getBoard)
	// r.Post("/game/id/move", h.getBoard)
}

func (h *Handler) moveTest(w http.ResponseWriter, r *http.Request) {
	engine.RunTests()
	w.Write([]byte("Test Finished"))
}

func (h *Handler) getCheck(w http.ResponseWriter, r *http.Request) {
	game := types.Game{}
	game.Board.Width = 8
	game.Board.Height = 8

	game.Board.Board = make([][]*types.Piece, game.Board.Height)
	for i := range game.Board.Board {
		game.Board.Board[i] = make([]*types.Piece, game.Board.Width)
	}

	game.Turn = 0
	game.Board.Board[7][4] = &types.Piece{Type: types.King, Owner: 0}
	game.Board.Board[3][4] = &types.Piece{Type: types.Rook, Owner: 1}
	game.Board.Board[0][4] = &types.Piece{Type: types.King, Owner: 1}

	result := engine.GetInCheck(game)
	if result {
		w.Write([]byte("Check"))
	} else {
		utils.WriteError(w, http.StatusBadRequest, fmt.Errorf("Not Check"))
	}
}

func (h *Handler) getCheckmate(w http.ResponseWriter, r *http.Request) {
	game := types.Game{}
	game.Board.Width = 8
	game.Board.Height = 8

	game.Board.Board = make([][]*types.Piece, game.Board.Height)
	for i := range game.Board.Board {
		game.Board.Board[i] = make([]*types.Piece, game.Board.Width)
	}

	game.Turn = 0
	game.Board.Board[7][4] = &types.Piece{Type: types.King, Owner: 0}
	game.Board.Board[6][4] = &types.Piece{Type: types.To, Owner: 1}
	game.Board.Board[0][4] = &types.Piece{Type: types.Hi, Owner: 1}

	result := engine.GetInCheckmate(game)

	if result == 1 {
		w.Write([]byte("Checkmate"))
	} else if result == 2 {
		utils.WriteError(w, http.StatusBadRequest, fmt.Errorf("Draw"))
	} else {
		utils.WriteError(w, http.StatusBadRequest, fmt.Errorf("Not Checkmate"))
	}
}

func (h *Handler) getBoard(w http.ResponseWriter, r *http.Request) {
	game := types.Game{}
	game.Board.Width = 8
	game.Board.Height = 8

	game.Board.Board = make([][]*types.Piece, game.Board.Height)
	for i := range game.Board.Board {
		game.Board.Board[i] = make([]*types.Piece, game.Board.Width)
	}
	game.Turn = 0
	game.Board.Board[6][0] = &types.Piece{Type: types.Pawn, Owner: 0}
	game.Board.Board[7][6] = &types.Piece{Type: types.Knight, Owner: 0}
	game.Board.Board[4][2] = &types.Piece{Type: types.Bishop, Owner: 0}
	game.Board.Board[6][4] = &types.Piece{Type: types.Rook, Owner: 0}
	game.Board.Board[4][6] = &types.Piece{Type: types.Queen, Owner: 0}
	game.Board.Board[3][3] = &types.Piece{Type: types.King, Owner: 1}
	game.Board.Board[3][4] = &types.Piece{Type: types.Fu, Owner: 0}
	game.Board.Board[3][7] = &types.Piece{Type: types.Kyou, Owner: 0}
	game.Board.Board[5][2] = &types.Piece{Type: types.Kei, Owner: 0}
	game.Board.Board[2][4] = &types.Piece{Type: types.Pawn, Owner: 1}
	game.Board.Board[3][1] = &types.Piece{Type: types.Pawn, Owner: 1}
	game.Board.Board[0][7] = &types.Piece{Type: types.Pawn, Owner: 1}

	result, err := engine.ConvertBoardToString(game)
	if err != nil {
		utils.WriteError(w, http.StatusBadRequest, err)
		return
	}
	utils.WriteJSON(w, http.StatusOK, result)
}

func (h *Handler) movePiece(w http.ResponseWriter, r *http.Request) {

}

func (h *Handler) validMove(w http.ResponseWriter, r *http.Request) {
	game := types.Game{}
	game.Board.Width = 8
	game.Board.Height = 8

	game.Board.Board = make([][]*types.Piece, game.Board.Height)
	for i := range game.Board.Board {
		game.Board.Board[i] = make([]*types.Piece, game.Board.Width)
	}

	game.Turn = 0

	game.Board.Board[6][0] = &types.Piece{Type: types.Pawn, Owner: 0}
	// game.Board.Board[7][6] = &types.Piece{Type: types.Knight, Owner: 0}
	// game.Board.Board[4][2] = &types.Piece{Type: types.Bishop, Owner: 0}
	// game.Board.Board[6][4] = &types.Piece{Type: types.Rook, Owner: 0}
	// game.Board.Board[4][6] = &types.Piece{Type: types.Queen, Owner: 0}
	// game.Board.Board[3][3] = &types.Piece{Type: types.King, Owner: 1}
	// game.Board.Board[3][4] = &types.Piece{Type: types.Fu, Owner: 0}
	// game.Board.Board[3][7] = &types.Piece{Type: types.Kyou, Owner: 0}
	// game.Board.Board[5][2] = &types.Piece{Type: types.Kei, Owner: 0}
	// game.Board.Board[2][4] = &types.Piece{Type: types.Pawn, Owner: 1}
	// game.Board.Board[3][1] = &types.Piece{Type: types.Pawn, Owner: 1}
	// game.Board.Board[0][7] = &types.Piece{Type: types.Pawn, Owner: 1}

	var postMove types.PostMoveString
	err := utils.ParseJSON(r, &postMove)
	if err != nil {
		utils.WriteError(w, http.StatusBadRequest, err)
		return
	}

	move, err := engine.ConvertStringToMove(postMove.Move, game)
	if err != nil {
		utils.WriteError(w, http.StatusBadRequest, err)
		return
	}

	err = engine.CheckValidMove(move, game)
	if err != nil {
		utils.WriteError(w, http.StatusBadRequest, err)
		return
	}

	utils.WriteJSON(w, http.StatusOK, "Valid move")
}
