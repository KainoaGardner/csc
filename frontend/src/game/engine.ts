import { type Vec2, PieceEnum } from "./util.ts"
import { Piece } from "./piece.ts"
import { Game } from "./game.ts"

export function getMoveDirection(owner: number): number {
  if (owner === 0) { return 1 }
  else { return -1 }
}

export function checkValidPieceMove(start: Vec2, end: Vec2, piece: Piece, game: Game): boolean {
  const dir = getMoveDirection(game.userSide)
  const possibleMoves = getPieceMoves(start, piece, game, dir)
  const filteredMoves = filterPossibleMoves(start, possibleMoves, game)

  return checkEndPosInPossibleMoves(end, filteredMoves)
}

export function getPieceMoves(start: Vec2, piece: Piece, game: Game, dir: number): Vec2[] {
  const possibleMoves: Vec2[] = []

  switch (piece.type) {
    case PieceEnum.Pawn: {
      break
    }
    case PieceEnum.Knight: {
      break
    }
    case PieceEnum.Bishop: {
      break
    }
    case PieceEnum.Rook: {
      break
    }
    case PieceEnum.Queen: {
      break
    }
    case PieceEnum.King: {
      break
    }
    case PieceEnum.Fu: {
      break
    }
    case PieceEnum.Kyou: {
      break
    }
    case PieceEnum.Kei: {
      break
    }
    case PieceEnum.Gin: {
      break
    }
    case PieceEnum.Kin: {
      break
    }
    case PieceEnum.Kaku: {
      break
    }
    case PieceEnum.Hi: {
      break
    }
    case PieceEnum.Ou: {
      break
    }
    case PieceEnum.To: {
      break
    }
    case PieceEnum.NariKyou: {
      break
    }
    case PieceEnum.NariKei: {
      break
    }
    case PieceEnum.NariGin: {
      break
    }
    case PieceEnum.Uma: {
      break
    }
    case PieceEnum.Ryuu: {
      break
    }
    case PieceEnum.Checker: {
      break
    }
    case PieceEnum.CheckerKing: {
      break
    }
  }

  return possibleMoves
}

export function filterPossibleMoves(start: Vec2, possibleMoves: Vec2[], game: Game): Vec2[] {
  return []
}

export function checkEndPosInPossibleMoves(end: Vec2, possibleMoves: Vec2[]): boolean {
  return false
}
