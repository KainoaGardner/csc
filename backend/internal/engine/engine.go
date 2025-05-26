package engine

import "fmt"

//in bounds
//moving your piece
//correct piece movement
//if promotiion check correct promotion
//if drop check drop
//if in check cant move unless not in check after

// type Move struct {
// 	Start      [2]int
// 	End        [2]int
// 	StartPiece Piece
// 	EndPiece   Piece
// 	TakenPiece Piece
// 	Promote    bool
// 	Drop       bool
// }

func setupMove(move *Move, game Game) {
	if !move.Drop {
		move.StartPiece = game.Board.Board[move.Start[1]][move.Start[0]]
	}
	if !move.Promote {
		move.EndPiece = move.StartPiece
	}

	move.TakenPiece = game.Board.Board[move.End[1]][move.End[0]]
}

func checkMoveInBounds(move Move, game Game) error {
	if move.Start[0] < 0 || move.Start[0] >= game.Board.Width {
		return fmt.Errorf("Start x out of board bounds")
	}
	if move.Start[1] < 0 || move.Start[1] >= game.Board.Height {
		return fmt.Errorf("Start y out of board bounds")
	}
	if move.End[0] < 0 || move.End[0] >= game.Board.Width {
		return fmt.Errorf("End x out of board bounds")
	}
	if move.End[1] < 0 || move.End[1] >= game.Board.Height {
		return fmt.Errorf("End y out of board bounds")
	}

	return nil
}

func CheckValidMove(move *Move, game Game) error {
	err := checkMoveInBounds(*move, game)
	if err != nil {
		return fmt.Errorf("Move outside board bounds")
	}

	setupMove(move, game)

	err = checkMovablePiece(*move, game)
	if err != nil {
		return fmt.Errorf("Cannot move this piece")
	}

	err = checkValidPieceMoves(*move, game)
	if err != nil {
		return err
	}

	return nil
}

func checkMovablePiece(move Move, game Game) error {
	if move.StartPiece.Type == Empty {
		return fmt.Errorf("Cant move empty piece")
	}

	if move.StartPiece.Owner != game.Turn {
		return fmt.Errorf("Cant move other players piece")
	}

	return nil
}

func checkValidPieceMoves(move Move, game Game) error {
	//check valid drop

	//check valid normal move
	//in normal moves if move has promotion check that

	if move.Drop {
		//drop check
	} else {
		switch move.StartPiece.Type {
		case Pawn:
			checkPawnMove(move, game)
		case Knight:
		case Bishop:
		case Rook:
		case Queen:
		case King:
		case Fu:
		case Kyou:
		case Kei:
		case Gin:
		case Kin:
		case Kaku:
		case Hi:
		case Ou:
		case To:
		case NariKyou:
		case NariKei:
		case NariGin:
		case Uma:
		case Ryuu:
		case Checker:
		}
	}

	return nil
}

func checkPawnMove(move Move, game Game) {
	direction := 1
	if game.Turn == 1 {
		direction = -1
	}

	var validMovePositions []Vec2
	//en passant

	//first turn move
	if !move.StartPiece.Moved {
		newY := move.Start[1] - 2*direction
		newPos := Vec2{x: move.Start[0], y: newY}
		if checkPositionInbounds(newPos, game) {
			if game.Board.Board[newPos.y][newPos.x].Type == Empty {
				validMovePositions = append(validMovePositions, newPos)
			}
		}
	}

	//move forward
	newY := move.Start[1] - 1*direction
	newPos := Vec2{x: move.Start[0], y: newY}
	if checkPositionInbounds(newPos, game) {
		if game.Board.Board[newPos.y][newPos.x].Type == Empty {
			validMovePositions = append(validMovePositions, newPos)
		}
	}

	relativeMovePos := [2]Vec2{Vec2{x: -1, y: -1}, Vec2{x: 1, y: -1}}
	//capture squares
	for i := 0; i < len(relativeMovePos); i++ {
		newPos = relativeMovePos[i]
		newPos.y *= direction
		if checkPositionInbounds(newPos, game) {
			if game.Board.Board[newPos.y][newPos.x].Owner == getEnemyTurnInt(game) {
				validMovePositions = append(validMovePositions, newPos)
			}
		}
	}
}

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
	if pos.x < 0 || pos.x >= game.Board.Width {
		return false
	}
	if pos.y < 0 || pos.y >= game.Board.Height {
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
	}

	return fmt.Errorf("No valid moves")
}
