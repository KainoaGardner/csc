package main

import (
	"fmt"
	"github.com/KainoaGardner/csc/internal/engine"
)

func main() {
	game := engine.Game{}
	game.Board.Width = 8
	game.Board.Height = 8

	game.Board.Board = make([][]engine.Piece, game.Board.Height)
	for i := range game.Board.Board {
		game.Board.Board[i] = make([]engine.Piece, game.Board.Width)
	}

	game.Turn = 0

	game.Board.Board[3][4] = engine.Piece{Type: engine.Pawn, Owner: 1}
	game.Board.Board[4][3] = engine.Piece{Type: engine.Pawn, Owner: 0}

	for i := 0; i < game.Board.Height; i++ {
		for j := 0; j < game.Board.Width; j++ {
			sign := 1
			if game.Board.Board[j][i].Owner == 1 {
				sign = -1
			}

			fmt.Print(game.Board.Board[i][j].Type * sign)
		}
		fmt.Println()
	}

	moveStr := "d4,e5"
	move, err := engine.ConvertStringToMove(moveStr, game)
	if err != nil {
		fmt.Println(err)
		return
	}

	fmt.Println(move)
	// engine.SetupMove(&move, game)
	fmt.Println(move)

}
