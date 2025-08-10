import { Game } from "./game.ts"
import { Piece } from "./piece.ts"

import { getMoveDirection, getPieceMoves, filterPossibleMoves, checkPieceOnBoard } from "./engine.ts"
import {
  PieceEnum,
  PieceTypeToPrice,
  fitTextToWidth,
  type Message,
  convertSecondsToTimeString,
  type Annotation,
  checkEqualAnnotation,
  getAnnotationType,
  AnnotationEnum,

} from "./util.ts"
import { Button, createGameButtons } from "./button.ts"

import { BoardThemeColors } from "./themes.ts"
import { InputHandler } from "./inputHandler.ts"

const whiteShopPieces = [
  [
    new Piece(0, 8, PieceEnum.Pawn, 0, false),
    new Piece(1, 8, PieceEnum.Knight, 0, false),
    new Piece(2, 8, PieceEnum.Bishop, 0, false),
    new Piece(3, 8, PieceEnum.Rook, 0, false),
    new Piece(4, 8, PieceEnum.Queen, 0, false),
    new Piece(5, 8, PieceEnum.King, 0, false),
  ],
  [
    new Piece(0, 8, PieceEnum.Fu, 0, false),
    new Piece(1, 8, PieceEnum.Kyou, 0, false),
    new Piece(2, 8, PieceEnum.Kei, 0, false),
    new Piece(3, 8, PieceEnum.Gin, 0, false),
    new Piece(4, 8, PieceEnum.Kin, 0, false),
    new Piece(5, 8, PieceEnum.Kaku, 0, false),
    new Piece(6, 8, PieceEnum.Hi, 0, false),
    new Piece(7, 8, PieceEnum.Ou, 0, false),
  ],
  [
    new Piece(0, 8, PieceEnum.Checker, 0, false),
  ]
]

const blackShopPieces = [
  [
    new Piece(0, -1, PieceEnum.Pawn, 1, false),
    new Piece(1, -1, PieceEnum.Knight, 1, false),
    new Piece(2, -1, PieceEnum.Bishop, 1, false),
    new Piece(3, -1, PieceEnum.Rook, 1, false),
    new Piece(4, -1, PieceEnum.Queen, 1, false),
    new Piece(5, -1, PieceEnum.King, 1, false),
  ],
  [
    new Piece(0, -1, PieceEnum.Fu, 1, false),
    new Piece(1, -1, PieceEnum.Kyou, 1, false),
    new Piece(2, -1, PieceEnum.Kei, 1, false),
    new Piece(3, -1, PieceEnum.Gin, 1, false),
    new Piece(4, -1, PieceEnum.Kin, 1, false),
    new Piece(5, -1, PieceEnum.Kaku, 1, false),
    new Piece(6, -1, PieceEnum.Hi, 1, false),
    new Piece(7, -1, PieceEnum.Ou, 1, false),
  ],
  [
    new Piece(0, -1, PieceEnum.Checker, 1, false),
  ]

]

const whiteMochigomaPieces = [
  new Piece(8, 7, PieceEnum.Fu, 0, false),
  new Piece(8, 6, PieceEnum.Kyou, 0, false),
  new Piece(8, 5, PieceEnum.Kei, 0, false),
  new Piece(8, 4, PieceEnum.Gin, 0, false),
  new Piece(8, 3, PieceEnum.Kin, 0, false),
  new Piece(8, 2, PieceEnum.Kaku, 0, false),
  new Piece(8, 1, PieceEnum.Hi, 0, false),
]

const blackMochigomaPieces = [
  new Piece(-1, 0, PieceEnum.Fu, 1, false),
  new Piece(-1, 1, PieceEnum.Kyou, 1, false),
  new Piece(-1, 2, PieceEnum.Kei, 1, false),
  new Piece(-1, 3, PieceEnum.Gin, 1, false),
  new Piece(-1, 4, PieceEnum.Kin, 1, false),
  new Piece(-1, 5, PieceEnum.Kaku, 1, false),
  new Piece(-1, 6, PieceEnum.Hi, 1, false),
]


export class BoardRenderer2D {
  ctx: CanvasRenderingContext2D;
  canvas: HTMLCanvasElement
  UIRatio: number
  tileSize: number
  shopScreen: number = 0

  whiteShopPieces: Piece[][] = whiteShopPieces
  blackShopPieces: Piece[][] = blackShopPieces
  whiteMochigomaPieces: Piece[] = whiteMochigomaPieces
  blackMochigomaPieces: Piece[] = blackMochigomaPieces

  currAnnotation: Annotation = { start: null, end: null }
  annotations: Annotation[] = []

  buttons: Map<string, Button>

  constructor(ctx: CanvasRenderingContext2D,
    canvas: HTMLCanvasElement,
    game: Game,
    handleNotif: (err: string) => void,
    sendMessage: (msg: Message<unknown>) => void) {

    this.ctx = ctx
    this.canvas = canvas
    this.UIRatio = this.canvas.width / 1000
    this.tileSize = 800 * this.UIRatio / Math.max(game.width, game.height)
    this.switchShopScreen = this.switchShopScreen.bind(this)

    this.buttons = createGameButtons(canvas,
      this.UIRatio,
      game,
      handleNotif,
      sendMessage,
      this.switchShopScreen,
      game.clearBoardPlace,
      game.readyUp,
      game.unreadyUp,
    )

    this.updateButtonScreen(game)

    this.updateButtonScreen = this.updateButtonScreen.bind(this)
  }


  draw(game: Game, boardTheme: number, input: InputHandler) {
    this.ctx.fillStyle = "#111"
    this.ctx.fillRect(0, 0, this.canvas.width, this.canvas.height)


    this.#drawButtons()
    switch (game.state) {
      case 0: {
        this.#drawConnect()
        break
      }
      case 1: {
        this.#drawPlace(game, boardTheme, input)
        break
      }
      case 2: {
        this.#drawMove(game, boardTheme, input)
        break
      }
      case 3: {
        this.#drawOver(game, boardTheme, input)
        break
      }
    }

  }

  #drawConnect() {
    this.ctx.textAlign = "center"
    this.ctx.textBaseline = "middle";
    this.ctx.font = `${50 * this.UIRatio}px Arial`
    this.ctx.fillStyle = "#fff"
    this.ctx.fillText("Waiting for other player to connect...", this.canvas.width / 2, this.canvas.height / 4)
  }

  #drawPlace(game: Game, boardTheme: number, input: InputHandler) {
    this.#drawBoard(game, boardTheme)
    this.#drawCover(game)
    this.#drawSelectedSpace(game)
    this.#drawPlaceSpace(game, input)
    this.#drawBoardPieces(game, input)
    if (!game.ready[game.userSide])
      this.#drawShopPieces(game, input)
    this.#drawReadyText(game)
  }

  #drawMove(game: Game, boardTheme: number, input: InputHandler) {
    this.#drawBoard(game, boardTheme)
    this.#drawAnnotations()
    this.#drawSelectedSpace(game)
    this.#drawPlaceSpace(game, input)
    this.#drawTime(game,)

    this.#drawMochigoma(game, input)
    this.#drawBoardPieces(game, input)
    this.#drawMovableSpaces(game)
  }

  #drawOver(game: Game, boardTheme: number, input: InputHandler) {
    this.#drawBoard(game, boardTheme)
    this.#drawTime(game,)
    this.#drawMochigoma(game, input)
    this.#drawBoardPieces(game, input)
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
      }
    }
  }

  #drawAnnotations() {
    this.ctx.globalAlpha = 0.75
    this.ctx.fillStyle = "#e74c3c"
    this.ctx.strokeStyle = "#e74c3c"
    for (let i = 0; i < this.annotations.length; i++) {
      const anno = this.annotations[i]
      const annoType = getAnnotationType(anno)
      switch (annoType) {
        case AnnotationEnum.singleSpace: {
          this.#drawSingleSpaceAnnotation(anno)
          break
        }
        case AnnotationEnum.straightArrow: {
          this.#drawArrowAnnotation(anno)
          break
        }
        case AnnotationEnum.diagonalArrow: {
          this.#drawArrowAnnotation(anno)
          break
        }
        case AnnotationEnum.turnArrow: {
          this.#drawTurnAnnotation(anno)
          break
        }
      }
    }

    this.ctx.globalAlpha = 1.0
  }

  #drawSingleSpaceAnnotation(anno: Annotation) {
    if (anno.start === null || anno.end === null) {
      return
    }

    this.ctx.fillRect((anno.start.x + 1) * this.tileSize, (anno.start.y + 1) * this.tileSize, this.tileSize, this.tileSize)
  }

  #drawArrowAnnotation(anno: Annotation) {
    if (anno.start === null || anno.end === null) {
      return
    }

    this.ctx.lineWidth = this.tileSize * 0.25
    const boxSize = this.tileSize * 0.5

    const dx = anno.start.x - anno.end.x
    const dy = anno.start.y - anno.end.y
    const angle = Math.atan2(dy, dx)

    const startX = (anno.start.x + 1) * this.tileSize + this.tileSize * 0.5
    const startY = (anno.start.y + 1) * this.tileSize + this.tileSize * 0.5
    const endX = (anno.end.x + 1) * this.tileSize + this.tileSize * 0.5
    const endY = (anno.end.y + 1) * this.tileSize + this.tileSize * 0.5
    const endXLine = endX + boxSize * 0.5 * Math.cos(angle)
    const endYLine = endY + boxSize * 0.5 * Math.sin(angle)

    this.ctx.beginPath()

    this.ctx.moveTo(startX, startY)
    this.ctx.lineTo(endXLine, endYLine)
    this.ctx.stroke()

    this.ctx.save()
    this.ctx.translate(endX, endY)
    this.ctx.rotate(angle)
    this.ctx.fillRect(-boxSize * 0.5, -boxSize * 0.5, boxSize, boxSize)
    this.ctx.restore()
  }

  #drawTurnAnnotation(anno: Annotation) {
    if (anno.start === null || anno.end === null) {
      return
    }

    this.ctx.lineWidth = this.tileSize * 0.25
    const boxSize = this.tileSize * 0.5


    const startX = (anno.start.x + 1) * this.tileSize + this.tileSize * 0.5
    const startY = (anno.start.y + 1) * this.tileSize + this.tileSize * 0.5
    const endX = (anno.end.x + 1) * this.tileSize + this.tileSize * 0.5
    const endY = (anno.end.y + 1) * this.tileSize + this.tileSize * 0.5

    const turndx = Math.abs(anno.start.x - anno.end.x)
    const turndy = Math.abs(anno.start.y - anno.end.y)

    let turnX
    let turnY
    if (turndx === 1 && turndy === 2) {
      turnX = startX
      turnY = endY
    } else {
      turnX = endX
      turnY = startY
    }

    const dx = turnX - endX
    const dy = turnY - endY
    const angle = Math.atan2(dy, dx)
    const endXLine = endX + boxSize * 0.5 * Math.cos(angle)
    const endYLine = endY + boxSize * 0.5 * Math.sin(angle)


    this.ctx.beginPath()

    this.ctx.moveTo(startX, startY)
    this.ctx.lineTo(turnX, turnY)
    this.ctx.lineTo(endXLine, endYLine)
    this.ctx.stroke()

    this.ctx.fillRect(endX - boxSize * 0.5, endY - boxSize * 0.5, boxSize, boxSize)
  }

  #drawSelectedSpace(game: Game) {
    this.ctx.globalAlpha = 0.5
    this.ctx.fillStyle = "#e74c3c"
    for (let i = 0; i < game.height; i++) {
      for (let j = 0; j < game.width; j++) {
        const piece = game.board[i][j]
        if (piece === null) {
          continue
        }

        if (piece.selected) {
          this.ctx.fillRect((j + 1) * this.tileSize, (i + 1) * this.tileSize, this.tileSize, this.tileSize)
          break
        }
      }
    }

    this.ctx.globalAlpha = 1.0
  }

  #drawPlaceSpace(game: Game, input: InputHandler) {
    this.ctx.globalAlpha = 0.5
    this.ctx.fillStyle = "#2ecc71"

    let selectedFound: boolean = false
    for (let i = 0; i < game.height; i++) {
      for (let j = 0; j < game.width; j++) {
        const piece = game.board[i][j]
        if (piece === null) {
          continue
        }

        if (piece.selected) {
          selectedFound = true
          break
        }
      }
    }

    if (game.state === 1) {
      let pieces
      if (game.userSide === 0) {
        pieces = this.whiteShopPieces[this.shopScreen]
      } else {
        pieces = this.blackShopPieces[this.shopScreen]
      }

      for (let i = 0; i < pieces.length; i++) {
        const piece = pieces[i]
        if (piece === null) {
          continue
        }

        if (piece.selected) {
          selectedFound = true
          break
        }
      }
    } else if (game.state === 2) {
      for (let i = 0; i < 14; i++) {
        let piece: Piece
        if (i < 7) {
          piece = this.whiteMochigomaPieces[i]
        } else {
          piece = this.blackMochigomaPieces[i - 7]
        }

        if (piece !== null && piece.selected) {
          selectedFound = true
          break
        }
      }
    }

    if (selectedFound) {
      const placeX = Math.floor(input.mouse.x / this.tileSize)
      const placeY = Math.floor(input.mouse.y / this.tileSize)
      if (placeX - 1 >= 0 && placeX - 1 < game.width && placeY - 1 >= 0 && placeY - 1 < game.height) {
        this.ctx.fillRect(placeX * this.tileSize, placeY * this.tileSize, this.tileSize, this.tileSize)
      }
    }

    this.ctx.globalAlpha = 1.0
  }


  #drawBoardPieces(game: Game, input: InputHandler) {
    let selectedPiece: Piece | null = null

    for (let i = 0; i < game.height; i++) {
      for (let j = 0; j < game.width; j++) {
        const piece = game.board[i][j]
        if (piece === null) {
          continue
        }

        if (game.state === 1 &&
          (game.userSide === 0 && i < game.placeLine) ||
          (game.userSide === 1 && i >= game.placeLine)
        ) {
          continue
        }

        if (piece.selected) {
          selectedPiece = piece
        } else {
          piece.draw(this.ctx, this.tileSize, input)
        }
      }
    }

    if (selectedPiece !== null) {
      selectedPiece.draw(this.ctx, this.tileSize, input)
    }
  }

  #drawTime(game: Game) {
    this.ctx.textAlign = "center"
    this.ctx.textBaseline = "middle";
    this.ctx.lineWidth = 1 * this.UIRatio;
    this.ctx.fillStyle = "#FFF"
    this.ctx.strokeStyle = "#000"


    const whiteTimeString = convertSecondsToTimeString(game.time[0])
    const whiteTimeFontSize = fitTextToWidth(this.ctx, whiteTimeString, 90 * this.UIRatio, 50 * this.UIRatio, 5 * this.UIRatio)
    this.ctx.font = `${whiteTimeFontSize}px Arial Black`
    this.ctx.fillText(whiteTimeString, 900 * this.UIRatio + this.tileSize / 2, 900 * this.UIRatio + this.tileSize / 2)
    this.ctx.strokeText(whiteTimeString, 900 * this.UIRatio + this.tileSize / 2, 900 * this.UIRatio + this.tileSize / 2)

    const blackTimeString = convertSecondsToTimeString(game.time[1])
    const blackTimeFontSize = fitTextToWidth(this.ctx, whiteTimeString, 90 * this.UIRatio, 50 * this.UIRatio, 5 * this.UIRatio)
    this.ctx.font = `${blackTimeFontSize}px Arial Black`
    this.ctx.fillText(blackTimeString, this.tileSize / 2, this.tileSize / 2)
    this.ctx.strokeText(blackTimeString, this.tileSize / 2, this.tileSize / 2)

  }

  #drawReadyText(game: Game) {
    this.ctx.textAlign = "center"
    this.ctx.textBaseline = "middle";
    this.ctx.lineWidth = 2 * this.UIRatio;
    this.ctx.font = `${50 * this.UIRatio}px Arial Black`
    this.ctx.fillStyle = "#2ecc71"
    this.ctx.strokeStyle = "#000"

    if (game.ready[0]) {
      this.ctx.fillText("Ready", 500 * this.UIRatio, 900 * this.UIRatio + this.tileSize / 2)
      this.ctx.strokeText("Ready", 500 * this.UIRatio, 900 * this.UIRatio + this.tileSize / 2)
    }
    if (game.ready[1]) {
      this.ctx.fillText("Ready", 500 * this.UIRatio, this.tileSize / 2)
      this.ctx.strokeText("Ready", 500 * this.UIRatio, this.tileSize / 2)
    }
  }

  #drawMochigomaCover(game: Game) {
    this.ctx.globalAlpha = 0.5
    this.ctx.fillStyle = "#000"
    for (let i = 0; i < 14; i++) {
      let piece: Piece
      if (i < 7) {
        piece = this.whiteMochigomaPieces[i]
      } else {
        piece = this.blackMochigomaPieces[i - 7]
      }
      const count = game.mochigoma[i]

      if (count === 0) {
        this.ctx.fillRect((piece.x + 1) * this.tileSize * this.UIRatio, (piece.y + 1) * this.tileSize * this.UIRatio, this.tileSize, this.tileSize)
      }
    }

    this.ctx.globalAlpha = 1

  }

  #drawMochigoma(game: Game, input: InputHandler) {
    this.ctx.textAlign = "center"
    this.ctx.textBaseline = "middle";
    this.ctx.lineWidth = 2 * this.UIRatio;
    this.ctx.font = `${40 * this.UIRatio}px Arial Black`
    this.ctx.fillStyle = "#FFF"
    this.ctx.strokeStyle = "#000"

    let selectedPiece: Piece | null = null
    for (let i = 0; i < 14; i++) {
      let piece: Piece
      let textX: number
      let textY: number
      if (i < 7) {
        piece = this.whiteMochigomaPieces[i]
        textX = (piece.x + 1) * this.tileSize * this.UIRatio + this.tileSize / 2 + this.tileSize / 4
        textY = (piece.y + 1) * this.tileSize * this.UIRatio + this.tileSize / 4
      } else {
        piece = this.blackMochigomaPieces[i - 7]
        textX = (piece.x + 1) * this.tileSize * this.UIRatio + this.tileSize / 4
        textY = (piece.y + 1) * this.tileSize * this.UIRatio + this.tileSize / 2 + this.tileSize / 4
      }


      if (piece.selected) {
        selectedPiece = piece
      } else {
        piece.draw(this.ctx, this.tileSize, input)
      }


      this.ctx.fillStyle = "#FFF"
      const count = game.mochigoma[i]
      this.ctx.fillText(count.toString(), textX, textY)
      this.ctx.strokeText(count.toString(), textX, textY)
    }

    this.#drawMochigomaCover(game)

    if (selectedPiece !== null) {
      selectedPiece.draw(this.ctx, this.tileSize, input)
    }
  }

  #drawShopPieces(game: Game, input: InputHandler) {
    this.ctx.textAlign = "center"
    this.ctx.textBaseline = "middle";
    this.ctx.lineWidth = 2 * this.UIRatio;
    this.ctx.font = `${35 * this.UIRatio}px Arial Black`
    this.ctx.fillStyle = "#FFF"
    this.ctx.strokeStyle = "#000"


    let pieces
    let yStartPriceFont
    let yStartMoneyFont
    let xStartPriceFont
    if (game.userSide === 0) {
      pieces = this.whiteShopPieces[this.shopScreen]
      yStartPriceFont = 920 * this.UIRatio
      yStartMoneyFont = 900 * this.UIRatio
      xStartPriceFont = 125 * this.UIRatio
    } else {
      pieces = this.blackShopPieces[this.shopScreen]
      yStartPriceFont = 80 * this.UIRatio
      yStartMoneyFont = 0
      xStartPriceFont = 75 * this.UIRatio
    }

    let selectedPiece: Piece | null = null
    for (let i = 0; i < pieces.length; i++) {
      const piece = pieces[i]

      if (piece.selected) {
        selectedPiece = piece
      } else {
        piece.draw(this.ctx, this.tileSize, input)
      }

      const price = PieceTypeToPrice.get(piece.type)
      if (price === undefined) {
        continue
      }

      this.ctx.fillText(price.toString(), xStartPriceFont + this.tileSize * i + this.tileSize / 2, yStartPriceFont)
      this.ctx.strokeText(price.toString(), xStartPriceFont + this.tileSize * i + this.tileSize / 2, yStartPriceFont)
    }

    const moneyFontSize = fitTextToWidth(this.ctx, game.money[game.userSide].toString(), 90 * this.UIRatio, 50 * this.UIRatio, 5 * this.UIRatio)
    this.ctx.font = `${moneyFontSize}px Arial Black`
    this.ctx.fillText(game.money[game.userSide].toString(), 900 * this.UIRatio + this.tileSize / 2, yStartMoneyFont + this.tileSize / 2)
    this.ctx.strokeText(game.money[game.userSide].toString(), 900 * this.UIRatio + this.tileSize / 2, yStartMoneyFont + this.tileSize / 2)


    if (selectedPiece) {
      selectedPiece.draw(this.ctx, this.tileSize, input)
    }
  }

  #drawCover(game: Game) {
    this.ctx.fillStyle = "#555"
    this.ctx.globalAlpha = 0.75

    const placeLinePixel = game.placeLine * this.tileSize

    if (game.userSide === 0) {
      this.ctx.fillRect(100 * this.UIRatio, 100 * this.UIRatio, 800 * this.UIRatio, placeLinePixel * this.UIRatio)
    } else {
      this.ctx.fillRect(100 * this.UIRatio, (placeLinePixel + 100) * this.UIRatio, 800 * this.UIRatio, 800 - placeLinePixel * this.UIRatio)
    }

    this.ctx.globalAlpha = 1.0
  }

  #drawMovableSpaces(game: Game) {
    let selectedPiece: Piece | null = null
    for (let i = 0; i < game.height; i++) {
      for (let j = 0; j < game.width; j++) {
        const piece = game.board[i][j]
        if (piece === null || piece.owner !== game.userSide) {
          continue
        }

        if (piece.selected) {
          selectedPiece = piece
        }
      }
    }

    if (selectedPiece === null) {
      return
    }

    const dir = getMoveDirection(game.turn)
    const startPos = { x: selectedPiece.x, y: selectedPiece.y }
    const possibleMoves = getPieceMoves(startPos, selectedPiece, game, dir)
    const filteredMoves = filterPossibleMoves(startPos, possibleMoves, game)

    if (game.userSide === 0) {
      this.ctx.fillStyle = "#FFF"
      this.ctx.strokeStyle = "#000"
    } else {
      this.ctx.fillStyle = "#000"
      this.ctx.strokeStyle = "#FFF"
    }

    this.ctx.globalAlpha = 0.8
    this.ctx.lineWidth = 2 * this.UIRatio

    for (let i = 0; i < filteredMoves.length; i++) {
      const move = filteredMoves[i]
      this.ctx.beginPath()
      this.ctx.arc((move.x + 1) * this.tileSize + this.tileSize / 2, (move.y + 1) * this.tileSize + this.tileSize / 2, this.tileSize / 10, 0, 2 * Math.PI)
      this.ctx.fill()
      this.ctx.stroke()
    }

    this.ctx.globalAlpha = 1.0
  }

  #drawOverMessage(game: Game) {

  }

  #drawButtons() {
    for (const button of this.buttons.values()) {
      if (button.visible) {
        button.draw(this.ctx)
      }
    }
  }

  update(game: Game, input: InputHandler, sendMessage: (msg: Message<unknown>) => void) {
    this.#updateButtons(input, game)

    switch (game.state) {
      case 1:
        this.#placeUpdate(game, input, sendMessage)
        break
      case 2:
        this.#moveUpdate(game, input, sendMessage)
        break
    }
  }

  #updateButtons(input: InputHandler, game: Game) {
    for (const button of this.buttons.values()) {
      if (button.visible) {
        button.update(input, game, this.updateButtonScreen)
      }
    }
  }

  updateButtonScreen(game: Game) {
    for (const button of this.buttons.values()) {
      button.visible = false
    }

    switch (game.state) {
      case 0: {
        this.buttons.get("id")!.visible = true
        break
      }
      case 1: {
        if (game.ready[game.userSide]) {
          this.buttons.get("unready")!.visible = true
        } else {
          this.buttons.get("shop")!.visible = true
          this.buttons.get("ready")!.visible = true
          this.buttons.get("clear")!.visible = true
        }
        break
      }
      case 2: {
        this.buttons.get("resign")!.visible = true
        if (game.draw[game.userSide]) {
          this.buttons.get("undraw")!.visible = true
        } else {
          this.buttons.get("draw")!.visible = true
        }
        break
      }
      case 3: {
        break
      }
    }
  }

  #placeUpdate(game: Game, input: InputHandler, sendMessage: (msg: Message<unknown>) => void) {
    let pieces
    if (game.userSide === 0) {
      pieces = this.whiteShopPieces[this.shopScreen]
    } else {
      pieces = this.blackShopPieces[this.shopScreen]
    }

    for (let i = 0; i < pieces.length; i++) {
      const piece = pieces[i]
      piece.placeUpdate(game, this.tileSize, input, sendMessage)
    }

    for (let i = 0; i < game.height; i++) {
      for (let j = 0; j < game.width; j++) {
        const piece = game.board[i][j]
        if (piece === null) {
          continue
        }

        piece.placeUpdate(game, this.tileSize, input, sendMessage)
      }
    }
  }

  #moveUpdate(game: Game, input: InputHandler, sendMessage: (msg: Message<unknown>) => void) {
    this.#annotationUpdate(game, input)

    for (let i = 0; i < game.height; i++) {
      for (let j = 0; j < game.width; j++) {
        const piece = game.board[i][j]
        if (piece === null) {
          continue
        }

        piece.moveUpdate(game, this.tileSize, input, sendMessage)
      }
    }

    let mochigomaPieces
    let mochigomaOffset
    if (game.userSide === 0) {
      mochigomaPieces = this.whiteMochigomaPieces
      mochigomaOffset = 0
    } else {
      mochigomaPieces = this.blackMochigomaPieces
      mochigomaOffset = 7
    }

    for (let i = 0; i < mochigomaPieces.length; i++) {
      const piece = mochigomaPieces[i]
      const amount = game.mochigoma[mochigomaOffset + i]
      if (amount > 0) {
        piece.moveMochigomaUpdate(game, this.tileSize, input, sendMessage)
      }
    }
  }

  #annotationUpdate(game: Game, input: InputHandler) {
    if (input.mouse.justPressed[0]) {
      this.annotations = []
    }

    if (input.mouse.justPressed[2]) {
      const placeX = Math.floor(input.mouse.x / this.tileSize) - 1
      const placeY = Math.floor(input.mouse.y / this.tileSize) - 1

      if (checkPieceOnBoard(placeX, placeY, game)) {
        this.currAnnotation.start = { x: placeX, y: placeY }
      } else {
        this.annotations = []
      }
    }
    if (input.mouse.justReleased[2]) {
      if (this.currAnnotation.start !== null) {
        const placeX = Math.floor(input.mouse.x / this.tileSize) - 1
        const placeY = Math.floor(input.mouse.y / this.tileSize) - 1

        if (checkPieceOnBoard(placeX, placeY, game)) {
          this.currAnnotation.end = { x: placeX, y: placeY }
          const newStartPos = { x: this.currAnnotation.start.x, y: this.currAnnotation.start.y }
          const newEndPos = { x: placeX, y: placeY }
          const newAnnotation = { start: newStartPos, end: newEndPos }

          let found = false
          for (let i = this.annotations.length - 1; i >= 0; i--) {
            const anno = this.annotations[i]
            if (checkEqualAnnotation(newAnnotation, anno)) {
              this.annotations.splice(i, 1)
              found = true
            }
          }

          if (!found) {
            this.annotations.push(newAnnotation)
          }
        }
      }
      this.currAnnotation = { start: null, end: null }
    }
  }

  switchShopScreen() {
    this.shopScreen = (this.shopScreen + 1) % 3
  }

}
