import { Game } from "./game.ts"
import { Button } from "./button.ts"

import { BoardThemeColors } from "./themes.ts"


export class BoardRenderer2D {
  ctx: CanvasRenderingContext2D;
  canvas: HTMLCanvasElement
  UIRatio: number
  tileSize: number

  constructor(ctx: CanvasRenderingContext2D, canvas: HTMLCanvasElement, game: Game) {
    this.ctx = ctx
    this.canvas = canvas

    this.UIRatio = this.canvas.width / 1000

    this.tileSize = 800 * this.UIRatio / Math.max(game.width, game.height)
  }


  draw(game: Game, boardTheme: number, pieceTheme: number, buttons: Button[]) {
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

    this.#drawButtons(game, buttons)
  }

  #drawConnect(game: Game) {
    this.ctx.textAlign = "center"
    this.ctx.textBaseline = "middle";
    this.ctx.font = `${50 * this.UIRatio}px Arial`
    this.ctx.fillStyle = "#fff"
    this.ctx.fillText("Waiting for other player to connect...", this.canvas.width / 2, this.canvas.height / 4)

    this.ctx.font = `${15 * this.UIRatio}px Arial`
    this.ctx.fillStyle = "#fff"
    this.ctx.fillText("Click ID Button to copy ID", this.canvas.width / 2, this.canvas.height / 2 + this.UIRatio * 100)
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
    this.ctx.fillStyle = "#FFF"
    this.ctx.fillRect(100 * this.UIRatio, 100 * this.UIRatio, 800 * this.UIRatio, 800 * this.UIRatio)

    const xStart = this.canvas.width / 2 - this.tileSize * (game.width / 2)
    const yStart = this.canvas.height / 2 - this.tileSize * (game.height / 2)

    let boardThemeColor = BoardThemeColors.get(boardTheme)
    if (boardThemeColor === undefined) {
      boardThemeColor = { x: "#FFF", y: "#000" }
    }

    for (let i = 0; i < game.height; i++) {
      for (let j = 0; j < game.width; j++) {
        if ((i + j) % 2 == 0) {
          this.ctx.fillStyle = boardThemeColor.x
        } else {
          this.ctx.fillStyle = boardThemeColor.y
        }

        this.ctx.fillRect(xStart + j * this.tileSize, yStart + i * this.tileSize, this.tileSize, this.tileSize)

        const piece = game.board[i][j]
        if (piece === null) {
          continue
        }

        let color = "#000"
        if (piece.owner === 0) {
          color = "#FFF"
        }

        this.ctx.fillStyle = color
        this.ctx.textAlign = "center"
        this.ctx.textBaseline = "middle";
        this.ctx.font = `${this.tileSize / 2 * this.UIRatio}px Arial`
        this.ctx.fillText(piece.type.toString(), xStart + j * this.tileSize + this.tileSize / 2, yStart + i * this.tileSize + this.tileSize / 2)
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

  #drawButtons(game: Game, buttons: Button[]) {
    let screen = ""
    if (game.state === 0) {
      screen = "connect"
    }

    for (const button of buttons) {
      if (button.screen === screen) {
        button.visible = true
        button.draw(this.ctx)
      } else {
        button.visible = false
      }
    }
  }

  update(dt: number) {
  }

}
