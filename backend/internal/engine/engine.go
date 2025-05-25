package engine

import ()

//in bounds
//moving your piece
//correct piece movement
//if promotiion check correct promotion
//if drop check drop
//if in check cant move unless not in check after

//

func CheckValidMove(game Game, move Move) bool {
	if !checkMoveInBounds(game, move) {
		return false
	}

	return true
}

func checkMoveInBounds(game Game, move Move) bool {

	return true
}
