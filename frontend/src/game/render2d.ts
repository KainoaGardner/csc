import { Game } from "./game.ts"
import { Piece } from "./piece.ts"
import { PieceEnum, PieceTypeToPrice } from "./util.ts"
import { PieceImages, PieceImageDimensions } from "./images.ts"
import { Button } from "./button.ts"

import { BoardThemeColors } from "./themes.ts"

const whiteShopPieces = [
  [
    new Piece(PieceEnum.Pawn, 0, false),
    new Piece(PieceEnum.Knight, 0, false),
    new Piece(PieceEnum.Bishop, 0, false),
    new Piece(PieceEnum.Rook, 0, false),
    new Piece(PieceEnum.Queen, 0, false),
    new Piece(PieceEnum.King, 0, false),
  ],
  [
    new Piece(PieceEnum.Fu, 0, false),
    new Piece(PieceEnum.Kyou, 0, false),
    new Piece(PieceEnum.Kei, 0, false),
    new Piece(PieceEnum.Gin, 0, false),
    new Piece(PieceEnum.Kin, 0, false),
    new Piece(PieceEnum.Kaku, 0, false),
    new Piece(PieceEnum.Hi, 0, false),
    new Piece(PieceEnum.Ou, 0, false),
  ],
  [
    new Piece(PieceEnum.Checker, 0, false),
  ]
]

const blackShopPieces = [
  [
    new Piece(PieceEnum.Pawn, 1, false),
    new Piece(PieceEnum.Knight, 1, false),
    new Piece(PieceEnum.Bishop, 1, false),
    new Piece(PieceEnum.Rook, 1, false),
    new Piece(PieceEnum.Queen, 1, false),
    new Piece(PieceEnum.King, 1, false),
  ],
  [
    new Piece(PieceEnum.Fu, 1, false),
    new Piece(PieceEnum.Kyou, 1, false),
    new Piece(PieceEnum.Kei, 1, false),
    new Piece(PieceEnum.Gin, 1, false),
    new Piece(PieceEnum.Kin, 1, false),
    new Piece(PieceEnum.Kaku, 1, false),
    new Piece(PieceEnum.Hi, 1, false),
    new Piece(PieceEnum.Ou, 1, false),
  ],
  [
    new Piece(PieceEnum.Checker, 1, false),
  ]

]


export class BoardRenderer2D {
  ctx: CanvasRenderingContext2D;
  canvas: HTMLCanvasElement
  UIRatio: number
  tileSize: number
  shopScreen: number = 0

  whiteShopPieces: Piece[][] = whiteShopPieces
  blackShopPieces: Piece[][] = blackShopPieces

  constructor(ctx: CanvasRenderingContext2D, canvas: HTMLCanvasElement, game: Game) {
    this.ctx = ctx
    this.canvas = canvas

    this.UIRatio = this.canvas.width / 1000

    this.tileSize = 800 * this.UIRatio / Math.max(game.width, game.height)

    this.switchShopScreen = this.switchShopScreen.bind(this)
  }


  draw(game: Game, boardTheme: number, buttons: Button[]) {
    this.ctx.fillStyle = "#111"
    this.ctx.fillRect(0, 0, this.canvas.width, this.canvas.height)

    switch (game.state) {
      case 0: {
        this.#drawConnect()
        break
      }
      case 1: {
        this.#drawPlace(game, boardTheme)
        break
      }
      case 2: {
        this.#drawMove(game, boardTheme)
        break
      }
      case 3: {
        this.#drawOver(game, boardTheme)
        break
      }
    }

    this.#drawButtons(game, buttons)
  }

  #drawConnect() {
    this.ctx.textAlign = "center"
    this.ctx.textBaseline = "middle";
    this.ctx.font = `${50 * this.UIRatio}px Arial`
    this.ctx.fillStyle = "#fff"
    this.ctx.fillText("Waiting for other player to connect...", this.canvas.width / 2, this.canvas.height / 4)

    this.ctx.font = `${15 * this.UIRatio}px Arial`
    this.ctx.fillStyle = "#fff"
    this.ctx.fillText("Click ID Button to copy ID", this.canvas.width / 2, this.canvas.height / 2 + this.UIRatio * 100)
  }

  #drawPlace(game: Game, boardTheme: number) {
    this.#drawBoard(game, boardTheme)
    this.#drawCover(game)
    this.#drawMochigoma(game, boardTheme)
    this.#drawShopPieces(game)
  }

  #drawMove(game: Game, boardTheme: number) {
    this.#drawBoard(game, boardTheme)
    this.#drawMochigoma(game, boardTheme)
  }

  #drawOver(game: Game, boardTheme: number) {
    this.#drawBoard(game, boardTheme)
    this.#drawMochigoma(game, boardTheme)
    this.#drawOverMessage(game)
  }

  #drawBoard(game: Game, boardTheme: number) {
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

        piece.draw(this.ctx, xStart + j * this.tileSize, yStart + i * this.tileSize, this.tileSize)
      }
    }
  }

  #drawMochigoma(game: Game, boardTheme: number) {

  }

  #drawShopPieces(game: Game) {
    this.ctx.textAlign = "center"
    this.ctx.textBaseline = "middle";
    this.ctx.lineWidth = 1 * this.UIRatio;
    this.ctx.font = `${35 * this.UIRatio}px Arial`
    this.ctx.fillStyle = "#FFD700"
    this.ctx.strokeStyle = "#000"


    if (game.userSide === 0) {
      const pieces = this.whiteShopPieces[this.shopScreen]

      for (let i = 0; i < pieces.length; i++) {
        const piece = pieces[i]
        piece.draw(this.ctx, 100 * this.UIRatio + this.tileSize * i, 900 * this.UIRatio, this.tileSize)
        const price = PieceTypeToPrice.get(piece.type)
        if (price === undefined) {
          continue
        }

        this.ctx.fillText(price.toString(), 100 * this.UIRatio + this.tileSize * i + this.tileSize / 2, 920 * this.UIRatio)
        this.ctx.strokeText(price.toString(), 100 * this.UIRatio + this.tileSize * i + this.tileSize / 2, 920 * this.UIRatio)
      }
    } else {
      const pieces = this.blackShopPieces[this.shopScreen]
      for (let i = 0; i < pieces.length; i++) {
        const piece = pieces[i]
        piece.draw(this.ctx, 100 * this.UIRatio + this.tileSize * i, 0, this.tileSize)
        const price = PieceTypeToPrice.get(piece.type)
        if (price === undefined) {
          continue
        }
        this.ctx.fillText(price.toString(), 100 * this.UIRatio + this.tileSize * i + this.tileSize / 2, 80 * this.UIRatio)
        this.ctx.strokeText(price.toString(), 100 * this.UIRatio + this.tileSize * i + this.tileSize / 2, 80 * this.UIRatio)
      }
    }
  }

  #drawCover(game: Game) {
    this.ctx.fillStyle = "#000"

    if (game.userSide === 0) {
      this.ctx.fillRect(100 * this.UIRatio, 100 * this.UIRatio, 800 * this.UIRatio, 400 * this.UIRatio)
    } else {
      this.ctx.fillRect(100 * this.UIRatio, 500 * this.UIRatio, 800 * this.UIRatio, 400 * this.UIRatio)
    }
  }

  #drawOverMessage(game: Game) {

  }

  #drawButtons(game: Game, buttons: Button[]) {
    let screen = ""
    if (game.state === 0) {
      screen = "connect"
    } else if (game.state === 1) {
      screen = "place"
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


  switchShopScreen() {
    this.shopScreen = (this.shopScreen + 1) % 3
    console.log(this.shopScreen)
  }

}
