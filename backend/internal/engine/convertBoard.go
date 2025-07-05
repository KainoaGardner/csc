package engine

import (
	"fmt"
	"github.com/KainoaGardner/csc/internal/types"
	"strconv"
	"strings"
)

func ConvertBoardToString(game types.Game) (string, error) {
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
	result += moveCountString + " "

	timeString := getTimeString(game)
	result += timeString

	return result, nil
}

func convertPiecePositionToString(game types.Game) (string, error) {
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

				pieceString, ok := types.FenPieceToString[piece.Type]
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

func convertMochigomaToString(game types.Game) string {
	result := ""
	for i := 0; i < len(game.Mochigoma); i++ {
		result += strconv.Itoa(game.Mochigoma[i]) + "/"
	}

	return result[:len(result)-1]
}

func getTurnString(game types.Game) string {
	if game.Turn == 1 {
		return "b"
	} else {

		return "w"
	}
}

func getEnPassantString(game types.Game) (string, error) {
	if game.EnPassant == nil {
		return "-", nil
	}
	result, err := convertPositionToString(*game.EnPassant, game)
	if err != nil {
		return result, err
	}

	return result, nil
}

func getCheckerJumpString(game types.Game) (string, error) {
	if game.CheckerJump == nil {
		return "-", nil
	}
	result, err := convertPositionToString(*game.CheckerJump, game)
	if err != nil {
		return result, err
	}

	return result, nil
}

func getHalfMoveCountString(game types.Game) string {
	return strconv.Itoa(game.HalfMoveCount)
}

func getMoveCountString(game types.Game) string {
	return strconv.Itoa(game.MoveCount)
}

func getTimeString(game types.Game) string {
	result := ""

	result += strconv.Itoa(game.Time[0]) + "/"
	result += strconv.Itoa(game.Time[1])

	return result
}
