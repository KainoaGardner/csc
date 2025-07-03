package engine

import (
	"fmt"
	"strconv"
	"strings"
)

// Alternate FEN 2 Char
func ConvertBoardToString(game Game) (string, error) {
	result := ""

	piecePositionString, err := convertPiecePositionToString(game)
	if err != nil {
		return result, err
	}
	result += piecePositionString + " "

	mochigomaString := convertMochigomaToString(game)
	result += mochigomaString + " "

	turnString := getTurnString(game)
	result += turnString + " "

	enPassantString, err := getEnPassantString(game)
	if err != nil {
		return result, err
	}
	result += enPassantString + " "

	checkerJumpString, err := getCheckerJumpString(game)
	if err != nil {
		return result, err
	}
	result += checkerJumpString + " "

	halfMoveCountString := getHalfMoveCountString(game)
	result += halfMoveCountString + " "

	moveCountString := getMoveCountString(game)
	result += moveCountString

	return result, nil
}

func convertPiecePositionToString(game Game) (string, error) {
	result := ""

	for i := 0; i < game.Board.Height; i++ {
		emptyCount := 0
		for j := 0; j < game.Board.Width; j++ {
			piece := game.Board.Board[i][j]
			if piece == nil {
				emptyCount++
			} else {
				if emptyCount != 0 {
					result += strconv.Itoa(emptyCount)
					emptyCount = 0
				}

				pieceString, ok := fenPieceToString[piece.Type]
				if !ok {
					return result, fmt.Errorf("Could not convert piece to fen string")
				}

				if piece.Owner == 1 {
					pieceString = strings.ToLower(pieceString)
				}

				result += pieceString

				if !piece.Moved {
					result += "*"
				} else {
					result += "-"
				}
			}
		}

		if emptyCount != 0 {
			result += strconv.Itoa(emptyCount)
		}

		if i != game.Board.Height-1 {
			result += "/"
		}

	}

	return result, nil
}

func convertMochigomaToString(game Game) string {
	result := ""
	for i := 0; i < len(game.Mochigoma); i++ {
		result += strconv.Itoa(game.Mochigoma[i]) + "/"
	}

	return result[:len(result)-1]
}

func getTurnString(game Game) string {
	if game.Turn == 1 {
		return "b"
	} else {

		return "w"
	}
}

func getEnPassantString(game Game) (string, error) {
	if game.EnPassant == nil {
		return "-", nil
	}
	result, err := convertPositionToString(*game.EnPassant, game)
	if err != nil {
		return result, err
	}

	return result, nil
}

func getCheckerJumpString(game Game) (string, error) {
	if game.CheckerJump == nil {
		return "-", nil
	}
	result, err := convertPositionToString(*game.CheckerJump, game)
	if err != nil {
		return result, err
	}

	return result, nil
}

func getHalfMoveCountString(game Game) string {
	return strconv.Itoa(game.HalfMoveCount)
}

func getMoveCountString(game Game) string {
	return strconv.Itoa(game.MoveCount)
}
