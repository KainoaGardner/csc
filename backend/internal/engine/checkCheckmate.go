package engine

import (
	"github.com/KainoaGardner/csc/internal/types"
)

func GetInCheck(game Game) bool {
	kings := getKingPos(game)
	for _, king := range kings {
		if checkUnderAttack(king, game) {
			return true
		}
	}

	return false
}

func getKingPos(game Game) []types.Vec2 {
	var result []types.Vec2

	for i := 0; i < game.Board.Height; i++ {
		for j := 0; j < game.Board.Width; j++ {
			space := game.Board.Board[i][j]

			if space != nil && space.Owner == game.Turn && (space.Type == King || space.Type == Ou) {
				result = append(result, types.Vec2{X: j, Y: i})
			}
		}
	}

	return result
}

func checkUnderAttack(pos types.Vec2, game Game) bool {
	attackSpace := map[types.Vec2]bool{}

	game.Turn = getEnemyTurnInt(game)
	for i := 0; i < game.Board.Height; i++ {
		for j := 0; j < game.Board.Width; j++ {
			space := game.Board.Board[i][j]
			if space != nil && space.Owner == game.Turn {
				pos := types.Vec2{X: j, Y: i}
				dir := getMoveDirection(game)
				possibleMoves := getPieceMoves(pos, *space, game, dir)
				for _, move := range possibleMoves {
					attackSpace[move] = true
				}
			}
		}
	}

	_, ok := attackSpace[pos]
	return ok
}

func GetInCheckmate(game Game) bool {
	if !GetInCheck(game) {
		return false
	}

	//check normal moves and promotions
	for i := 0; i < game.Board.Height; i++ {
		for j := 0; j < game.Board.Width; j++ {
			space := game.Board.Board[i][j]
			if space != nil && space.Owner == game.Turn {

			}
		}
	}

	//check all drops

	return false
}

func checkMovePieceOutOfCheck(pos types.Vec2, piece Piece, game Game) bool {
	dir := getMoveDirection(game)

	if piece.Type >= Pawn && piece.Type <= Ryuu {
		possibleMoves := []types.Vec2{}
		switch piece.Type {
		case Pawn:
			possibleMoves = getPawnMoves(pos, piece, game, dir)
		case Knight:
			possibleMoves = getKnightMoves(pos, piece, game)
		case Bishop:
			possibleMoves = getBishopMoves(pos, piece, game)
		case Rook:
			possibleMoves = getRookMoves(pos, piece, game)
		case Queen:
			possibleMoves = getQueenMoves(pos, piece, game)
		case King:
			possibleMoves = getKingMoves(pos, piece, game)
		case Fu:
			possibleMoves = getFuMoves(pos, piece, game, dir)
		case Kyou:
			possibleMoves = getKyouMoves(pos, piece, game, dir)
		case Kei:
			possibleMoves = getKeiMoves(pos, piece, game, dir)
		case Gin:
			possibleMoves = getGinMoves(pos, piece, game, dir)
		case Kin:
			possibleMoves = getKinMoves(pos, piece, game, dir)
		case Kaku:
			possibleMoves = getBishopMoves(pos, piece, game)
		case Hi:
			possibleMoves = getRookMoves(pos, piece, game)
		case Ou:
			possibleMoves = getKingMoves(pos, piece, game)
		case To:
			possibleMoves = getKinMoves(pos, piece, game, dir)
		case NariKyou:
			possibleMoves = getKinMoves(pos, piece, game, dir)
		case NariKei:
			possibleMoves = getKinMoves(pos, piece, game, dir)
		case NariGin:
			possibleMoves = getKinMoves(pos, piece, game, dir)
		case Uma:
			possibleMoves = getUmaMoves(pos, piece, game)
		case Ryuu:
			possibleMoves = getRyuuMoves(pos, piece, game)
		case Checker:
		}

		for _, move := range possibleMoves {
			if chessShogiOutOfCheck(piece, pos, move, game) {
				return true
			}
		}

	} else if piece.Type >= Checker && piece.Type <= CheckerKing {

	}

	return false
}

func chessShogiOutOfCheck(piece Piece, startPos types.Vec2, endPos types.Vec2, game Game) bool {
	game.Board.Board[startPos.Y][startPos.X] = nil
	game.Board.Board[endPos.Y][endPos.X] = &piece
	if !GetInCheck(game) {
		return true
	}
	return false
}
