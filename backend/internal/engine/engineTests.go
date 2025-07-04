package engine

import (
	"fmt"
	"github.com/KainoaGardner/csc/internal/types"
)

func RunTests() {
	game := types.Game{}
	game.Board.Width = 8
	game.Board.Height = 8

	game.Board.Board = make([][]*types.Piece, game.Board.Height)
	for i := range game.Board.Board {
		game.Board.Board[i] = make([]*types.Piece, game.Board.Width)
	}

	game.Turn = 0

	moves := []string{
		// Chess
		"b1,c3",
		"g8,f6",
		"e5,f6",
		"f6,g7",
		"h1,f1",
		"f3,e5",
		"c8,g4",
		"g4,e6",
		"e6,d7",

		// Chess promotions
		"b7,b8Q",
		"c7,c8N",
		"d7,d8R",
		"f7,f8B",

		// Chess en passant-like (can be treated specially)
		"e5,d6", // capture behind if tracked correctly
		"d5,c6",

		// Shogi moves
		"c3,b2+",
		"g6,f5",
		"e2,d1+",
		"f4,e3+",
		"d6,e5",
		"e5,d4+",
		"b7,a6",

		// Shogi drops
		"P*,f6",
		"L*,d3",
		"N*,e2",
		"S*,c5",
		"G*,b4",
		"B*,a8",
		"R*,g3",

		// Extra shogi promoted drops
		"N*,c8",
		"S*,e1",
		"P*,g2",

		// Shogi aggressive drops
		"B*,d4",
		"R*,f5",

		// Checkers
		"b6,c5",
		"c5,d6",
		"d6,e7",
		"e7,f6",
		"f6,g7",
		"g7,h8",

		// Checkers multi-capture pattern
		"a3,c5",
		"c5,e7",
		"e7,g5",

		// More Shogi promotions
		"d4,e3+",
		"e3,f2+",
		"f2,g1+",

		// More Chess edge cases
		"a2,a4",
		"h2,h4",
		"g1,h3",
		"b8,a6",

		// Shogi drop chain
		"P*,e5",
		"P*,e6",
		"P*,e7",
		"P*,e8",

		//checkers
		"c3,d4",
		"f6,e5",
		"b2,d4",
		"g5,e3",
		"a2,c4",
		"h5,f3+",
		"e3,c5",
		"h8,g7",
		"c5,e3",

		// Final few for randomness
		"g2,f3",
		"f3,e4",
		"e4,d5+",
		"P*,f7",
		"N*,e8",
	}

	for i := 0; i < len(moves); i++ {
		err := convertStringToMoveTest(moves[i], game)
		if err != nil {
			fmt.Println(err)
			fmt.Println("FAILED", moves[i])
			return
		}

		fmt.Println("Passed", moves[i])
	}
}

func convertStringToMoveTest(input string, game types.Game) error {
	movesResult, err := ConvertStringToMove(input, game)
	if err != nil {
		return err
	}

	stringResult, err := ConvertMoveToString(movesResult, game)
	if err != nil {
		return err
	}

	result := input == stringResult
	if !result {
		return fmt.Errorf("Results do not match. %s %s", input, stringResult)
	}

	return nil
}
