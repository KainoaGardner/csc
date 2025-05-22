package engine

import (
	"fmt"
)

func Main() {
	var board Board
	board.width = 8
	board.height = 8

	fmt.Print(board.width)

}
