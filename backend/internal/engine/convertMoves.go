package engine

import (
	"fmt"
	"github.com/KainoaGardner/csc/internal/utils"
	"strconv"
	"strings"
)

func ConvertStringToMoves(move string, game Game) ([]Move, error) {
	var result []Move

	moveStrings := strings.Split(move, ",")
	for i := 0; i < len(moveStrings)-1; i++ {
		moveStartString := moveStrings[i]
		moveEndString := moveStrings[i+1]
		var move Move

		drop := checkDropPiece(moveStartString)
		if drop == nil {
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

func checkDropPiece(move string) *int {
	var result *int

	if len(move) < 2 {
		return nil
	}

	if move[1] != '*' {
		return nil
	}

	komaChar := move[0]
	koma, ok := shogiDropCharToPiece[komaChar]
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

func ConvertMovesToString(moves []Move, game Game) (string, error) {
	result := ""

	movesLength := len(moves)

	//add start
	if movesLength < 1 {
		return result, fmt.Errorf("No moves")
	}

	startStr, err := convertStartMoveToString(moves[0], game)
	if err != nil {
		return result, err
	}

	result += startStr

	//add ends
	for i := 0; i < len(moves); i++ {
		endStr, err := convertEndMoveToString(moves[i], game)
		if err != nil {
			return result, err
		}
		result += "," + endStr
	}

	return result, nil
}

func convertStartMoveToString(move Move, game Game) (string, error) {
	result := ""

	startWidthStr, err := utils.ConvertNumberToLowercase(move.Start.X + 1)

	if err != nil {
		return "", err
	}

	startHeightStr := strconv.Itoa(game.Board.Height - move.Start.Y)
	result = startWidthStr + startHeightStr

	if move.Drop != nil {
		dropPiece := *move.Drop
		pieceChar, ok := shogiDropPieceToChar[dropPiece]
		if !ok {
			return "", fmt.Errorf("Invalid Drop Piece")
		}

		result = string(pieceChar) + "*"
	}

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
