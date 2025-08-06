import { type Vec2, PieceEnum, checkVec2Equal } from "./util.ts"
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
  let possibleMoves: Vec2[] = []

  switch (piece.type) {
    case PieceEnum.Pawn: {
      possibleMoves = getPawnMoves(start, piece, game, dir)
      break
    }
    case PieceEnum.Knight: {
      possibleMoves = getKnightMoves(start, piece, game)
      break
    }
    case PieceEnum.Bishop: {
      possibleMoves = getBishopMoves(start, piece, game)
      break
    }
    case PieceEnum.Rook: {
      possibleMoves = getRookMoves(start, piece, game)
      break
    }
    case PieceEnum.Queen: {
      possibleMoves = getQueenMoves(start, piece, game)
      break
    }
    case PieceEnum.King: {
      possibleMoves = getKingMoves(start, piece, game)
      possibleMoves.push(...getCastleMoves(start, piece, game))
      break
    }
    case PieceEnum.Fu: {
      possibleMoves = getFuMoves(start, piece, game, dir)
      break
    }
    case PieceEnum.Kyou: {
      possibleMoves = getKyouMoves(start, piece, game, dir)
      break
    }
    case PieceEnum.Kei: {
      possibleMoves = getKeiMoves(start, piece, game, dir)
      break
    }
    case PieceEnum.Gin: {
      possibleMoves = getGinMoves(start, piece, game, dir)
      break
    }
    case PieceEnum.Kin: {
      possibleMoves = getKinMoves(start, piece, game, dir)
      break
    }
    case PieceEnum.Kaku: {
      possibleMoves = getBishopMoves(start, piece, game)
      break
    }
    case PieceEnum.Hi: {
      possibleMoves = getRookMoves(start, piece, game)
      break
    }
    case PieceEnum.Ou: {
      possibleMoves = getKingMoves(start, piece, game)
      break
    }
    case PieceEnum.To: {
      possibleMoves = getKinMoves(start, piece, game, dir)
      break
    }
    case PieceEnum.NariKyou: {
      possibleMoves = getKinMoves(start, piece, game, dir)
      break
    }
    case PieceEnum.NariKei: {
      possibleMoves = getKinMoves(start, piece, game, dir)
      break
    }
    case PieceEnum.NariGin: {
      possibleMoves = getKinMoves(start, piece, game, dir)
      break
    }
    case PieceEnum.Uma: {
      possibleMoves = getUmaMoves(start, piece, game)
      break
    }
    case PieceEnum.Ryuu: {
      possibleMoves = getRyuuMoves(start, piece, game)
      break
    }
    case PieceEnum.Checker: {
      possibleMoves = getCheckerMoves(start, piece, game, dir)
      break
    }
    case PieceEnum.CheckerKing: {
      possibleMoves = getCheckerKingMoves(start, piece, game)
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

function checkPositionInbounds(pos: Vec2, game: Game): boolean {
  if (pos.x < 0 || pos.x >= game.width) {
    return false
  }
  if (pos.y < 0 || pos.y >= game.height) {
    return false
  }

  return true
}

function getPawnMoves(pos: Vec2, piece: Piece, game: Game, direction: number): Vec2[] {
  const result: Vec2[] = []

  //move forward
  const newY = pos.y - 1 * direction
  const newPos = { x: pos.x, y: newY }
  const newPos2 = { x: pos.x, y: newY - direction }

  if (checkPositionInbounds(newPos, game)) {
    let space = game.board[newPos.y][newPos.x]

    if (space === null) {
      result.push(newPos)

      //check starting move 2 space
      if (checkPositionInbounds(newPos2, game)) {
        space = game.board[newPos2.y][newPos2.x]
        if (space === null && !piece.moved) {
          result.push(newPos2)
        }
      }
    }
  }

  //capture squares
  const relativeMovePos = [{ x: -1, y: -1 }, { x: 1, y: -1 }]
  for (let i = 0; i < relativeMovePos.length; i++) {
    const newPos = relativeMovePos[i]
    newPos.y *= direction

    newPos.x += pos.x
    newPos.y += pos.y
    if (checkPositionInbounds(newPos, game)) {
      const space = game.board[newPos.y][newPos.x]
      if (space !== null && space.owner == piece.owner) {
        result.push(newPos)
      } else if (space === null && game.enPassant !== null && checkVec2Equal(newPos, game.enPassant)) {
        result.push(newPos)
      }
    }
  }

  return result
}

function getKnightMoves(pos: Vec2, piece: Piece, game: Game): Vec2[] {
  const result: Vec2[] = []

  const relativeMovePos: Vec2[] = [
    { x: -1, y: -2 },
    { x: 1, y: -2 },
    { x: -2, y: -1 },
    { x: 2, y: -1 },
    { x: -2, y: 1 },
    { x: 2, y: 1 },
    { x: -1, y: 2 },
    { x: 1, y: 2 },
  ]


  for (let i = 0; i < relativeMovePos.length; i++) {
    const newPos = relativeMovePos[i]
    newPos.x += pos.x
    newPos.y += pos.y

    if (checkPositionInbounds(newPos, game)) {
      const space = game.board[newPos.y][newPos.x]
      if (space === null || space.owner !== piece.owner) {
        result.push(newPos)
      }
    }
  }

  return result
}

function getBishopMoves(pos: Vec2, piece: Piece, game: Game): Vec2[] {
  const result: Vec2[] = []

  const directions: Vec2[] = [
    { x: -1, y: -1 },
    { x: -1, y: 1 },
    { x: 1, y: -1 },
    { x: 1, y: 1 },
  ]

  for (let i = 0; i < directions.length; i++) {
    const dir = directions[i]

    let j = 0
    while (j >= 0) {
      j++
      const newPos = pos

      newPos.x += dir.x * j
      newPos.y += dir.y * j

      if (!checkPositionInbounds(newPos, game)) {
        break
      }

      const space = game.board[newPos.y][newPos.x]
      if (space === null) {
        result.push(newPos)
      } else if (space.owner != piece.owner) {
        result.push(newPos)
        break
      } else {
        break
      }
    }
  }

  return result
}

function getRookMoves(pos: Vec2, piece: Piece, game: Game): Vec2[] {
  const result: Vec2[] = []

  const directions: Vec2[] = [
    { x: -1, y: 0 },
    { x: 1, y: 0 },
    { x: 0, y: -1 },
    { x: 0, y: 1 },
  ]

  for (let i = 0; i < directions.length; i++) {
    const dir = directions[i]

    let j = 0
    while (j >= 0) {
      j++
      const newPos = pos

      newPos.x += dir.x * j
      newPos.y += dir.y * j

      if (!checkPositionInbounds(newPos, game)) {
        break
      }

      const space = game.board[newPos.y][newPos.x]
      if (space === null) {
        result.push(newPos)
      } else if (space.owner !== piece.owner) {
        result.push(newPos)
        break
      } else {
        break
      }
    }
  }
  return result
}

function getQueenMoves(pos: Vec2, piece: Piece, game: Game): Vec2[] {
  const result: Vec2[] = []

  const bishopMoves = getBishopMoves(pos, piece, game)
  const rookMoves = getRookMoves(pos, piece, game)

  result.push(...bishopMoves)
  result.push(...rookMoves)

  return result
}

function getKingMoves(pos: Vec2, piece: Piece, game: Game): Vec2[] {
  const result: Vec2[] = []

  const directions: Vec2[] = [
    { x: -1, y: -1 },
    { x: 0, y: -1 },
    { x: 1, y: -1 },
    { x: -1, y: 0 },
    { x: 1, y: 0 },
    { x: -1, y: 1 },
    { x: 0, y: 1 },
    { x: 1, y: 1 },
  ]

  for (let i = 0; i < directions.length; i++) {
    const dir = directions[i]

    const newPos = pos
    newPos.x += dir.x
    newPos.y += dir.y

    if (checkPositionInbounds(newPos, game)) {
      const space = game.board[newPos.y][newPos.x]
      if (space === null || space.owner != piece.owner) {
        result.push(newPos)
      }
    }
  }

  return result
}

function getCastleMoves(pos: Vec2, piece: Piece, game: Game): Vec2[] {
  const result: Vec2[] = []

  //left
  for (let i = pos.x - 1; i >= 0; i--) {
    const targetPiece = game.board[pos.y][i]
    if (targetPiece !== null) {
      if (targetPiece.type === PieceEnum.Rook && targetPiece.owner == piece.owner) {
        result.push({ x: i, y: pos.y })
      }
      break
    }
  }

  for (let i = pos.x + 1; i < game.width; i++) {
    const targetPiece = game.board[pos.y][i]
    if (targetPiece !== null) {
      if (targetPiece.type === PieceEnum.Rook && targetPiece.owner == piece.owner) {
        result.push({ x: i, y: pos.y })
      }
      break
    }
  }

  return result
}

function getFuMoves(pos: Vec2, piece: Piece, game: Game, direction: number): Vec2[] {
  const result: Vec2[] = []

  const newPos = pos
  newPos.y += -1 * direction

  if (checkPositionInbounds(newPos, game)) {
    const space = game.board[newPos.y][newPos.x]
    if (space === null || space.owner != piece.owner) {
      result.push(newPos)
    }
  }

  return result
}

function getKyouMoves(pos: Vec2, piece: Piece, game: Game, dir: number): Vec2[] {
  const result: Vec2[] = []

  let i = 0
  while (i >= 0) {
    i++
    const newPos = pos

    newPos.y += -i * dir

    if (!checkPositionInbounds(newPos, game)) {
      break
    }

    const space = game.board[newPos.y][newPos.x]
    if (space === null) {
      result.push(newPos)
    } else if (space.owner !== piece.owner) {
      result.push(newPos)
      break
    } else {
      break
    }
  }

  return result
}

function getKeiMoves(pos: Vec2, piece: Piece, game: Game, direction: number): Vec2[] {
  const result: Vec2[] = []

  const directions: Vec2[] = [
    { x: -1, y: -2 * direction },
    { x: 1, y: -2 * direction },
  ]

  for (let i = 0; i < directions.length; i++) {
    const dir = directions[i]

    const newPos = pos
    newPos.x += dir.x
    newPos.y += dir.y

    if (checkPositionInbounds(newPos, game)) {
      const space = game.board[newPos.y][newPos.x]
      if (space === null || space.owner != piece.owner) {
        result.push(newPos)
      }
    }
  }

  return result
}

function getGinMoves(pos: Vec2, piece: Piece, game: Game, dir: number): Vec2[] {
  const result: Vec2[] = []

  const directions: Vec2[] = [
    { x: -1, y: -1 * dir },
    { x: 0, y: -1 * dir },
    { x: 1, y: -1 * dir },
    { x: -1, y: 1 * dir },
    { x: 1, y: 1 * dir },
  ]

  for (let i = 0; i < directions.length; i++) {
    const dir = directions[i]

    const newPos = pos
    newPos.x += dir.x
    newPos.y += dir.y

    if (checkPositionInbounds(newPos, game)) {
      const space = game.board[newPos.y][newPos.x]
      if (space === null || space.owner != piece.owner) {
        result.push(newPos)
      }
    }

  }
  return result
}

function getKinMoves(pos: Vec2, piece: Piece, game: Game, dir: number): Vec2[] {
  const result: Vec2[] = []

  const directions: Vec2[] = [
    { x: -1, y: -1 * dir },
    { x: 0, y: -1 * dir },
    { x: 1, y: -1 * dir },
    { x: -1, y: 0 },
    { x: 1, y: 0 },
    { x: 0, y: 1 * dir },
  ]

  for (let i = 0; i < directions.length; i++) {
    const dir = directions[i]

    const newPos = pos
    newPos.x += dir.x
    newPos.y += dir.y

    if (checkPositionInbounds(newPos, game)) {
      const space = game.board[newPos.y][newPos.x]
      if (space === null || space.owner != piece.owner) {
        result.push(newPos)
      }
    }

  }
  return result
}

function getUmaMoves(pos: Vec2, piece: Piece, game: Game): Vec2[] {
  const result: Vec2[] = []

  const directions: Vec2[] = [
    { x: -1, y: 0 },
    { x: 1, y: 0 },
    { x: 0, y: -1 },
    { x: 0, y: 1 },
  ]

  for (let i = 0; i < directions.length; i++) {
    const dir = directions[i]

    const newPos = pos
    newPos.x += dir.x
    newPos.y += dir.y

    if (checkPositionInbounds(newPos, game)) {
      const space = game.board[newPos.y][newPos.x]
      if (space === null || space.owner != piece.owner) {
        result.push(newPos)
      }
    }
  }

  const bishopMoves = getBishopMoves(pos, piece, game)
  result.push(...bishopMoves)
  return result
}

function getRyuuMoves(pos: Vec2, piece: Piece, game: Game): Vec2[] {
  const result: Vec2[] = []

  const directions: Vec2[] = [
    { x: -1, y: -1 },
    { x: 1, y: -1 },
    { x: -1, y: 1 },
    { x: 1, y: 1 },
  ]

  for (let i = 0; i < directions.length; i++) {
    const dir = directions[i]

    const newPos = pos
    newPos.x += dir.x
    newPos.y += dir.y

    if (checkPositionInbounds(newPos, game)) {
      const space = game.board[newPos.y][newPos.x]
      if (space === null || space.owner != piece.owner) {
        result.push(newPos)
      }
    }
  }

  const rookMoves = getRookMoves(pos, piece, game)
  result.push(...rookMoves)
  return result
}

function getCheckerMoves(pos: Vec2, piece: Piece, game: Game, dir: number): Vec2[] {
  const result: Vec2[] = []

  const directions: Vec2[] = [
    { x: -1, y: -1 * dir },
    { x: 1, y: -1 * dir },
  ]

  for (let i = 0; i < directions.length; i++) {
    const dir = directions[i]

    const jumpPos = pos
    const landPos = pos

    jumpPos.x += dir.x
    jumpPos.y += dir.y

    landPos.x += dir.x * 2
    landPos.y += dir.y * 2

    if (checkPositionInbounds(jumpPos, game)) {
      const jumpSpace = game.board[jumpPos.y][jumpPos.x]
      if (jumpSpace === null) {
        result.push(jumpPos)
      }

      if (checkPositionInbounds(landPos, game)) {
        const landSpace = game.board[landPos.y][landPos.x]
        if (landSpace === null && (jumpSpace !== null && jumpSpace.owner != piece.owner)) {
          result.push(landPos)
        }

      }
    }
  }

  return result
}

function getCheckerKingMoves(pos: Vec2, piece: Piece, game: Game): Vec2[] {
  const result: Vec2[] = []

  const directions: Vec2[] = [
    { x: -1, y: -1 },
    { x: 1, y: -1 },
    { x: -1, y: 1 },
    { x: 1, y: 1 },
  ]

  for (let i = 0; i < directions.length; i++) {
    const dir = directions[i]
    const jumpPos = pos
    const landPos = pos

    jumpPos.x += dir.x
    jumpPos.y += dir.y

    landPos.x += dir.x * 2
    landPos.y += dir.y * 2

    if (checkPositionInbounds(jumpPos, game)) {
      const jumpSpace = game.board[jumpPos.y][jumpPos.x]
      if (jumpSpace === null) {
        result.push(jumpPos)
      }

      if (checkPositionInbounds(landPos, game)) {
        const landSpace = game.board[landPos.y][landPos.x]
        if (landSpace === null && (jumpSpace !== null && jumpSpace.owner != piece.owner)) {
          result.push(landPos)
        }
      }
    }
  }

  return result
}
