import { Game } from "./game.ts"
import { BoardRenderer2D } from "./render2d.ts"
import { convertStringToPosition } from "./util.ts"

export class GameLog {
  id: string
  date: Date
  moveCount: number
  moves: string[]
  boardStates: string[]

  winner: number
  reason: string

  moveIndex: number = 0

  constructor(
    id: string,
    date: Date,
    moveCount: number,
    moves: string[],
    boardStates: string[],
    winner: number,
    reason: string,
  ) {

    this.id = id
    this.date = date
    this.moveCount = moveCount
    this.moves = moves
    this.boardStates = boardStates
    this.winner = winner
    this.reason = reason
  }

  updateSettings(
    id: string,
    date: Date,
    moveCount: number,
    moves: string[],
    boardStates: string[],
    winner: number,
    reason: string,
  ) {

    this.id = id
    this.date = date
    this.moveCount = moveCount
    this.moves = moves
    this.boardStates = boardStates
    this.winner = winner
    this.reason = reason

  }

  prevMove(game: Game, renderer: BoardRenderer2D) {
    if (this.moveIndex > 0) {
      this.moveIndex--

      this.updateLastMove(game, renderer)

      game.updateGame(this.boardStates[this.moveIndex])
    }
  }

  nextMove(game: Game, renderer: BoardRenderer2D) {
    if (this.moveIndex < this.boardStates.length - 1) {
      this.moveIndex++

      this.updateLastMove(game, renderer)
      game.updateGame(this.boardStates[this.moveIndex])
    }
  }

  updateLastMove(game: Game, renderer: BoardRenderer2D) {
    renderer.lastMove = {
      start: null,
      end: null,
    }

    if (this.moveIndex < this.moves.length) {
      const move = this.moves[this.moveIndex].split(",")
      if (move.length === 2) {
        const start = convertStringToPosition(move[0], game.height)
        const end = convertStringToPosition(move[1], game.height)
        renderer.lastMove = {
          start: start,
          end: end,
        }
      }
    }
  }
}

