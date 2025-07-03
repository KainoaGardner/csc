package engine

import (
	"fmt"
	"github.com/KainoaGardner/csc/internal/types"
	"github.com/KainoaGardner/csc/internal/utils"
	"strconv"
	"strings"
)

func ConvertStringToMove(moveString string, game Game) (Move, error) {
	var result Move

	moveStrings := strings.Split(moveString, ",")
	if len(moveStrings) != 2 {
		return result, fmt.Errorf("Must have 2 move positions. A0,A0")
	}

	moveStartString := moveStrings[0]
	moveEndString := moveStrings[1]

	drop := checkDropPiece(moveStartString)
	if drop == nil {
		start, err := convertStringToPosition(moveStartString, game.Board.Height)
		if err != nil {
			return result, err
		}
		result.Start = start
	} else {
		result.Drop = drop
	}

	promote := checkPromotePiece(moveEndString)
	result.Promote = promote

	if result.Drop != nil && result.Promote != nil {
		return result, fmt.Errorf("Can't promote when dropping")
	}

	end, err := convertStringToPosition(moveEndString, game.Board.Height)
	if err != nil {
		return result, err
	}
	result.End = end

	return result, nil
}

func convertStringToPosition(move string, boardHeight int) (types.Vec2, error) {
	var result types.Vec2

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

func checkDropPiece(move string) *int {
	var result *int

	if len(move) < 2 {
		return nil
	}

	if move[1] != '*' {
		return nil
	}

	komaChar := move[0]
	koma, ok := shogiDropCharToMochiPiece[komaChar]
	if !ok {
		return nil
	}

	result = &koma
	return result
}

func checkPromotePiece(move string) *int {
	var result *int

	moveLength := len(move)
	if moveLength < 1 {
		return result
	}

	if move[moveLength-1] == '+' { //if shogi promotion or chess
		found := 0
		result = &found
		return result
	}

	pieceChar := move[moveLength-1]
	piece, ok := chessPromoteCharToPiece[pieceChar]
	if !ok {
		return result
	}
	result = &piece

	return result
}

func ConvertMoveToString(move Move, game Game) (string, error) {
	result := ""

	startStr, err := convertStartMoveToString(move, game)
	if err != nil {
		return result, err
	}

	result += startStr

	endStr, err := convertEndMoveToString(move, game)
	if err != nil {
		return result, err
	}
	result += "," + endStr

	return result, nil
}

func convertStartMoveToString(move Move, game Game) (string, error) {
	result, err := convertPositionToString(move.Start, game)
	if err != nil {
		return result, err
	}

	if move.Drop != nil {
		dropPiece := *move.Drop
		pieceChar, ok := shogiMochiPieceToChar[dropPiece]
		if !ok {
			return "", fmt.Errorf("Invalid Drop Piece")
		}

		result = string(pieceChar) + "*"
	}

	return result, nil
}

func convertPositionToString(pos types.Vec2, game Game) (string, error) {
	result := ""

	startWidthStr, err := utils.ConvertNumberToLowercase(pos.X + 1)

	if err != nil {
		return "", err
	}

	startHeightStr := strconv.Itoa(game.Board.Height - pos.Y)
	result = startWidthStr + startHeightStr

	return result, nil
}

func convertEndMoveToString(move Move, game Game) (string, error) {
	result := ""

	endWidthStr, err := utils.ConvertNumberToLowercase(move.End.X + 1)

	if err != nil {
		return "", err
	}

	endHeightStr := strconv.Itoa(game.Board.Height - move.End.Y)
	result = endWidthStr + endHeightStr

	if move.Promote != nil {
		if *move.Promote != 0 {
			promotePiece, ok := chessPromotePieceToChar[*move.Promote]
			if !ok {
				return "", fmt.Errorf("Invalid Promote Piece")
			}
			result += string(promotePiece)
		} else {
			result += "+"
		}
	}

	return result, nil
}
