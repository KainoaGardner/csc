import { type Vec2, PieceEnum, checkVec2Equal } from "./util.ts"
import { Piece } from "./piece.ts"
import { Game } from "./game.ts"

export function getMoveDirection(owner: number): number {
  if (owner === 0) { return 1 }
  else { return -1 }
}

export function getEnemyTurnInt(game: Game): number {
  if (game.turn === 0) {
    return 1
  } else {
    return 0
  }
}

export function checkValidPieceMove(start: Vec2, end: Vec2, piece: Piece, game: Game): boolean {
  if (!checkPieceOnBoard(start.x, start.y, game) || !checkPieceOnBoard(end.x, end.y, game)) {
    return false
  }

  const dir = getMoveDirection(game.turn)
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
  for (let i = possibleMoves.length - 1; i >= 0; i--) {
    const movePos = possibleMoves[i]
    const gameCopy: Game = copyGame(game)
    const piece = gameCopy.board[start.y][start.x]

    if (piece !== null && piece.type >= PieceEnum.Pawn && piece.type <= PieceEnum.Ryuu) {
      gameCopy.board[start.y][start.x] = null
      gameCopy.board[movePos.y][movePos.x] = piece
      if (getInCheck(gameCopy)) {
        possibleMoves.splice(i, 1)
      }
    } else if (piece !== null && piece.type >= PieceEnum.Checker && piece.type <= PieceEnum.CheckerKing) {
      if (checkerMovesInCheck(start, movePos, piece, gameCopy)) {
        possibleMoves.splice(i, 1)
      }
    }
  }

  return possibleMoves
}

export function checkEndPosInPossibleMoves(end: Vec2, possibleMoves: Vec2[]): boolean {
  for (let i = 0; i < possibleMoves.length; i++) {
    const possibleMove = possibleMoves[i]

    if (checkVec2Equal(end, possibleMove)) {
      return true
    }
  }

  return false
}

export function getPieceDrops(piece: Piece, game: Game): Vec2[] {
  const possibleDrops: Vec2[] = []

  for (let i = 0; i < game.height; i++) {
    for (let j = 0; j < game.width; j++) {
      const endPos: Vec2 = { x: j, y: i }
      const result = checkValidDropMove(endPos, piece, game)
      if (result) {
        possibleDrops.push(endPos)
      }
    }
  }

  return possibleDrops
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
    const newPos = { x: relativeMovePos[i].x, y: relativeMovePos[i].y }
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
    const newPos = { x: relativeMovePos[i].x, y: relativeMovePos[i].y }
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
      const newPos = { x: pos.x, y: pos.y }

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
      const newPos = { x: pos.x, y: pos.y }

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

    const newPos = { x: pos.x, y: pos.y }
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

  const newPos = { x: pos.x, y: pos.y }
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
    const newPos = { x: pos.x, y: pos.y }

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

    const newPos = { x: pos.x, y: pos.y }
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

    const newPos = { x: pos.x, y: pos.y }
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

    const newPos = { x: pos.x, y: pos.y }
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

    const newPos = { x: pos.x, y: pos.y }
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

    const newPos = { x: pos.x, y: pos.y }
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

    const jumpPos = { x: pos.x, y: pos.y }
    const landPos = { x: pos.x, y: pos.y }

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
    const jumpPos = { x: pos.x, y: pos.y }
    const landPos = { x: pos.x, y: pos.y }

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


export function copyGame(game: Game): Game {
  const gameCopy = new Game(game.id, game.width, game.height, game.placeLine, game.userSide, game.money, game.time)

  const boardCopy = Array.from({ length: gameCopy.height }, () => Array(gameCopy.width).fill(null))

  for (let i = 0; i < game.height; i++) {
    for (let j = 0; j < game.width; j++) {
      const piece = game.board[i][j]
      if (piece !== null) {
        const pieceCopy = new Piece(piece.x, piece.y, piece.type, piece.owner, piece.moved)
        boardCopy[i][j] = pieceCopy
      } else {
        boardCopy[i][j] = null
      }
    }
  }

  gameCopy.board = boardCopy
  gameCopy.turn = game.turn

  if (game.enPassant !== null) {
    const enPassantPos = { x: game.enPassant.x, y: game.enPassant.y }
    gameCopy.enPassant = enPassantPos
  } else {
    gameCopy.enPassant = null
  }

  if (game.checkerJump !== null) {
    const checkerJump = { x: game.checkerJump.x, y: game.checkerJump.y }
    gameCopy.checkerJump = checkerJump
  } else {
    gameCopy.checkerJump = null
  }

  return gameCopy
}

function getInCheckmate(game: Game): boolean {
  if (!getInCheck(game)) {
    return false
  }

  const possibleMoves = getAllPossibleMovesCheckmate(game)
  if (possibleMoves.length > 0) {
    return false
  }

  const possibleDrops = getAllPossibleDrops(game)
  for (let i = 0; i < possibleDrops.length; i++) {
    const movePos = possibleDrops[i]

    const piece = new Piece(movePos.x, movePos.y, PieceEnum.Fu, game.turn, false)
    const gameCopy = copyGame(game)
    gameCopy.board[movePos.y][movePos.x] = piece
    if (!getInCheck(gameCopy)) {
      return false
    }
  }

  return true
}

function getAllPossibleMovesCheckmate(game: Game): Vec2[] {
  const possibleMoves: Vec2[] = []
  //check normal moves
  for (let i = 0; i < game.height; i++) {
    for (let j = 0; j < game.width; j++) {
      const space = game.board[i][j]
      if (space !== null && space.owner === game.turn) {
        const pos: Vec2 = { x: j, y: i }
        possibleMoves.push(...getValidPieceMovesForCheckmate(pos, space, game))
      }
    }
  }

  return possibleMoves
}

function getValidPieceMovesForCheckmate(pos: Vec2, piece: Piece, game: Game): Vec2[] {
  const dir = getMoveDirection(game.turn)
  let possibleMoves: Vec2[] = []
  if (piece.type == PieceEnum.King) {
    possibleMoves = getKingMoves(pos, piece, game)
  } else {
    possibleMoves = getPieceMoves(pos, piece, game, dir)
  }

  possibleMoves = filterPossibleMoves(pos, possibleMoves, game)

  return possibleMoves
}

function getAllPossibleDrops(game: Game): Vec2[] {
  const possibleDrops: Vec2[] = []

  for (let i = 0; i < game.height; i++) {
    for (let j = 0; j < game.width; j++) {
      for (let k = 0; k < 7; k++) {
        const endPos: Vec2 = { x: j, y: i }
        const piece = new Piece(endPos.x, endPos.y, PieceEnum.Fu + k, game.turn, false)
        const result = checkValidDropMove(endPos, piece, game)
        if (result) {
          possibleDrops.push(endPos)
        }
      }
    }
  }

  return possibleDrops
}


function getInCheck(game: Game): boolean {
  const kings = getKingPos(game)
  for (let i = 0; i < kings.length; i++) {
    const king = kings[i]

    if (checkUnderAttack(king, game)) {
      return true
    }
  }

  return false
}

function getKingPos(game: Game): Vec2[] {
  const result: Vec2[] = []

  for (let i = 0; i < game.height; i++) {
    for (let j = 0; j < game.width; j++) {
      const space = game.board[i][j]

      if (space !== null && space.owner === game.turn && (space.type === PieceEnum.King || space.type === PieceEnum.Ou)) {
        result.push({ x: j, y: i })
      }
    }
  }

  return result
}

function checkUnderAttack(pos: Vec2, game: Game): boolean {
  const gameCopy = copyGame(game)
  const attackSpace: Map<string, boolean> = new Map<string, boolean>()

  gameCopy.turn = getEnemyTurnInt(game)
  for (let i = 0; i < gameCopy.height; i++) {
    for (let j = 0; j < gameCopy.width; j++) {
      const space = gameCopy.board[i][j]

      if (space !== null && space.owner === gameCopy.turn) {
        const pos = { x: j, y: i }
        const dir = getMoveDirection(gameCopy.turn)
        const possibleMoves = getPieceMoves(pos, space, gameCopy, dir)
        for (let k = 0; k < possibleMoves.length; k++) {
          const move = possibleMoves[k]
          const moveString = `${move.x},${move.y}`
          attackSpace.set(moveString, true)
        }
      }
    }
  }


  const posString = `${pos.x},${pos.y}`
  const underAttack = attackSpace.get(posString)
  if (underAttack === undefined) {
    return false
  } else {
    return underAttack
  }
}

function checkerMovesInCheck(startPos: Vec2, endPos: Vec2, piece: Piece, game: Game): boolean {
  if (!checkCheckerNextJumps(startPos, endPos, piece, game)) {
    game.board[startPos.y][startPos.x] = null
    game.board[endPos.y][endPos.x] = piece
    if (getInCheck(game)) {
      return true
    } else {
      return false
    }
  } else {

    const dir = getMoveDirection(game.turn)
    let possibleMoves: Vec2[] = []
    switch (piece.type) {
      case PieceEnum.Checker: {
        possibleMoves = getCheckerMoves(endPos, piece, game, dir)
        break
      }
      case PieceEnum.CheckerKing: {
        possibleMoves = getCheckerKingMoves(endPos, piece, game)
        break
      }
    }
    for (let i = possibleMoves.length - 1; i >= 0; i--) {
      if (!checkCheckerTake(endPos, possibleMoves[i])) {
        possibleMoves.splice(i, 1)
      }
    }

    for (let i = 0; i < possibleMoves.length; i++) {
      const movePos = { x: possibleMoves[i].x, y: possibleMoves[i].y }
      const gameCopy = copyGame(game)
      gameCopy.board[startPos.y][startPos.x] = null
      gameCopy.board[endPos.y][endPos.x] = piece

      const result = checkerMovesInCheck(endPos, movePos, piece, gameCopy)

      if (!result) {
        return false
      }
    }
  }

  return true
}

function checkCheckerNextJumps(startPos: Vec2, endPos: Vec2, piece: Piece, game: Game): boolean {
  if (!checkCheckerTake(startPos, endPos)) {
    return false
  }

  const dir = getMoveDirection(game.turn)
  let possibleMoves: Vec2[] = []

  switch (piece.type) {
    case PieceEnum.Checker: {
      possibleMoves = getCheckerMoves(endPos, piece, game, dir)
      break
    }
    case PieceEnum.CheckerKing: {
      possibleMoves = getCheckerKingMoves(endPos, piece, game)
      break
    }
    default:
      return false
  }

  for (let i = possibleMoves.length - 1; i >= 0; i--) {
    if (!checkCheckerTake(endPos, possibleMoves[i])) {
      possibleMoves.splice(i, 1)
    }
  }

  return possibleMoves.length > 0
}

function checkCheckerTake(startPos: Vec2, endPos: Vec2): boolean {
  const dx = Math.abs(startPos.x - endPos.x)
  const dy = Math.abs(startPos.y - endPos.y)

  if (dx === 2 && dy === 2) {
    return true
  }

  return false
}

export function checkValidDropMove(end: Vec2, piece: Piece, game: Game): boolean {
  if (!checkPieceOnBoard(end.x, end.y, game)) {
    return false
  }

  if (!checkHaveDropPiece(piece, game)) { return false }

  if (!checkEmptySpace(end, game)) { return false }

  if (!checkNifu(end, piece, game)) { return false }

  if (!checkIkidokoronoNaiKoma(end, piece, game)) { return false }

  if (!checkUtifudume(end, piece, game)) { return false }

  return true
}


function checkHaveDropPiece(piece: Piece, game: Game): boolean {
  if (!(PieceEnum.Fu <= piece.type && piece.type <= PieceEnum.Hi)) { return false }

  const mochigomaIndex = game.userSide * 7 + piece.type - PieceEnum.Fu
  if (game.mochigoma[mochigomaIndex] <= 0) {
    return false
  }

  return true
}

function checkEmptySpace(end: Vec2, game: Game): boolean {
  if (game.board[end.y][end.x] !== null) { return false }

  return true
}

function checkNifu(end: Vec2, piece: Piece, game: Game): boolean {
  if (piece.type !== PieceEnum.Fu) { return true }

  for (let i = 0; i < game.height; i++) {
    const space = game.board[i][end.x]
    if (space !== null && space.type === PieceEnum.Fu && space.owner == piece.owner) {
      return false
    }
  }

  return true
}

function checkIkidokoronoNaiKoma(end: Vec2, piece: Piece, game: Game): boolean {
  let row0
  let row1

  if (piece.owner == 0) {
    row0 = 0
    row1 = 1
  } else {
    row0 = game.height - 1
    row1 = game.height - 2
  }

  if (piece.type === PieceEnum.Fu || piece.type === PieceEnum.Kyou) {
    if (end.y === row0) {
      return false
    }
  } else if (piece.type === PieceEnum.Kei) {
    if (end.y == row0 || end.y === row1) {
      return false
    }
  }

  return true
}

function checkUtifudume(end: Vec2, piece: Piece, game: Game): boolean {
  if (piece.type !== PieceEnum.Fu) {
    return true
  }

  const gameCopy = copyGame(game)
  gameCopy.board[end.y][end.x] = piece

  gameCopy.turn = getEnemyTurnInt(gameCopy)

  const result = getInCheckmate(gameCopy)

  if (result) { return false }

  return true
}

export function checkPieceOnBoard(x: number, y: number, game: Game): boolean {
  const result = 0 <= x && x < game.width && 0 <= y && y < game.height
  return result
}
