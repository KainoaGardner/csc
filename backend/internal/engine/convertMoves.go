package engine

import (
	// "fmt"
	"github.com/KainoaGardner/csc/internal/utils"
	"strconv"
	"strings"
)

// func ConvertStringToMove(move string, game Game) (Move, error) {
// 	result := Move{}
// 	var err error
//
// 	commaIndex := utils.GetIndexFirstChar(move, ",")
// 	if commaIndex == -1 {
// 		return result, fmt.Errorf("Invalid Format. No ,")
// 	}
//
// 	result.Drop, result.StartPiece = checkDropPiece(move, game)
// 	if !result.Drop {
// 		result.Start, err = getStartPosition(move, commaIndex, game.Board.Height)
// 		if err != nil {
// 			return result, fmt.Errorf("Invalid Start Position")
// 		}
//
// 	}
//
// 	result.Promote, result.EndPiece = checkPromote(move, game)
//
// 	result.End, err = getEndPosition(move, commaIndex, game.Board.Height)
//
// 	if err != nil {
// 		return result, fmt.Errorf("Invalid End Position")
// 	}
//
// 	return result, nil
// }

func ConvertStringToMoves(move string, game Game) ([]Move, error) {
	var result []Move

	moveStrings := strings.Split(move, ",")
	for i := 0; i < len(moveStrings)-1; i++ {
		moveStartString := moveStrings[i]
		moveEndString := moveStrings[i+1]
		var move Move

		drop := checkDropPiece(moveStartString)
		if drop == 0 {
			start, err := convertStringToPosition(moveStartString, game.Board.Height)
			if err != nil {
				return result, err
			}
			move.Start = start
		} else {
			move.Drop = drop
		}

		promote := checkPromotePiece(moveEndString)
		move.Promote = promote

		end, err := convertStringToPosition(moveEndString, game.Board.Height)
		if err != nil {
			return result, err
		}
		move.End = end

		result = append(result, move)
	}

	// result := Move{}
	// var err error

	// commaIndex := utils.GetIndexFirstChar(move, ",")
	// if commaIndex == -1 {
	// 	return result, fmt.Errorf("Invalid Format. No ,")
	// }

	// result.Drop, result.StartPiece = checkDropPiece(move, game)
	// if !result.Drop {
	// 	result.Start, err = getStartPosition(move, commaIndex, game.Board.Height)
	// 	if err != nil {
	// 		return result, fmt.Errorf("Invalid Start Position")
	// 	}
	//
	// }
	//
	// result.Promote, result.EndPiece = checkPromote(move, game)
	//
	// result.End, err = getEndPosition(move, commaIndex, game.Board.Height)
	//
	// if err != nil {
	// 	return result, fmt.Errorf("Invalid End Position")
	// }
	return result, nil
}

func convertStringToPosition(move string, boardHeight int) (Vec2, error) {
	var result Vec2

	xStr := ""
	yStr := ""

	for i := 0; i < len(move); i++ {
		char := move[i]
		if utils.IsDigit(char) {
			yStr += string(char)
		} else if utils.IsLower(char) {
			xStr += string(char)
		}
	}

	x, err := utils.ConvertLowercaseToNumber(xStr)
	if err != nil {
		return result, err
	}
	x-- //index base zero

	y, err := strconv.Atoi(yStr)
	if err != nil {
		return result, err
	}

	result.X = x
	result.Y = boardHeight - y

	return result, nil
}

func checkDropPiece(move string) int {
	if len(move) < 2 {
		return 0
	}

	if move[1] != '*' {
		return 0
	}

	komaChar := move[0]
	koma, ok := shogiDropCharToPiece[komaChar]
	if !ok {
		return 0
	}

	return koma
}

func checkPromotePiece(move string) int {

	moveLength := len(move)
	if moveLength < 1 {
		return 0
	}

	if move[moveLength-1] == '+' { //if shogi promotion or chess
		return 0
	}

	pieceChar := move[moveLength-1]
	piece, ok := chessPromoteCharToPiece[pieceChar]
	if !ok {
		return 0
	}

	return piece
}

//
// func ConvertMoveToString(move Move, game Game) (string, error) {
// 	var result string
//
// 	startStr := ""
// 	endStr := ""
// 	promoteStr := ""
//
// 	endWidthStr, err := utils.ConvertNumberToLowercase(move.End.X + 1)
// 	if err != nil {
// 		return "", err
// 	}
// 	endHeightStr := strconv.Itoa(game.Board.Height - move.End.Y)
// 	endStr = endWidthStr + endHeightStr
//
// 	startWidthStr, err := utils.ConvertNumberToLowercase(move.Start.X + 1)
// 	if err != nil {
// 		return "", err
// 	}
// 	startHeightStr := strconv.Itoa(game.Board.Height - move.Start.Y)
// 	startStr = startWidthStr + startHeightStr
//
// 	if move.Drop {
// 		dropPiece := move.StartPiece.Type
// 		pieceChar, ok := shogiDropPieceToChar[dropPiece]
// 		if !ok {
// 			return "", fmt.Errorf("Invalid Drop Piece")
// 		}
//
// 		startStr = string(pieceChar) + "*"
// 	}
//
// 	if move.Promote {
// 		if move.EndPiece.Type != 0 {
// 			promotePiece, ok := chessPromotePieceToChar[move.EndPiece.Type]
// 			if !ok {
// 				return "", fmt.Errorf("Invalid Promote Piece")
// 			}
// 			promoteStr = string(promotePiece)
// 		} else {
// 			promoteStr = "+"
// 		}
// 	}
//
// 	result = startStr + "," + endStr + promoteStr
//
// 	return result, nil
// }
