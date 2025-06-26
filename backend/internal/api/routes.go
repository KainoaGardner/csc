package api

import (
	"github.com/KainoaGardner/csc/internal/engine"
	"github.com/KainoaGardner/csc/internal/utils"
	"github.com/go-chi/chi/v5"
	"net/http"
)

type Handler struct {
}

func NewHandler() *Handler {
	return &Handler{}
}

func (h *Handler) RegisterRoutes(r chi.Router) {
	r.Get("/test/move", h.moveTest)
	r.Get("/game/board", h.getBoard)
	r.Post("/game/move", h.movePiece)
	r.Post("/game/validMove", h.validMove)

	//use ID for each game access
	// r.Post("/game/id/create", h.getBoard)
	// r.Post("/game/id/move", h.getBoard)
}

func (h *Handler) moveTest(w http.ResponseWriter, r *http.Request) {
	engine.RunTests()
	w.Write([]byte("Test Finished"))
}

func (h *Handler) getBoard(w http.ResponseWriter, r *http.Request) {
	game := engine.Game{}
	game.Board.Width = 8
	game.Board.Height = 8

	game.Board.Board = make([][]*engine.Piece, game.Board.Height)
	for i := range game.Board.Board {
		game.Board.Board[i] = make([]*engine.Piece, game.Board.Width)
	}
	game.Turn = 0
	game.Board.Board[6][0] = &engine.Piece{Type: engine.Pawn, Owner: 0}
	game.Board.Board[7][6] = &engine.Piece{Type: engine.Knight, Owner: 0}
	game.Board.Board[4][2] = &engine.Piece{Type: engine.Bishop, Owner: 0}
	game.Board.Board[6][4] = &engine.Piece{Type: engine.Rook, Owner: 0}
	game.Board.Board[4][6] = &engine.Piece{Type: engine.Queen, Owner: 0}
	game.Board.Board[3][3] = &engine.Piece{Type: engine.King, Owner: 1}
	game.Board.Board[3][4] = &engine.Piece{Type: engine.Fu, Owner: 0}
	game.Board.Board[3][7] = &engine.Piece{Type: engine.Kyou, Owner: 0}
	game.Board.Board[5][2] = &engine.Piece{Type: engine.Kei, Owner: 0}
	game.Board.Board[2][4] = &engine.Piece{Type: engine.Pawn, Owner: 1}
	game.Board.Board[3][1] = &engine.Piece{Type: engine.Pawn, Owner: 1}
	game.Board.Board[0][7] = &engine.Piece{Type: engine.Pawn, Owner: 1}

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
	game := engine.Game{}
	game.Board.Width = 8
	game.Board.Height = 8

	game.Board.Board = make([][]*engine.Piece, game.Board.Height)
	for i := range game.Board.Board {
		game.Board.Board[i] = make([]*engine.Piece, game.Board.Width)
	}

	game.Turn = 0

	game.Board.Board[6][0] = &engine.Piece{Type: engine.Pawn, Owner: 0}
	// game.Board.Board[7][6] = &engine.Piece{Type: engine.Knight, Owner: 0}
	// game.Board.Board[4][2] = &engine.Piece{Type: engine.Bishop, Owner: 0}
	// game.Board.Board[6][4] = &engine.Piece{Type: engine.Rook, Owner: 0}
	// game.Board.Board[4][6] = &engine.Piece{Type: engine.Queen, Owner: 0}
	// game.Board.Board[3][3] = &engine.Piece{Type: engine.King, Owner: 1}
	// game.Board.Board[3][4] = &engine.Piece{Type: engine.Fu, Owner: 0}
	// game.Board.Board[3][7] = &engine.Piece{Type: engine.Kyou, Owner: 0}
	// game.Board.Board[5][2] = &engine.Piece{Type: engine.Kei, Owner: 0}
	// game.Board.Board[2][4] = &engine.Piece{Type: engine.Pawn, Owner: 1}
	// game.Board.Board[3][1] = &engine.Piece{Type: engine.Pawn, Owner: 1}
	// game.Board.Board[0][7] = &engine.Piece{Type: engine.Pawn, Owner: 1}

	var postMove PostMoveString
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
