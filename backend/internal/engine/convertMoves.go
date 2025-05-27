package engine

import (
	"fmt"
	"github.com/KainoaGardner/csc/internal/utils"
	"strconv"
)

func ConvertStringToMove(move string, game Game) (Move, error) {
	result := Move{}
	var err error

	commaIndex := utils.GetIndexFirstChar(move, ",")
	if commaIndex == -1 {
		return result, fmt.Errorf("Invalid Format. No ,")
	}

	result.Drop, result.StartPiece = checkDropPiece(move, game)
	if !result.Drop {
		result.Start, err = getStartPosition(move, commaIndex, game.Board.Height)
		if err != nil {
			return result, fmt.Errorf("Invalid Start Position")
		}

	}

	result.Promote, result.EndPiece = checkPromote(move, game)

	result.End, err = getEndPosition(move, commaIndex, game.Board.Height)

	if err != nil {
		return result, fmt.Errorf("Invalid End Position")
	}

	return result, nil
}

func getStartPosition(move string, commaIndex int, boardHeight int) (Vec2, error) {
	result := Vec2{}

	startWidthStr := ""
	startHeightStr := ""

	for i := 0; i < commaIndex; i++ {
		char := move[i]
		if utils.IsDigit(char) {
			startHeightStr += string(char)
		} else if utils.IsLower(char) {
			startWidthStr += string(char)
		}
	}

	startWidth, err := utils.ConvertLowercaseToNumber(startWidthStr)
	if err != nil {
		return result, err
	}
	startWidth-- //index base zero

	startHeight, err := strconv.Atoi(startHeightStr)
	if err != nil {
		return result, err
	}

	result.X = startWidth
	result.Y = boardHeight - startHeight

	return result, nil
}

func getEndPosition(move string, commaIndex int, boardHeight int) (Vec2, error) {
	result := Vec2{}

	endWidthStr := ""
	endHeightStr := ""

	for i := commaIndex + 1; i < len(move); i++ {
		char := move[i]
		if utils.IsDigit(char) {
			endHeightStr += string(char)
		} else if utils.IsLower(char) {
			endWidthStr += string(char)
		}
	}

	endWidth, err := utils.ConvertLowercaseToNumber(endWidthStr)
	if err != nil {
		return result, err
	}

	endHeight, err := strconv.Atoi(endHeightStr)
	if err != nil {
		return result, err
	}
	endWidth-- //index base zero

	result.X = endWidth
	result.Y = boardHeight - endHeight

	return result, nil
}

func checkDropPiece(move string, game Game) (bool, Piece) {
	var koma Piece
	if len(move) < 2 {
		return false, koma
	}

	if move[1] != '*' {
		return false, koma
	}

	komaChar := move[0]
	komaInt, ok := shogiDropCharToPiece[komaChar]
	koma.Type = komaInt
	koma.Owner = game.Turn
	if !ok {
		return false, koma
	}

	return true, koma
}

func checkPromote(move string, game Game) (bool, Piece) {
	var piece Piece

	moveLength := len(move)
	if moveLength < 1 {
		return false, piece
	}

	if move[moveLength-1] == '+' { //if shogi promotion
		return true, piece
	}

	pieceChar := move[moveLength-1]
	pieceInt, ok := chessPromoteCharToPiece[pieceChar]
	piece.Type = pieceInt
	piece.Owner = game.Turn
	if !ok {
		return false, piece
	}

	return true, piece
}

func ConvertMoveToString(move Move, game Game) (string, error) {
	var result string

	startStr := ""
	endStr := ""
	promoteStr := ""

	endWidthStr, err := utils.ConvertNumberToLowercase(move.End.X + 1)
	if err != nil {
		return "", err
	}
	endHeightStr := strconv.Itoa(game.Board.Height - move.End.Y)
	endStr = endWidthStr + endHeightStr

	startWidthStr, err := utils.ConvertNumberToLowercase(move.Start.X + 1)
	if err != nil {
		return "", err
	}
	startHeightStr := strconv.Itoa(game.Board.Height - move.Start.Y)
	startStr = startWidthStr + startHeightStr

	if move.Drop {
		dropPiece := move.StartPiece.Type
		pieceChar, ok := shogiDropPieceToChar[dropPiece]
		if !ok {
			return "", fmt.Errorf("Invalid Drop Piece")
		}

		startStr = string(pieceChar) + "*"
	}

	if move.Promote {
		if move.EndPiece.Type != 0 {
			promotePiece, ok := chessPromotePieceToChar[move.EndPiece.Type]
			if !ok {
				return "", fmt.Errorf("Invalid Promote Piece")
			}
			promoteStr = string(promotePiece)
		} else {
			promoteStr = "+"
		}
	}

	result = startStr + "," + endStr + promoteStr

	return result, nil
}
