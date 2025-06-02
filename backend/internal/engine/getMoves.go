package engine

func getPawnMoves(move Move, piece Piece, game Game, direction int) []Vec2 {
	var validMovePositions []Vec2

	//en passant

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

//NEED TO DO RECUSIVELY
// func getCheckerMoves(move Move, piece Piece, game Game, dir int) []Vec2 {
// 	var validMovePositions []Vec2
//
// 	directions := []Vec2{
// 		{X: -1, Y: -1 * dir},
// 		{X: 1, Y: -1 * dir},
// 	}
//
// 	for i := 0; i < len(directions); i++ {
// 		dir := directions[i]
// 		newPos := move.Start
//
// 		newPos.X += dir.X
// 		newPos.Y += dir.Y
//
// 		if !checkPositionInbounds(newPos, game) {
// 			break
// 		}
//
// 		space := game.Board.Board[newPos.Y][newPos.X]
// 		if space == nil || space.Owner != piece.Owner {
// 			validMovePositions = append(validMovePositions, newPos)
// 		}
// 	}
//
// 	return validMovePositions
// }

// func getCheckerKingMoves(move Move, piece Piece, game Game) []Vec2 {
// 	var validMovePositions []Vec2
//
// 	directions := []Vec2{
// 		{X: -1, Y: -1 * dir},
// 		{X: 1, Y: -1 * dir},
// 	}
//
// 	for i := 0; i < len(directions); i++ {
// 		dir := directions[i]
// 		newPos := move.Start
//
// 		newPos.X += dir.X
// 		newPos.Y += dir.Y
//
// 		if !checkPositionInbounds(newPos, game) {
// 			break
// 		}
//
// 		space := game.Board.Board[newPos.Y][newPos.X]
// 		if space == nil || space.Owner != piece.Owner {
// 			validMovePositions = append(validMovePositions, newPos)
// 		}
// 	}
//
// 	return validMovePositions
// }
