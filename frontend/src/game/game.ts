import {
  FenStringToPieceInt,
  isCharDigit,
  isCharUppercase,
  convertStringToPosition,
  type Vec2,
} from "./util.ts"


import { Piece } from "./piece.ts"

export class Game {
  id: string
  userSide: number
  state = 0

  width: number
  height: number
  board: (Piece | null)[][]
  placeLine: number

  mochigoma: number[] = new Array(14).fill(0)

  money: Vec2
  time: Vec2

  turn: number = 0
  halfMoves: number = 0
  moves: number = 0

  enPassant: Vec2 | null = null
  checkerJump: Vec2 | null = null

  winner: number | null = null
  reason: string = ""

  constructor(id: string, width: number, height: number, placeLine: number, userSide: number, money: Vec2, time: Vec2) {
    this.id = id

    this.board = Array.from({ length: height }, () => Array(width).fill(null))
    this.width = width
    this.height = height
    this.placeLine = placeLine
    this.money = money
    this.time = time

    this.userSide = userSide
  }

  #updateBoard(fenPos: string) {
    const rows = fenPos.split("/")

    const pieceStrings: string[] = []
    const piecePositions: Vec2[] = []
    for (let i = 0; i < rows.length; i++) {
      let j = 0;
      let k = 0;

      let currPiece = ""
      while (j < rows[i].length) {
        if (currPiece.length === 3) {
          pieceStrings.push(currPiece)
          const pos: Vec2 = { x: k, y: i, }
          piecePositions.push(pos)
          currPiece = ""
          k++
        }

        const c = rows[i][j]
        if (isCharDigit(c)) {
          j++
          k += parseInt(c)
          currPiece = ""
          continue
        }

        currPiece += c
        j++
      }

      if (currPiece.length === 3) {
        pieceStrings.push(currPiece)
        const pos: Vec2 = { x: k, y: i, }
        piecePositions.push(pos)

      }
    }


    for (let i = 0; i < pieceStrings.length; i++) {
      const pos = piecePositions[i]
      const pieceString = pieceStrings[i]

      let owner = 0
      if (!isCharUppercase(pieceString[0])) {
        owner = 1
      }

      const type = FenStringToPieceInt.get(pieceString.slice(0, 2).toUpperCase())
      if (type === undefined) {
        console.log("Incorrect piece")
        continue
      }

      const moved = pieceString[2] === "-"

      const piece = new Piece(type, owner, moved)
      this.#setPiece(pos, piece)
    }
  }

  #setPiece(pos: Vec2, piece: Piece) {
    this.board[pos.y][pos.x] = piece
  }

  #clearBoard() {
    for (let i = 0; i < this.height; i++) {
      for (let j = 0; j < this.width; j++) {
        this.board[i][j] = null
      }
    }
  }

  #updateMochigoma(fenMochi: string) {
    const pieces = fenMochi.split("/")
    for (let i = 0; i < pieces.length; i++) {
      this.mochigoma[i] = parseInt(pieces[i])
    }
  }

  #updateTurn(turn: string) {
    if (turn === "w") {
      this.turn = 0
    } else {
      this.turn = 1
    }
  }

  #updateEnPassant(posString: string) {
    if (posString === "-") {
      this.enPassant = null
      return
    }

    this.enPassant = convertStringToPosition(posString, this.height)
  }

  #updateCheckerJump(posString: string) {
    if (posString === "-") {
      this.checkerJump = null
      return
    }

    this.checkerJump = convertStringToPosition(posString, this.height)
  }

  #updateMoveCounts(halfMoves: string, moves: string) {
    this.halfMoves = parseInt(halfMoves)
    this.moves = parseInt(moves)
  }

  #updateTime(fenTime: string) {
    const times = fenTime.split("/")
    this.time.x = parseInt(times[0])
    this.time.y = parseInt(times[1])
  }

  updateGame(fen: string) {

    const parts = fen.split(" ")
    if (parts.length !== 8) {
      console.log("ERROR fen part size incorrect")
      return
    }

    this.#clearBoard()

    this.#updateBoard(parts[0])
    this.#updateMochigoma(parts[1])
    this.#updateTurn(parts[2])
    this.#updateEnPassant(parts[3])
    this.#updateCheckerJump(parts[4])
    this.#updateMoveCounts(parts[5], parts[6])
    this.#updateTime(parts[7])
  }
}


