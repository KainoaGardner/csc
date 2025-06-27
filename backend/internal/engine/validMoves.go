package engine

import "fmt"

func checkValidPieceMoves(move Move, piece Piece, game Game) error {
	dir := getMoveDirection(game)
	possibleMoves := getPieceMoves(move, piece, game, dir)
	return checkEndPosInPossibleMoves(possibleMoves, move)
}

func getPieceMoves(move Move, piece Piece, game Game, dir int) []Vec2 {
	var possibleMoves []Vec2
	switch piece.Type {
	case Pawn:
		possibleMoves = getPawnMoves(move, piece, game, dir)
	case Knight:
		possibleMoves = getKnightMoves(move, piece, game)
	case Bishop:
		possibleMoves = getBishopMoves(move, piece, game)
	case Rook:
		possibleMoves = getRookMoves(move, piece, game)
	case Queen:
		possibleMoves = getQueenMoves(move, piece, game)
	case King:
		possibleMoves = getKingMoves(move, piece, game)
		possibleMoves = append(possibleMoves, getCastleMoves(move, piece, game)...)
	case Fu:
		possibleMoves = getFuMoves(move, piece, game, dir)
	case Kyou:
		possibleMoves = getKyouMoves(move, piece, game, dir)
	case Kei:
		possibleMoves = getKeiMoves(move, piece, game, dir)
	case Gin:
		possibleMoves = getGinMoves(move, piece, game, dir)
	case Kin:
		possibleMoves = getKinMoves(move, piece, game, dir)
	case Kaku:
		possibleMoves = getBishopMoves(move, piece, game)
	case Hi:
		possibleMoves = getRookMoves(move, piece, game)
	case Ou:
		possibleMoves = getKingMoves(move, piece, game)
	case To:
		possibleMoves = getKinMoves(move, piece, game, dir)
	case NariKyou:
		possibleMoves = getKinMoves(move, piece, game, dir)
	case NariKei:
		possibleMoves = getKinMoves(move, piece, game, dir)
	case NariGin:
		possibleMoves = getKinMoves(move, piece, game, dir)
	case Uma:
		possibleMoves = getUmaMoves(move, piece, game)
	case Ryuu:
		possibleMoves = getRyuuMoves(move, piece, game)
	case Checker:
		possibleMoves = getCheckerMoves(move, piece, game, dir)
	case CheckerKing:
		possibleMoves = getCheckerKingMoves(move, piece, game)
	}

	return possibleMoves
}

func getPawnMoves(move Move, piece Piece, game Game, direction int) []Vec2 {
	var validMovePositions []Vec2
	//move forward
	newY := move.Start.Y - 1*direction
	newPos := Vec2{X: move.Start.X, Y: newY}
	newPos2 := Vec2{X: move.Start.X, Y: newY - direction}
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
	relativeMovePos := []Vec2{{X: -1, Y: -1}, {X: 1, Y: -1}}
	for i := 0; i < len(relativeMovePos); i++ {
		newPos = relativeMovePos[i]
		newPos.Y *= direction

		newPos.X += move.Start.X
		newPos.Y += move.Start.Y
		if checkPositionInbounds(newPos, game) {
			space := game.Board.Board[newPos.Y][newPos.X]
			if space != nil && space.Owner == getEnemyTurnInt(game) {
				validMovePositions = append(validMovePositions, newPos)
			} else if space == nil && newPos == *game.EnPassant {
				validMovePositions = append(validMovePositions, newPos)
			}
		}
	}

	return validMovePositions
}

func getKnightMoves(move Move, piece Piece, game Game) []Vec2 {
	var validMovePositions []Vec2

	relativeMovePos := []Vec2{
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
		newPos.X += move.Start.X
		newPos.Y += move.Start.Y

		if checkPositionInbounds(newPos, game) {
			space := game.Board.Board[newPos.Y][newPos.X]
			if space == nil || space.Owner != piece.Owner {
				validMovePositions = append(validMovePositions, newPos)
			}
		}
	}

	return validMovePositions
}

func getBishopMoves(move Move, piece Piece, game Game) []Vec2 {
	var validMovePositions []Vec2

	directions := []Vec2{
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
			newPos := move.Start

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

func getRookMoves(move Move, piece Piece, game Game) []Vec2 {
	var validMovePositions []Vec2

	directions := []Vec2{
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
			newPos := move.Start

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

func getQueenMoves(move Move, piece Piece, game Game) []Vec2 {
	var validMovePositions []Vec2

	bishopMoves := getBishopMoves(move, piece, game)
	rookMoves := getRookMoves(move, piece, game)
	validMovePositions = append(bishopMoves, rookMoves...)

	return validMovePositions
}

func getKingMoves(move Move, piece Piece, game Game) []Vec2 {
	var validMovePositions []Vec2

	directions := []Vec2{
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

		newPos := move.Start
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

func getFuMoves(move Move, piece Piece, game Game, direction int) []Vec2 {
	var validMovePositions []Vec2

	newPos := move.Start
	newPos.Y += -1 * direction

	if checkPositionInbounds(newPos, game) {
		space := game.Board.Board[newPos.Y][newPos.X]
		if space == nil || space.Owner != piece.Owner {
			validMovePositions = append(validMovePositions, newPos)
		}
	}

	return validMovePositions
}

func getKyouMoves(move Move, piece Piece, game Game, dir int) []Vec2 {
	var validMovePositions []Vec2

	i := 0
	for i >= 0 {
		i++
		newPos := move.Start

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

func getKeiMoves(move Move, piece Piece, game Game, direction int) []Vec2 {
	var validMovePositions []Vec2

	directions := []Vec2{
		{X: -1, Y: -2 * direction},
		{X: 1, Y: -2 * direction},
	}

	for i := 0; i < len(directions); i++ {
		dir := directions[i]

		newPos := move.Start
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

func getGinMoves(move Move, piece Piece, game Game, dir int) []Vec2 {
	var validMovePositions []Vec2

	directions := []Vec2{
		{X: -1, Y: -1 * dir},
		{X: 0, Y: -1 * dir},
		{X: 1, Y: -1 * dir},
		{X: -1, Y: 1 * dir},
		{X: 1, Y: 1 * dir},
	}

	for i := 0; i < len(directions); i++ {
		dir := directions[i]

		newPos := move.Start
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

func getKinMoves(move Move, piece Piece, game Game, dir int) []Vec2 {
	var validMovePositions []Vec2

	directions := []Vec2{
		{X: -1, Y: -1 * dir},
		{X: 0, Y: -1 * dir},
		{X: 1, Y: -1 * dir},
		{X: -1, Y: 0},
		{X: 1, Y: 0},
		{X: 0, Y: 1 * dir},
	}

	for i := 0; i < len(directions); i++ {
		dir := directions[i]

		newPos := move.Start
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

func getUmaMoves(move Move, piece Piece, game Game) []Vec2 {
	var validMovePositions []Vec2

	directions := []Vec2{
		{X: -1, Y: 0},
		{X: 1, Y: 0},
		{X: 0, Y: -1},
		{X: 0, Y: 1},
	}

	for i := 0; i < len(directions); i++ {
		dir := directions[i]

		newPos := move.Start
		newPos.X += dir.X
		newPos.Y += dir.Y

		if checkPositionInbounds(newPos, game) {
			space := game.Board.Board[newPos.Y][newPos.X]
			if space == nil || space.Owner != piece.Owner {
				validMovePositions = append(validMovePositions, newPos)
			}
		}
	}

	bishopMoves := getBishopMoves(move, piece, game)
	validMovePositions = append(validMovePositions, bishopMoves...)
	return validMovePositions
}

func getRyuuMoves(move Move, piece Piece, game Game) []Vec2 {
	var validMovePositions []Vec2

	directions := []Vec2{
		{X: -1, Y: -1},
		{X: 1, Y: -1},
		{X: -1, Y: 1},
		{X: 1, Y: 1},
	}

	for i := 0; i < len(directions); i++ {
		dir := directions[i]

		newPos := move.Start
		newPos.X += dir.X
		newPos.Y += dir.Y

		if checkPositionInbounds(newPos, game) {
			space := game.Board.Board[newPos.Y][newPos.X]
			if space == nil || space.Owner != piece.Owner {
				validMovePositions = append(validMovePositions, newPos)
			}
		}
	}

	rookMoves := getRookMoves(move, piece, game)
	validMovePositions = append(validMovePositions, rookMoves...)
	return validMovePositions
}

func getCheckerMoves(move Move, piece Piece, game Game, dir int) []Vec2 {
	var validMovePositions []Vec2

	directions := []Vec2{
		{X: -1, Y: -1 * dir},
		{X: 1, Y: -1 * dir},
	}

	for i := 0; i < len(directions); i++ {
		dir := directions[i]

		jumpPos := move.Start
		landPos := move.Start

		jumpPos.X += dir.X
		jumpPos.Y += dir.Y

		landPos.X += dir.X * 2
		landPos.Y += dir.Y * 2

		if !checkPositionInbounds(jumpPos, game) || !checkPositionInbounds(landPos, game) {
			break
		}

		landSpace := game.Board.Board[landPos.Y][landPos.X]
		jumpSpace := game.Board.Board[jumpPos.Y][jumpPos.X]

		if landSpace == nil && (jumpSpace != nil && jumpSpace.Owner != piece.Owner) {
			validMovePositions = append(validMovePositions, landPos)
		}
	}

	return validMovePositions
}

func getCheckerKingMoves(move Move, piece Piece, game Game) []Vec2 {
	var validMovePositions []Vec2

	directions := []Vec2{
		{X: -1, Y: -1},
		{X: 1, Y: -1},
		{X: -1, Y: 1},
		{X: 1, Y: 1},
	}

	for i := 0; i < len(directions); i++ {
		dir := directions[i]

		jumpPos := move.Start
		landPos := move.Start

		jumpPos.X += dir.X
		jumpPos.Y += dir.Y

		landPos.X += dir.X * 2
		landPos.Y += dir.Y * 2

		if !checkPositionInbounds(jumpPos, game) || !checkPositionInbounds(landPos, game) {
			break
		}

		landSpace := game.Board.Board[landPos.Y][landPos.X]
		jumpSpace := game.Board.Board[jumpPos.Y][jumpPos.X]

		if landSpace == nil && (jumpSpace != nil && jumpSpace.Owner != piece.Owner) {
			validMovePositions = append(validMovePositions, landPos)
		}
	}

	return validMovePositions
}

func checkEndPosInPossibleMoves(possibleMoves []Vec2, move Move) error {
	for i := 0; i < len(possibleMoves); i++ {
		possibleMove := possibleMoves[i]

		if move.End == possibleMove {
			return nil
		}
	}

	return fmt.Errorf("No valid moves")
}

func getMoveDirection(game Game) int {
	direction := 1
	if game.Turn == 1 {
		direction = -1
	}
	return direction
}

func getCastleMoves(move Move, piece Piece, game Game) []Vec2 {
	var validMovePositions []Vec2

	//left
	for i := move.Start.X - 1; i >= 0; i-- {
		targetPiece := game.Board.Board[move.Start.Y][i]
		if targetPiece != nil {
			if targetPiece.Type == Rook && targetPiece.Owner == piece.Owner {
				validMovePositions = append(validMovePositions, Vec2{X: i, Y: move.Start.Y})
			}
			break
		}
	}

	for i := move.Start.X + 1; i < game.Board.Width; i++ {
		targetPiece := game.Board.Board[move.Start.Y][i]
		if targetPiece != nil {
			if targetPiece.Type == Rook && targetPiece.Owner == piece.Owner {
				validMovePositions = append(validMovePositions, Vec2{X: i, Y: move.Start.Y})
			}
			break
		}
	}

	return validMovePositions
}
