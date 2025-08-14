package engine

import (
	"fmt"

	"github.com/KainoaGardner/csc/internal/types"
	"github.com/KainoaGardner/csc/internal/utils"
)

func checkValidPieceMoves(move types.Move, piece types.Piece, game types.Game) error {
	dir := getMoveDirection(game)
	possibleMoves := getPieceMoves(move.Start, piece, game, dir)
	filterPossibleMoves(move.Start, &possibleMoves, game)
	return checkEndPosInPossibleMoves(possibleMoves, move)
}

// change move to Start types.Vec2
func getPieceMoves(pos types.Vec2, piece types.Piece, game types.Game, dir int) []types.Vec2 {
	var possibleMoves []types.Vec2
	switch piece.Type {
	case types.Pawn:
		possibleMoves = getPawnMoves(pos, piece, game, dir)
	case types.Knight:
		possibleMoves = getKnightMoves(pos, piece, game)
	case types.Bishop:
		possibleMoves = getBishopMoves(pos, piece, game)
	case types.Rook:
		possibleMoves = getRookMoves(pos, piece, game)
	case types.Queen:
		possibleMoves = getQueenMoves(pos, piece, game)
	case types.King:
		possibleMoves = getKingMoves(pos, piece, game)
		possibleMoves = append(possibleMoves, getCastleMoves(pos, piece, game)...)
	case types.Fu:
		possibleMoves = getFuMoves(pos, piece, game, dir)
	case types.Kyou:
		possibleMoves = getKyouMoves(pos, piece, game, dir)
	case types.Kei:
		possibleMoves = getKeiMoves(pos, piece, game, dir)
	case types.Gin:
		possibleMoves = getGinMoves(pos, piece, game, dir)
	case types.Kin:
		possibleMoves = getKinMoves(pos, piece, game, dir)
	case types.Kaku:
		possibleMoves = getBishopMoves(pos, piece, game)
	case types.Hi:
		possibleMoves = getRookMoves(pos, piece, game)
	case types.Ou:
		possibleMoves = getKingMoves(pos, piece, game)
	case types.To:
		possibleMoves = getKinMoves(pos, piece, game, dir)
	case types.NariKyou:
		possibleMoves = getKinMoves(pos, piece, game, dir)
	case types.NariKei:
		possibleMoves = getKinMoves(pos, piece, game, dir)
	case types.NariGin:
		possibleMoves = getKinMoves(pos, piece, game, dir)
	case types.Uma:
		possibleMoves = getUmaMoves(pos, piece, game)
	case types.Ryuu:
		possibleMoves = getRyuuMoves(pos, piece, game)
	case types.Checker:
		possibleMoves = getCheckerMoves(pos, piece, game, dir)
	case types.CheckerKing:
		possibleMoves = getCheckerKingMoves(pos, piece, game)
	}

	return possibleMoves
}

func getPawnMoves(pos types.Vec2, piece types.Piece, game types.Game, direction int) []types.Vec2 {
	var validMovePositions []types.Vec2
	//move forward
	newY := pos.Y - 1*direction
	newPos := types.Vec2{X: pos.X, Y: newY}
	newPos2 := types.Vec2{X: pos.X, Y: newY - direction}
	if checkPositionInbounds(newPos, game) {
		space := game.Board.Board[newPos.Y][newPos.X]
		if space == nil {
			validMovePositions = append(validMovePositions, newPos)

			//check starting move 2 space
			if checkPositionInbounds(newPos2, game) {
				space = game.Board.Board[newPos2.Y][newPos2.X]
				if space == nil && !piece.Moved {
					validMovePositions = append(validMovePositions, newPos2)
				}
			}
		}
	}

	//capture squares
	relativeMovePos := []types.Vec2{{X: -1, Y: -1}, {X: 1, Y: -1}}
	for i := 0; i < len(relativeMovePos); i++ {
		newPos = relativeMovePos[i]
		newPos.Y *= direction

		newPos.X += pos.X
		newPos.Y += pos.Y
		if checkPositionInbounds(newPos, game) {
			space := game.Board.Board[newPos.Y][newPos.X]
			if space != nil && space.Owner == piece.Owner {
				validMovePositions = append(validMovePositions, newPos)
			} else if space == nil && game.EnPassant != nil && utils.CheckVec2Equal(newPos, *game.EnPassant) {
				validMovePositions = append(validMovePositions, newPos)
			}
		}
	}

	return validMovePositions
}

func getKnightMoves(pos types.Vec2, piece types.Piece, game types.Game) []types.Vec2 {
	var validMovePositions []types.Vec2

	relativeMovePos := []types.Vec2{
		{X: -1, Y: -2},
		{X: 1, Y: -2},
		{X: -2, Y: -1},
		{X: 2, Y: -1},
		{X: -2, Y: 1},
		{X: 2, Y: 1},
		{X: -1, Y: 2},
		{X: 1, Y: 2},
	}

	for i := 0; i < len(relativeMovePos); i++ {
		newPos := relativeMovePos[i]
		newPos.X += pos.X
		newPos.Y += pos.Y

		if checkPositionInbounds(newPos, game) {
			space := game.Board.Board[newPos.Y][newPos.X]
			if space == nil || space.Owner != piece.Owner {
				validMovePositions = append(validMovePositions, newPos)
			}
		}
	}

	return validMovePositions
}

func getBishopMoves(pos types.Vec2, piece types.Piece, game types.Game) []types.Vec2 {
	var validMovePositions []types.Vec2

	directions := []types.Vec2{
		{X: -1, Y: -1},
		{X: -1, Y: 1},
		{X: 1, Y: -1},
		{X: 1, Y: 1},
	}

	for i := 0; i < len(directions); i++ {
		dir := directions[i]

		j := 0
		for j >= 0 {
			j++
			newPos := pos

			newPos.X += dir.X * j
			newPos.Y += dir.Y * j

			if !checkPositionInbounds(newPos, game) {
				break
			}

			space := game.Board.Board[newPos.Y][newPos.X]
			if space == nil {
				validMovePositions = append(validMovePositions, newPos)
			} else if space.Owner != piece.Owner {
				validMovePositions = append(validMovePositions, newPos)
				break
			} else {
				break
			}
		}
	}

	return validMovePositions
}

func getRookMoves(pos types.Vec2, piece types.Piece, game types.Game) []types.Vec2 {
	var validMovePositions []types.Vec2

	directions := []types.Vec2{
		{X: -1, Y: 0},
		{X: 1, Y: 0},
		{X: 0, Y: -1},
		{X: 0, Y: 1},
	}

	for i := 0; i < len(directions); i++ {
		dir := directions[i]

		j := 0
		for j >= 0 {
			j++
			newPos := pos

			newPos.X += dir.X * j
			newPos.Y += dir.Y * j

			if !checkPositionInbounds(newPos, game) {
				break
			}

			space := game.Board.Board[newPos.Y][newPos.X]
			if space == nil {
				validMovePositions = append(validMovePositions, newPos)
			} else if space.Owner != piece.Owner {
				validMovePositions = append(validMovePositions, newPos)
				break
			} else {
				break
			}
		}
	}
	return validMovePositions
}

func getQueenMoves(pos types.Vec2, piece types.Piece, game types.Game) []types.Vec2 {
	var validMovePositions []types.Vec2

	bishopMoves := getBishopMoves(pos, piece, game)
	rookMoves := getRookMoves(pos, piece, game)
	validMovePositions = append(bishopMoves, rookMoves...)

	return validMovePositions
}

func getKingMoves(pos types.Vec2, piece types.Piece, game types.Game) []types.Vec2 {
	var validMovePositions []types.Vec2

	directions := []types.Vec2{
		{X: -1, Y: -1},
		{X: 0, Y: -1},
		{X: 1, Y: -1},
		{X: -1, Y: 0},
		{X: 1, Y: 0},
		{X: -1, Y: 1},
		{X: 0, Y: 1},
		{X: 1, Y: 1},
	}

	for i := 0; i < len(directions); i++ {
		dir := directions[i]

		newPos := pos
		newPos.X += dir.X
		newPos.Y += dir.Y

		if checkPositionInbounds(newPos, game) {
			space := game.Board.Board[newPos.Y][newPos.X]
			if space == nil || space.Owner != piece.Owner {
				validMovePositions = append(validMovePositions, newPos)
			}
		}
	}

	return validMovePositions
}

func getFuMoves(pos types.Vec2, piece types.Piece, game types.Game, direction int) []types.Vec2 {
	var validMovePositions []types.Vec2

	newPos := pos
	newPos.Y += -1 * direction

	if checkPositionInbounds(newPos, game) {
		space := game.Board.Board[newPos.Y][newPos.X]
		if space == nil || space.Owner != piece.Owner {
			validMovePositions = append(validMovePositions, newPos)
		}
	}

	return validMovePositions
}

func getKyouMoves(pos types.Vec2, piece types.Piece, game types.Game, dir int) []types.Vec2 {
	var validMovePositions []types.Vec2

	i := 0
	for i >= 0 {
		i++
		newPos := pos

		newPos.Y += -i * dir

		if !checkPositionInbounds(newPos, game) {
			break
		}

		space := game.Board.Board[newPos.Y][newPos.X]
		if space == nil {
			validMovePositions = append(validMovePositions, newPos)
		} else if space.Owner != piece.Owner {
			validMovePositions = append(validMovePositions, newPos)
			break
		} else {
			break
		}
	}

	return validMovePositions
}

func getKeiMoves(pos types.Vec2, piece types.Piece, game types.Game, direction int) []types.Vec2 {
	var validMovePositions []types.Vec2

	directions := []types.Vec2{
		{X: -1, Y: -2 * direction},
		{X: 1, Y: -2 * direction},
	}

	for i := 0; i < len(directions); i++ {
		dir := directions[i]

		newPos := pos
		newPos.X += dir.X
		newPos.Y += dir.Y

		if checkPositionInbounds(newPos, game) {
			space := game.Board.Board[newPos.Y][newPos.X]
			if space == nil || space.Owner != piece.Owner {
				validMovePositions = append(validMovePositions, newPos)
			}
		}
	}

	return validMovePositions
}

func getGinMoves(pos types.Vec2, piece types.Piece, game types.Game, dir int) []types.Vec2 {
	var validMovePositions []types.Vec2

	directions := []types.Vec2{
		{X: -1, Y: -1 * dir},
		{X: 0, Y: -1 * dir},
		{X: 1, Y: -1 * dir},
		{X: -1, Y: 1 * dir},
		{X: 1, Y: 1 * dir},
	}

	for i := 0; i < len(directions); i++ {
		dir := directions[i]

		newPos := pos
		newPos.X += dir.X
		newPos.Y += dir.Y

		if checkPositionInbounds(newPos, game) {
			space := game.Board.Board[newPos.Y][newPos.X]
			if space == nil || space.Owner != piece.Owner {
				validMovePositions = append(validMovePositions, newPos)
			}
		}

	}
	return validMovePositions
}

func getKinMoves(pos types.Vec2, piece types.Piece, game types.Game, dir int) []types.Vec2 {
	var validMovePositions []types.Vec2

	directions := []types.Vec2{
		{X: -1, Y: -1 * dir},
		{X: 0, Y: -1 * dir},
		{X: 1, Y: -1 * dir},
		{X: -1, Y: 0},
		{X: 1, Y: 0},
		{X: 0, Y: 1 * dir},
	}

	for i := 0; i < len(directions); i++ {
		dir := directions[i]

		newPos := pos
		newPos.X += dir.X
		newPos.Y += dir.Y

		if checkPositionInbounds(newPos, game) {
			space := game.Board.Board[newPos.Y][newPos.X]
			if space == nil || space.Owner != piece.Owner {
				validMovePositions = append(validMovePositions, newPos)
			}
		}

	}
	return validMovePositions
}

func getUmaMoves(pos types.Vec2, piece types.Piece, game types.Game) []types.Vec2 {
	var validMovePositions []types.Vec2

	directions := []types.Vec2{
		{X: -1, Y: 0},
		{X: 1, Y: 0},
		{X: 0, Y: -1},
		{X: 0, Y: 1},
	}

	for i := 0; i < len(directions); i++ {
		dir := directions[i]

		newPos := pos
		newPos.X += dir.X
		newPos.Y += dir.Y

		if checkPositionInbounds(newPos, game) {
			space := game.Board.Board[newPos.Y][newPos.X]
			if space == nil || space.Owner != piece.Owner {
				validMovePositions = append(validMovePositions, newPos)
			}
		}
	}

	bishopMoves := getBishopMoves(pos, piece, game)
	validMovePositions = append(validMovePositions, bishopMoves...)
	return validMovePositions
}

func getRyuuMoves(pos types.Vec2, piece types.Piece, game types.Game) []types.Vec2 {
	var validMovePositions []types.Vec2

	directions := []types.Vec2{
		{X: -1, Y: -1},
		{X: 1, Y: -1},
		{X: -1, Y: 1},
		{X: 1, Y: 1},
	}

	for i := 0; i < len(directions); i++ {
		dir := directions[i]

		newPos := pos
		newPos.X += dir.X
		newPos.Y += dir.Y

		if checkPositionInbounds(newPos, game) {
			space := game.Board.Board[newPos.Y][newPos.X]
			if space == nil || space.Owner != piece.Owner {
				validMovePositions = append(validMovePositions, newPos)
			}
		}
	}

	rookMoves := getRookMoves(pos, piece, game)
	validMovePositions = append(validMovePositions, rookMoves...)
	return validMovePositions
}

func getCheckerMoves(pos types.Vec2, piece types.Piece, game types.Game, dir int) []types.Vec2 {
	var validMovePositions []types.Vec2

	directions := []types.Vec2{
		{X: -1, Y: -1 * dir},
		{X: 1, Y: -1 * dir},
	}

	for i := 0; i < len(directions); i++ {
		dir := directions[i]

		jumpPos := pos
		landPos := pos

		jumpPos.X += dir.X
		jumpPos.Y += dir.Y

		landPos.X += dir.X * 2
		landPos.Y += dir.Y * 2

		if checkPositionInbounds(jumpPos, game) {
			jumpSpace := game.Board.Board[jumpPos.Y][jumpPos.X]
			if jumpSpace == nil {
				validMovePositions = append(validMovePositions, jumpPos)
			}

			if checkPositionInbounds(landPos, game) {
				landSpace := game.Board.Board[landPos.Y][landPos.X]
				if landSpace == nil && (jumpSpace != nil && jumpSpace.Owner != piece.Owner) {
					validMovePositions = append(validMovePositions, landPos)
				}

			}
		}
	}

	return validMovePositions
}

func getCheckerKingMoves(pos types.Vec2, piece types.Piece, game types.Game) []types.Vec2 {
	var validMovePositions []types.Vec2

	directions := []types.Vec2{
		{X: -1, Y: -1},
		{X: 1, Y: -1},
		{X: -1, Y: 1},
		{X: 1, Y: 1},
	}

	for i := 0; i < len(directions); i++ {
		dir := directions[i]
		jumpPos := pos
		landPos := pos

		jumpPos.X += dir.X
		jumpPos.Y += dir.Y

		landPos.X += dir.X * 2
		landPos.Y += dir.Y * 2

		if checkPositionInbounds(jumpPos, game) {
			jumpSpace := game.Board.Board[jumpPos.Y][jumpPos.X]
			if jumpSpace == nil {
				validMovePositions = append(validMovePositions, jumpPos)
			}

			if checkPositionInbounds(landPos, game) {
				landSpace := game.Board.Board[landPos.Y][landPos.X]
				if landSpace == nil && (jumpSpace != nil && jumpSpace.Owner != piece.Owner) {
					validMovePositions = append(validMovePositions, landPos)
				}
			}
		}
	}

	return validMovePositions
}

func checkEndPosInPossibleMoves(possibleMoves []types.Vec2, move types.Move) error {
	for i := 0; i < len(possibleMoves); i++ {
		possibleMove := possibleMoves[i]

		if utils.CheckVec2Equal(move.End, possibleMove) {
			return nil
		}
	}

	return fmt.Errorf("No valid moves")
}

func getMoveDirection(game types.Game) int {
	direction := 1
	if game.Turn == 1 {
		direction = -1
	}
	return direction
}

func getCastleMoves(pos types.Vec2, piece types.Piece, game types.Game) []types.Vec2 {
	var validMovePositions []types.Vec2

	//left
	for i := pos.X - 1; i >= 0; i-- {
		targetPiece := game.Board.Board[pos.Y][i]
		if targetPiece != nil {
			if targetPiece.Type == types.Rook && targetPiece.Owner == piece.Owner {
				validMovePositions = append(validMovePositions, types.Vec2{X: i, Y: pos.Y})
			}
			break
		}
	}

	for i := pos.X + 1; i < game.Board.Width; i++ {
		targetPiece := game.Board.Board[pos.Y][i]
		if targetPiece != nil {
			if targetPiece.Type == types.Rook && targetPiece.Owner == piece.Owner {
				validMovePositions = append(validMovePositions, types.Vec2{X: i, Y: pos.Y})
			}
			break
		}
	}

	return validMovePositions
}

func checkCheckerNextJumps(startPos types.Vec2, endPos types.Vec2, piece types.Piece, game types.Game) bool {
	if !checkCheckerTake(startPos, endPos) {
		return false
	}

	dir := getMoveDirection(game)
	var possibleMoves []types.Vec2
	switch piece.Type {
	case types.Checker:
		possibleMoves = getCheckerMoves(endPos, piece, game, dir)
	case types.CheckerKing:
		possibleMoves = getCheckerKingMoves(endPos, piece, game)
	default:
		return false
	}

	for i := len(possibleMoves) - 1; i >= 0; i-- {
		if !checkCheckerTake(endPos, possibleMoves[i]) {
			possibleMoves = append(possibleMoves[:i], possibleMoves[i+1:]...)
		}
	}

	return len(possibleMoves) > 0
}

func checkCheckerTake(startPos types.Vec2, endPos types.Vec2) bool {
	dx := utils.AbsoluteValueInt(startPos.X - endPos.X)
	dy := utils.AbsoluteValueInt(startPos.Y - endPos.Y)

	if dx == 2 && dy == 2 {
		return true
	}

	return false
}

func filterPossibleMoves(startPos types.Vec2, possibleMoves *[]types.Vec2, game types.Game) {
	for i := len(*possibleMoves) - 1; i >= 0; i-- {
		movePos := (*possibleMoves)[i]
		gameCopy := copyGame(game)
		piece := gameCopy.Board.Board[startPos.Y][startPos.X]
		if piece != nil && piece.Type >= types.Pawn && piece.Type <= types.Ryuu {
			takePiece := gameCopy.Board.Board[movePos.Y][movePos.X]
			validCastle := takePiece != nil && piece.Type == types.King && takePiece.Type == types.Rook && takePiece.Owner == piece.Owner
			move := types.Move{
				Start:   startPos,
				End:     movePos,
				Promote: nil,
				Drop:    nil,
			}
			if validCastle {
				castleDir := getCastleDirection(move)
				dx := utils.AbsoluteValueInt(move.End.X-move.Start.X) - 1
				kingX := (dx/2+1)*castleDir + move.Start.X
				startToKingDist := utils.AbsoluteValueInt(startPos.X - kingX)
				for j := 0; j < startToKingDist+1; j++ {
					gameCopy.Board.Board[startPos.Y][startPos.X] = nil
					gameCopy.Board.Board[movePos.Y][startPos.X+j*castleDir] = piece
					if GetInCheck(*gameCopy) {
						*possibleMoves = append((*possibleMoves)[:i], (*possibleMoves)[i+1:]...)
						break
					}
				}
			} else {
				dir := getMoveDirection(*gameCopy)
				if checkEnPassantTake(move, *gameCopy, piece) {
					game.Board.Board[move.End.Y+dir][move.End.X] = nil
				}
				gameCopy.Board.Board[startPos.Y][startPos.X] = nil
				gameCopy.Board.Board[movePos.Y][movePos.X] = piece
				if GetInCheck(*gameCopy) {
					*possibleMoves = append((*possibleMoves)[:i], (*possibleMoves)[i+1:]...)
				}
			}
		} else if piece != nil && piece.Type >= types.Checker && piece.Type <= types.CheckerKing {
			if checkerMovesInCheck(startPos, movePos, piece, *gameCopy) {
				*possibleMoves = append((*possibleMoves)[:i], (*possibleMoves)[i+1:]...)
			}
		}
	}
}

func checkerMovesInCheck(startPos types.Vec2, endPos types.Vec2, piece *types.Piece, game types.Game) bool {
	if !checkCheckerNextJumps(startPos, endPos, *piece, game) {
		game.Board.Board[startPos.Y][startPos.X] = nil
		game.Board.Board[endPos.Y][endPos.X] = piece
		if GetInCheck(game) {
			return true
		} else {
			return false
		}
	} else {

		dir := getMoveDirection(game)
		var possibleMoves []types.Vec2
		switch piece.Type {
		case types.Checker:
			possibleMoves = getCheckerMoves(endPos, *piece, game, dir)
		case types.CheckerKing:
			possibleMoves = getCheckerKingMoves(endPos, *piece, game)
		}
		for i := len(possibleMoves) - 1; i >= 0; i-- {
			if !checkCheckerTake(endPos, possibleMoves[i]) {
				possibleMoves = append(possibleMoves[:i], possibleMoves[i+1:]...)
			}
		}

		for _, movePos := range possibleMoves {
			gameCopy := copyGame(game)
			gameCopy.Board.Board[startPos.Y][startPos.X] = nil
			gameCopy.Board.Board[endPos.Y][endPos.X] = piece

			result := checkerMovesInCheck(endPos, movePos, piece, *gameCopy)

			if !result {
				return false
			}

		}

	}

	return true
}
