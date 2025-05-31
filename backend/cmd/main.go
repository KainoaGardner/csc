package main

import (
	"fmt"
	"github.com/KainoaGardner/csc/internal/engine"
	"strconv"
)

func main() {
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

	for i := 0; i < game.Board.Height; i++ {
		for j := 0; j < game.Board.Width; j++ {
			piece := game.Board.Board[i][j]
			if piece != nil {

				sign := 1
				if game.Board.Board[i][j].Owner == 1 {
					sign = -1
				}

				pieceStr := strconv.Itoa(game.Board.Board[i][j].Type * sign)
				if len(pieceStr) < 2 {
					pieceStr = " " + pieceStr
				}
				fmt.Print(pieceStr)

			} else {
				fmt.Print("__")
			}
		}
		fmt.Println()
	}

	// movesString := "a2,a3"
	movesString := "e2,a3"
	moves, err := engine.ConvertStringToMoves(movesString, game)
	if err != nil {
		fmt.Println(err)
		return
	}

	engine.CheckValidMove(&moves, game)

}
