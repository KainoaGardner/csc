package main

import (
	"fmt"

	"github.com/KainoaGardner/csc/internal/engine"
)

func main() {
	game := engine.Game{}
	game.Board.Width = 8
	game.Board.Height = 8

	// game.Board.Board = make([][]int, game.Board.Height)
	// for i := range game.Board.Board {
	// 	game.Board.Board[i] = make([]int, game.Board.Width)
	// }

	game.Board.Board = [][]int{
		{-4, -2, -3, -5, -6, -3, -2, -4},
		{-1, -1, -1, -1, -1, -1, -1, -1},
		{0, 0, 0, 0, 0, 0, 0, 0},
		{0, 0, 0, 0, 0, 0, 0, 0},
		{0, 0, 0, 0, 0, 0, 0, 0},
		{0, 0, 0, 0, 0, 0, 0, 0},
		{1, 1, 1, 1, 1, 1, 1, 1},
		{4, 2, 3, 5, 6, 3, 2, 4},
	}

	game.Turn = 0

	moveStr := "e2,e4"

	move, err := engine.ConvertStringToMove(moveStr, game)
	if err != nil {
		fmt.Print(err)
		return
	}

	fmt.Println(moveStr)
	fmt.Println(move)
}
