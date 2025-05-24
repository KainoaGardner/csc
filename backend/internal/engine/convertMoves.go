package engine

import (
	"fmt"
	"github.com/KainoaGardner/csc/internal/utils"
	"strconv"
)

func ConvertStringToMove(move string, game Game) (Move, error) {
	result := Move{}
	var err error

	moveLength := len(move)
	if moveLength != 5 && moveLength != 6 {
		return result, fmt.Errorf("Invalid Length")
	}

	commaIndex := utils.GetIndexFirstChar(move, ",")
	if commaIndex == -1 {
		return result, fmt.Errorf("Invalid Format. No ,")
	}

	result.Drop, result.DropPiece = checkDropPiece(move, game.Turn)
	result.MovePiece = result.DropPiece
	if !result.Drop {
		result.Start, err = getStartPosition(move, commaIndex, game.Board.Height)
		if err != nil {
			return result, fmt.Errorf("Invalid Start Position")
		}

		if result.Start[0] >= 0 && result.Start[0] < game.Board.Width && result.Start[1] >= 0 && result.Start[1] < game.Board.Height {
			result.MovePiece = game.Board.Board[result.Start[1]][result.Start[0]]
		}

	}

	result.Promote, result.PromotePiece = checkPromote(move, game.Turn)

	result.End, err = getEndPosition(move, commaIndex, game.Board.Height)
	if result.End[0] >= 0 && result.End[0] < game.Board.Width && result.End[1] >= 0 && result.End[1] < game.Board.Height {
		result.TakenPiece = game.Board.Board[result.End[1]][result.End[0]]
	}

	if err != nil {
		return result, fmt.Errorf("Invalid End Position")
	}

	return result, nil
}

func getStartPosition(move string, commaIndex int, boardHeight int) ([2]int, error) {
	result := [2]int{}

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

	startHeight, err := strconv.Atoi(startHeightStr)
	if err != nil {
		return result, err
	}

	result[0] = startWidth
	result[1] = boardHeight - startHeight

	return result, nil
}

func getEndPosition(move string, commaIndex int, boardHeight int) ([2]int, error) {
	result := [2]int{}

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

	result[0] = endWidth
	result[1] = boardHeight - endHeight

	return result, nil
}

func checkDropPiece(move string, turn int) (bool, int) {
	var koma int
	if len(move) < 2 {
		return false, 0
	}

	if move[1] != '*' {
		return false, 0
	}

	komaChar := move[0]
	koma, ok := shogiDropCharToPiece[komaChar]
	if !ok {
		return false, 0
	}

	if turn == 1 {
		koma += MochigomaBlackOffset
	}

	return true, koma
}

func checkPromote(move string, turn int) (bool, int) {
	var piece int
	moveLength := len(move)
	if moveLength < 1 {
		return false, 0
	}

	if move[moveLength-1] == '+' { //if shogi promotion
		return true, 0
	}

	pieceChar := move[moveLength-1]
	piece, ok := chessPromoteCharToPiece[pieceChar]
	if !ok {
		return false, 0
	}

	if turn == 1 {
		piece *= -1
	}

	return true, piece
}

func ConvertMoveToString(move Move) (string, error) {

	var result string

	return result, nil
}
