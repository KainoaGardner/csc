package engine

import "fmt"

// in bounds
// moving your piece
// correct piece movement
// if promotiion check correct promotion
// if drop check drop
// if in check cant move unless not in check after

func CheckValidMove(moves *[]Move, game Game) error {
	moveCount := len(*moves)
	if moveCount == 1 {
		move := (*moves)[0]
		err := checkMoveInBounds(move, game)
		if err != nil {
			return err
		}

		piece, err := getPiece(move, game)
		if err != nil {
			return err
		}

		err = checkValidPieceMoves(move, piece, game)
		if err != nil {
			return err
		}
	}

	return nil
}

func checkMoveInBounds(move Move, game Game) error {
	if move.Start.X < 0 || move.Start.X >= game.Board.Width {
		return fmt.Errorf("Start x out of board bounds")
	}
	if move.Start.Y < 0 || move.Start.Y >= game.Board.Height {
		return fmt.Errorf("Start y out of board bounds")
	}
	if move.End.X < 0 || move.End.X >= game.Board.Width {
		return fmt.Errorf("End x out of board bounds")
	}
	if move.End.Y < 0 || move.End.Y >= game.Board.Height {
		return fmt.Errorf("End y out of board bounds")
	}

	return nil
}

func getPiece(move Move, game Game) (Piece, error) {
	piece := game.Board.Board[move.Start.Y][move.Start.X]
	if piece == nil {
		return *piece, fmt.Errorf("Cant move empty piece")
	}

	if piece.Owner != game.Turn {
		return *piece, fmt.Errorf("Cant move other players piece")
	}

	return *piece, nil
}

func checkValidPieceMoves(move Move, piece Piece, game Game) error {
	//check valid drop

	//check valid normal move
	//in normal moves if move has promotion check that

	var result error
	dir := getMoveDirection(game)

	possibleMoves := []Vec2{}

	if move.Drop != nil {
		//drop check
	} else {
		switch piece.Type {
		case Pawn: //add en passant
			possibleMoves = getPawnMoves(move, piece, game, dir)
			result = checkEndPosInPossibleMoves(possibleMoves, move)
		case Knight:
			possibleMoves = getKnightMoves(move, piece, game)
			result = checkEndPosInPossibleMoves(possibleMoves, move)
		case Bishop:
			possibleMoves = getBishopMoves(move, piece, game)
			result = checkEndPosInPossibleMoves(possibleMoves, move)
		case Rook:
			possibleMoves = getRookMoves(move, piece, game)
			result = checkEndPosInPossibleMoves(possibleMoves, move)
		case Queen:
			possibleMoves = getQueenMoves(move, piece, game)
			result = checkEndPosInPossibleMoves(possibleMoves, move)
		case King:
			possibleMoves = getKingMoves(move, piece, game)
			result = checkEndPosInPossibleMoves(possibleMoves, move)
		case Fu:
			possibleMoves = getFuMoves(move, piece, game, dir)
			result = checkEndPosInPossibleMoves(possibleMoves, move)
		case Kyou:
			possibleMoves = getKyouMoves(move, piece, game, dir)
			result = checkEndPosInPossibleMoves(possibleMoves, move)
		case Kei:
			possibleMoves = getKeiMoves(move, piece, game, dir)
			result = checkEndPosInPossibleMoves(possibleMoves, move)
		case Gin:
			possibleMoves = getGinMoves(move, piece, game, dir)
			result = checkEndPosInPossibleMoves(possibleMoves, move)
		case Kin:
			possibleMoves = getKinMoves(move, piece, game, dir)
			result = checkEndPosInPossibleMoves(possibleMoves, move)
		case Kaku:
			possibleMoves = getBishopMoves(move, piece, game)
			result = checkEndPosInPossibleMoves(possibleMoves, move)
		case Hi:
			possibleMoves = getRookMoves(move, piece, game)
			result = checkEndPosInPossibleMoves(possibleMoves, move)
		case Ou:
			possibleMoves = getKingMoves(move, piece, game)
			result = checkEndPosInPossibleMoves(possibleMoves, move)
		case To:
			possibleMoves = getKinMoves(move, piece, game, dir)
			result = checkEndPosInPossibleMoves(possibleMoves, move)
		case NariKyou:
			possibleMoves = getKinMoves(move, piece, game, dir)
			result = checkEndPosInPossibleMoves(possibleMoves, move)
		case NariKei:
			possibleMoves = getKinMoves(move, piece, game, dir)
			result = checkEndPosInPossibleMoves(possibleMoves, move)
		case NariGin:
			possibleMoves = getKinMoves(move, piece, game, dir)
			result = checkEndPosInPossibleMoves(possibleMoves, move)
		case Uma:
			possibleMoves = getUmaMoves(move, piece, game)
			result = checkEndPosInPossibleMoves(possibleMoves, move)
		case Ryuu:
			possibleMoves = getRyuuMoves(move, piece, game)
			result = checkEndPosInPossibleMoves(possibleMoves, move)
		case Checker:
		}

		if move.Promote != nil {
			//promote check
		}
	}

	for i := 0; i < len(possibleMoves); i++ {
		possibleMove := possibleMoves[i]
		fmt.Println(possibleMove)
	}

	return result
}

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

// func getUmaMoves(move Move, game Game) []Vec2 {
// 	var validMovePositions []Vec2
// 	bishopMoves := getBishopMoves(move, game)
//
// 	return validMovePositions
// }

// const (
// 	Empty = iota
// 	Pawn
// 	Knight
// 	Bishop
// 	Rook
// 	Queen
// 	King
// 	Fu
// 	Kyou
// 	Kei
// 	Gin
// 	Kin
// 	Kaku
// 	Hi
// 	Ou
// 	To
// 	NariKyou
// 	NariKei
// 	NariGin
// 	Uma
// 	Ryuu
// 	Checker
// )

func checkPositionInbounds(pos Vec2, game Game) bool {
	if pos.X < 0 || pos.X >= game.Board.Width {
		return false
	}
	if pos.Y < 0 || pos.Y >= game.Board.Height {
		return false
	}

	return true
}

func getEnemyTurnInt(game Game) int {
	if game.Turn == 0 {
		return 1
	} else {
		return 0
	}
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
