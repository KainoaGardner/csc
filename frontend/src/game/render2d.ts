import { Game } from "./game.ts"


export class BoardRenderer2D {
  ctx: CanvasRenderingContext2D;
  canvas: HTMLCanvasElement

  constructor(ctx: CanvasRenderingContext2D, canvas: HTMLCanvasElement) {
    this.ctx = ctx
    this.canvas = canvas
  }

  draw(game: Game, boardTheme: number, pieceTheme: number) {
    this.ctx.fillStyle = "#111"
    this.ctx.fillRect(0, 0, this.canvas.width, this.canvas.height)

    switch (game.state) {
      case 0: {
        this.#drawConnect(game)
        break
      }
      case 1: {
        this.#drawPlace(game, boardTheme, pieceTheme)
        break
      }
      case 2: {
        this.#drawMove(game, boardTheme, pieceTheme)
        break
      }
      case 3: {
        this.#drawOver(game, boardTheme, pieceTheme)
        break
      }



    }
    this.#drawBoard(game, boardTheme, pieceTheme)
    this.#drawMochigoma(game, boardTheme, pieceTheme)
  }

  #drawConnect(game: Game) {
  }

  #drawPlace(game: Game, boardTheme: number, pieceTheme: number) {
    this.#drawBoard(game, boardTheme, pieceTheme)
    this.#drawCover(game)
    this.#drawMochigoma(game, boardTheme, pieceTheme)
    this.#drawPieces(game, pieceTheme)
  }

  #drawMove(game: Game, boardTheme: number, pieceTheme: number) {
    this.#drawBoard(game, boardTheme, pieceTheme)
    this.#drawMochigoma(game, boardTheme, pieceTheme)
  }

  #drawOver(game: Game, boardTheme: number, pieceTheme: number) {
    this.#drawBoard(game, boardTheme, pieceTheme)
    this.#drawMochigoma(game, boardTheme, pieceTheme)
    this.#drawOverMessage(game)
  }

  #drawBoard(game: Game, boardTheme: number, pieceTheme: number) {
    for (let i = 0; i < game.height; i++) {
      for (let j = 0; j < game.width; j++) {
        const piece = game.board[i][j]
        if (piece === null) {
          continue
        }
      }
    }
  }

  #drawMochigoma(game: Game, boardTheme: number, pieceTheme: number) {

  }

  #drawPieces(game: Game, pieceTheme: number) {

  }

  #drawCover(game: Game) {

  }

  #drawOverMessage(game: Game) {

  }

  update(dt: number) {
  }
}

