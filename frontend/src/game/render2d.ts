import { Game } from "./game.ts"
import { GameLog } from "./gameLog.ts"
import { Piece } from "./piece.ts"

import {
  getMoveDirection,
  getPieceMoves,
  filterPossibleMoves,
  checkPieceOnBoard,
  getPieceDrops,
  copyGame,
  getEnemyTurnInt,
} from "./engine.ts"
import {
  PieceEnum,
  PieceTypeToPrice,
  fitTextToWidth,
  convertSecondsToTimeString,
  checkEqualAnnotation,
  getAnnotationType,
  AnnotationEnum,
  sendResignMessage,
  convertMoveToString,
  sendMoveMessage,
  PromoteTypeEnum,
  checkVec2Equal,
  convertStringToPosition,
  type PendingMove,
  type Message,
  type Annotation,
  type Vec2,
  type Move,

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
    new Piece(7, -1, PieceEnum.Pawn, 1, false),
    new Piece(6, -1, PieceEnum.Knight, 1, false),
    new Piece(5, -1, PieceEnum.Bishop, 1, false),
    new Piece(4, -1, PieceEnum.Rook, 1, false),
    new Piece(3, -1, PieceEnum.Queen, 1, false),
    new Piece(2, -1, PieceEnum.King, 1, false),
  ],
  [
    new Piece(7, -1, PieceEnum.Fu, 1, false),
    new Piece(6, -1, PieceEnum.Kyou, 1, false),
    new Piece(5, -1, PieceEnum.Kei, 1, false),
    new Piece(4, -1, PieceEnum.Gin, 1, false),
    new Piece(3, -1, PieceEnum.Kin, 1, false),
    new Piece(2, -1, PieceEnum.Kaku, 1, false),
    new Piece(1, -1, PieceEnum.Hi, 1, false),
    new Piece(0, -1, PieceEnum.Ou, 1, false),
  ],
  [
    new Piece(7, -1, PieceEnum.Checker, 1, false),
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

const whitePromotePieces = [
  new Piece(0, 0, PieceEnum.Knight, 0, false),
  new Piece(0, 0, PieceEnum.Bishop, 0, false),
  new Piece(0, 0, PieceEnum.Rook, 0, false),
  new Piece(0, 0, PieceEnum.Queen, 0, false),
]

const blackPromotePieces = [
  new Piece(0, 0, PieceEnum.Knight, 1, false),
  new Piece(0, 0, PieceEnum.Bishop, 1, false),
  new Piece(0, 0, PieceEnum.Rook, 1, false),
  new Piece(0, 0, PieceEnum.Queen, 1, false),
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
  whitePromotePieces: Piece[] = whitePromotePieces
  blackPromotePieces: Piece[] = blackPromotePieces

  currAnnotation: Annotation = { start: null, end: null }
  annotations: Annotation[] = []

  buttons: Map<string, Button>

  pendingMove: PendingMove | null = null

  lastFrameTime = Date.now()
  lastMoveTime = Date.now()


  lastMove: Annotation = { start: null, end: null }

  resignPressed: boolean = false


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

    this.pressResign = this.pressResign.bind(this)
    this.cancelResign = this.cancelResign.bind(this)
    this.confirmResign = this.confirmResign.bind(this)

    this.setPendingMove = this.setPendingMove.bind(this)
    this.sendPendingMove = this.sendPendingMove.bind(this)

    this.buttons = createGameButtons(canvas,
      this.UIRatio,
      game,
      handleNotif,
      sendMessage,
      this.switchShopScreen,
      game.clearBoardPlace,
      game.readyUp,
      game.unreadyUp,
      game.drawUp,
      game.undrawUp,
      this.pressResign,
      this.cancelResign,
      this.confirmResign,
      this.sendPendingMove,
    )

    this.updateButtonScreen(game)

    this.updateButtonScreen = this.updateButtonScreen.bind(this)
  }

  updateButtons(
    canvas: HTMLCanvasElement,
    game: Game,
    handleNotif: (err: string) => void,
    sendMessage: (msg: Message<unknown>) => void) {

    this.buttons = createGameButtons(
      canvas,
      this.UIRatio,
      game,
      handleNotif,
      sendMessage,
      this.switchShopScreen,
      game.clearBoardPlace,
      game.readyUp,
      game.unreadyUp,
      game.drawUp,
      game.undrawUp,
      this.pressResign,
      this.cancelResign,
      this.confirmResign,
      this.sendPendingMove,
    )
  }

  draw(game: Game, boardTheme: number, input: InputHandler) {
    this.ctx.fillStyle = "#111"
    this.ctx.fillRect(0, 0, this.canvas.width, this.canvas.height)

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
    this.#drawButtons(game, input)
  }

  drawGameLog(game: Game, gameLog: GameLog, boardTheme: number, input: InputHandler) {
    this.ctx.fillStyle = "#111"
    this.ctx.fillRect(0, 0, this.canvas.width, this.canvas.height)

    this.#drawGameLog(game, gameLog, boardTheme, input)

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
    this.#drawSpaceAnnotations()
    this.#drawPlaceSpace(game, input)
    this.#drawBoardPieces(game, input)
    if (!game.ready[game.userSide])
      this.#drawShopPieces(game, input)
    this.#drawAnnotations()
    this.#drawReadyText(game)
  }

  #drawMove(game: Game, boardTheme: number, input: InputHandler) {
    this.#drawBoard(game, boardTheme)
    this.#drawSpaceAnnotations()
    this.#drawSelectedSpace(game)
    this.#drawPlaceSpace(game, input)
    this.#drawLastMoveSpace()
    this.#drawTime(game)
    this.#drawDrawText(game)
    this.#drawTurnText(game)

    this.#drawMochigoma(game, input)
    this.#drawBoardPieces(game, input)
    this.#drawAnnotations()
    this.#drawMovableSpaces(game)
  }

  #drawOver(game: Game, boardTheme: number, input: InputHandler) {
    this.#drawBoard(game, boardTheme)
    this.#drawSpaceAnnotations()
    this.#drawTime(game,)
    this.#drawMochigoma(game, input)
    this.#drawBoardPieces(game, input)
    this.#drawAnnotations()
    this.#drawOverMessage(game)
  }

  #drawGameLog(game: Game, gameLog: GameLog, boardTheme: number, input: InputHandler) {
    this.#drawBoard(game, boardTheme)
    this.#drawSpaceAnnotations()
    this.#drawLastMoveSpace()
    this.#drawTime(game)
    this.#drawTurnText(game)

    this.#drawMochigoma(game, input)
    this.#drawBoardPieces(game, input)
    this.#drawAnnotations()


    if (gameLog.moveIndex === gameLog.boardStates.length) {
      this.#drawOverMessage(game)
    }
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

  #drawSpaceAnnotations() {
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
      }
    }

    this.ctx.globalAlpha = 1.0
  }


  #drawLastMoveSpace() {
    if (this.lastMove.start === null || this.lastMove.end === null) return

    this.ctx.globalAlpha = 0.5
    this.ctx.fillStyle = "#f1c40f"

    this.ctx.fillRect((this.lastMove.start.x + 1) * this.tileSize, (this.lastMove.start.y + 1) * this.tileSize, this.tileSize, this.tileSize)

    this.ctx.fillStyle = "#c0392b"
    this.ctx.fillRect((this.lastMove.end.x + 1) * this.tileSize, (this.lastMove.end.y + 1) * this.tileSize, this.tileSize, this.tileSize)


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

        if (game.state === 1 && (
          (game.userSide === 0 && i < game.placeLine) ||
          (game.userSide === 1 && i >= game.placeLine))
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

  #drawDrawText(game: Game) {
    this.ctx.textAlign = "center"
    this.ctx.textBaseline = "middle";
    this.ctx.lineWidth = 2 * this.UIRatio;
    this.ctx.font = `${35 * this.UIRatio}px Arial Black`
    this.ctx.fillStyle = "#2ecc71"
    this.ctx.strokeStyle = "#000"

    if (game.draw[0] && game.userSide !== 0) {
      this.ctx.fillText("Draw", 50 * this.UIRatio, 900 * this.UIRatio + this.tileSize / 2)
      this.ctx.strokeText("Draw", 50 * this.UIRatio, 900 * this.UIRatio + this.tileSize / 2)
    }
    if (game.draw[1] && game.userSide !== 1) {
      this.ctx.fillText("Draw", 950 * this.UIRatio, this.tileSize / 2)
      this.ctx.strokeText("Draw", 950 * this.UIRatio, this.tileSize / 2)
    }
  }

  #drawTurnText(game: Game) {
    this.ctx.textAlign = "center"
    this.ctx.textBaseline = "middle";
    this.ctx.lineWidth = 2 * this.UIRatio;
    this.ctx.font = `${50 * this.UIRatio}px Arial Black`

    if (game.turn === 0) {
      this.ctx.fillStyle = "#FFF"
      this.ctx.strokeStyle = "#000"
      this.ctx.fillText("White Turn", 500 * this.UIRatio, 900 * this.UIRatio + this.tileSize / 2)
      this.ctx.strokeText("White Turn", 500 * this.UIRatio, 900 * this.UIRatio + this.tileSize / 2)
    } else {
      this.ctx.fillStyle = "#000"
      this.ctx.strokeStyle = "#FFF"
      this.ctx.fillText("Black Turn", 500 * this.UIRatio, this.tileSize / 2)
      this.ctx.strokeText("Black Turn", 500 * this.UIRatio, this.tileSize / 2)
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
    let xStartPriceFont
    let xStartMoneyFont
    let yStartMoneyFont
    let shopDirection
    if (game.userSide === 0) {
      pieces = this.whiteShopPieces[this.shopScreen]
      yStartPriceFont = 920 * this.UIRatio
      xStartPriceFont = 125 * this.UIRatio
      xStartMoneyFont = 900 * this.UIRatio
      yStartMoneyFont = 900 * this.UIRatio
      shopDirection = 1
    } else {
      pieces = this.blackShopPieces[this.shopScreen]
      yStartPriceFont = 80 * this.UIRatio
      // xStartPriceFont = 75 * this.UIRatio
      xStartPriceFont = 825 * this.UIRatio
      xStartMoneyFont = 0 * this.UIRatio
      yStartMoneyFont = 0
      shopDirection = -1
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

      this.ctx.fillText(price.toString(), xStartPriceFont + this.tileSize * i * shopDirection + this.tileSize / 2, yStartPriceFont)
      this.ctx.strokeText(price.toString(), xStartPriceFont + this.tileSize * i * shopDirection + this.tileSize / 2, yStartPriceFont)
    }

    const moneyFontSize = fitTextToWidth(this.ctx, game.money[game.userSide].toString(), 90 * this.UIRatio, 50 * this.UIRatio, 5 * this.UIRatio)
    this.ctx.font = `${moneyFontSize}px Arial Black`
    this.ctx.fillText(game.money[game.userSide].toString(), xStartMoneyFont + this.tileSize / 2, yStartMoneyFont + this.tileSize / 2)
    this.ctx.strokeText(game.money[game.userSide].toString(), xStartMoneyFont + this.tileSize / 2, yStartMoneyFont + this.tileSize / 2)


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
    if (game.turn !== game.userSide) return
    let filteredMoves: Vec2[] = []

    let selectedPieceType: string = ""
    let selectedPiece: Piece | null = null
    for (let i = 0; i < game.height; i++) {
      for (let j = 0; j < game.width; j++) {
        const piece = game.board[i][j]
        if (piece === null || piece.owner !== game.userSide) {
          continue
        }

        if (piece.selected) {
          selectedPiece = piece
          selectedPieceType = "board"
        }
      }
    }

    let mochigomaPieces: Piece[]
    if (game.userSide === 0) {
      mochigomaPieces = this.whiteMochigomaPieces
    } else {
      mochigomaPieces = this.blackMochigomaPieces
    }

    for (let i = 0; i < mochigomaPieces.length; i++) {
      const piece = mochigomaPieces[i]

      if (piece.selected) {
        selectedPiece = piece
        selectedPieceType = "mochigoma"
      }
    }

    if (selectedPiece === null) {
      return
    }

    if (selectedPieceType === "board") {
      const dir = getMoveDirection(game.userSide)
      const startPos = { x: selectedPiece.x, y: selectedPiece.y }
      if (game.checkerJump === null || checkVec2Equal(startPos, game.checkerJump)) {
        const gameCopy = copyGame(game)
        gameCopy.turn = game.userSide
        const possibleMoves = getPieceMoves(startPos, selectedPiece, gameCopy, dir)
        filteredMoves = filterPossibleMoves(startPos, possibleMoves, gameCopy)
      }
    } else {
      if (game.checkerJump === null) {
        filteredMoves = getPieceDrops(selectedPiece, game)
      }
    }


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
    this.ctx.textAlign = "center"
    this.ctx.textBaseline = "middle";
    if (game.winner === null) return

    let winnerText
    let reasonText

    const winnerTextY = 450 * this.UIRatio
    const reasonTextY = 550 * this.UIRatio
    if (game.winner === 0) {
      this.ctx.fillStyle = "#FFF"
      this.ctx.strokeStyle = "#000"
      winnerText = "White Wins"
      reasonText = game.reason
    } else if (game.winner === 1) {
      this.ctx.fillStyle = "#000"
      this.ctx.strokeStyle = "#FFF"
      winnerText = "Black Wins"
      reasonText = game.reason
    } else {
      this.ctx.fillStyle = "#DDD"
      this.ctx.strokeStyle = "#333"
      winnerText = "Draw"
      reasonText = game.reason
    }


    this.ctx.lineWidth = 3 * this.UIRatio;
    this.ctx.font = `${75 * this.UIRatio}px Arial Black`

    this.ctx.fillText(winnerText, 500 * this.UIRatio, winnerTextY)
    this.ctx.strokeText(winnerText, 500 * this.UIRatio, winnerTextY)

    this.ctx.lineWidth = 2 * this.UIRatio;
    this.ctx.font = `${50 * this.UIRatio}px Arial Black`
    this.ctx.fillText(reasonText, 500 * this.UIRatio, reasonTextY)
    this.ctx.strokeText(reasonText, 500 * this.UIRatio, reasonTextY)


  }

  #drawButtons(game: Game, input: InputHandler) {
    for (const button of this.buttons.values()) {
      if (!button.visible) { continue }

      button.draw(this.ctx)

      this.#drawChessPromoteButtons(button, game, input)
    }
  }

  #drawChessPromoteButtons(button: Button, game: Game, input: InputHandler) {
    if (button.text === "K" || button.text === "B" || button.text === "R" || button.text === "Q") {
      let pieceIndex
      let pieces
      if (game.userSide === 0) {
        pieces = this.whitePromotePieces
      } else {
        pieces = this.blackPromotePieces
      }

      switch (button.text) {
        case "K": {
          pieceIndex = 0
          break
        }
        case "B": {
          pieceIndex = 1
          break
        }
        case "R": {
          pieceIndex = 2
          break
        }
        case "Q": {
          pieceIndex = 3
          break
        }
      }

      const piece = pieces[pieceIndex]
      piece.x = (button.x / this.tileSize) - 1
      piece.y = (button.y / this.tileSize) - 1
      piece.draw(this.ctx, this.tileSize, input)
    }

  }

  update(game: Game, input: InputHandler, sendMessage: (msg: Message<unknown>) => void) {
    this.#updateButtons(input, game)

    switch (game.state) {
      case 1:
        this.#annotationUpdate(game, input)
        this.#placeUpdate(game, input, sendMessage)
        break
      case 2:
        this.#annotationUpdate(game, input)
        this.#moveUpdate(game, input, sendMessage)
        this.#timeUpdate(game)
        break
      case 3:
        this.#annotationUpdate(game, input)
        break
      case 4:
        this.#annotationUpdate(game, input)
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

  #timeUpdate(game: Game) {
    game.updateClientTime(this.lastMoveTime, this.lastFrameTime)
    this.lastFrameTime = Date.now()
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
        if (this.resignPressed) {
          this.buttons.get("resign")!.visible = false
          this.buttons.get("confirm")!.visible = true
          this.buttons.get("cancel")!.visible = true
        } else {
          this.buttons.get("resign")!.visible = true
          this.buttons.get("confirm")!.visible = false
          this.buttons.get("cancel")!.visible = false
        }

        if (game.draw[game.userSide]) {
          this.buttons.get("undraw")!.visible = true
          this.buttons.get("draw")!.visible = false
        } else {
          this.buttons.get("undraw")!.visible = false
          this.buttons.get("draw")!.visible = true
        }

        this.#updatePromoteButtons(game)

        break
      }
      case 3: {
        break
      }
    }
  }

  #updatePromoteButtons(game: Game) {
    if (this.pendingMove === null) { return }

    if (this.pendingMove.type === PromoteTypeEnum.chess) {
      this.#updateChessPromoteButtons(game)
    } else if (this.pendingMove.type === PromoteTypeEnum.shogi) {
      this.#updateShogiPromoteButtons(game)
    }
  }

  #updateChessPromoteButtons(game: Game) {
    if (this.pendingMove === null) { return }

    const end = this.pendingMove.move.end
    const knightButton = this.buttons.get("chessPromoteK")!
    const bishopButton = this.buttons.get("chessPromoteB")!
    const rookButton = this.buttons.get("chessPromoteR")!
    const queenButton = this.buttons.get("chessPromoteQ")!

    let startY
    let startX = (end.x - 1) * this.tileSize + this.tileSize / 2
    if (game.userSide === 0) {
      startY = end.y * this.tileSize
    } else {
      startY = (end.y + 2) * this.tileSize
      startX = (end.x - 1) * this.tileSize + this.tileSize / 2
    }

    if (startX < 100 * this.UIRatio) {
      startX = 100 * this.UIRatio
    }
    if (startX + 4 * this.tileSize > 900 * this.UIRatio) {
      startX = 500 * this.UIRatio
    }

    knightButton.x = startX
    bishopButton.x = startX + this.tileSize * 1
    rookButton.x = startX + this.tileSize * 2
    queenButton.x = startX + this.tileSize * 3

    knightButton.y = startY
    bishopButton.y = startY
    rookButton.y = startY
    queenButton.y = startY

    knightButton.visible = true
    bishopButton.visible = true
    rookButton.visible = true
    queenButton.visible = true


  }

  #updateShogiPromoteButtons(game: Game) {
    if (this.pendingMove === null) { return }

    const end = this.pendingMove.move.end
    const cancelButton = this.buttons.get("shogiPromoteCancel")!
    const promoteButton = this.buttons.get("shogiPromote")!

    let startY
    let startX = end.x * this.tileSize + this.tileSize / 2
    if (game.userSide === 0) {
      startY = end.y * this.tileSize
    } else {
      startY = (end.y + 2) * this.tileSize
    }

    if (startX < 100 * this.UIRatio) {
      startX = 100 * this.UIRatio
    }
    if (startX + 2 * this.tileSize > 900 * this.UIRatio) {
      startX = 700 * this.UIRatio
    }

    cancelButton.x = startX
    promoteButton.x = startX + this.tileSize

    cancelButton.y = startY
    promoteButton.y = startY

    cancelButton.visible = true
    promoteButton.visible = true

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

    for (let i = 0; i < game.height; i++) {
      for (let j = 0; j < game.width; j++) {
        const piece = game.board[i][j]
        if (piece === null) {
          continue
        }

        piece.moveUpdate(game, this.tileSize, input, sendMessage, this.setPendingMove)
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


  pressResign() {
    if (this.resignPressed) return
    this.resignPressed = true
  }

  cancelResign() {
    if (!this.resignPressed) return
    this.resignPressed = false
  }

  confirmResign(sendMessage: (msg: Message<unknown>) => void) {
    sendResignMessage(sendMessage)
  }

  setPendingMove(move: Move, type: number, game: Game) {
    const pendingMove: PendingMove = {
      move: move,
      type: type
    }

    this.pendingMove = pendingMove
    this.updateButtonScreen(game)
  }

  sendPendingMove(promoteType: number | null, game: Game, sendMessage: (msg: Message<unknown>) => void) {
    if (this.pendingMove === null) {
      return
    }

    game.turn = getEnemyTurnInt(game)
    const move = this.pendingMove.move
    move.Promote = promoteType

    const moveString = convertMoveToString(move, game.height)
    sendMoveMessage(moveString, sendMessage)
    this.pendingMove = null

    this.updateButtonScreen(game)
  }

  updateLastMove(game: Game, moveString: string) {
    this.lastMove = {
      start: null,
      end: null,
    }

    const move = moveString.split(",")
    if (move.length === 2) {
      const start = convertStringToPosition(move[0], game.height)
      const end = convertStringToPosition(move[1], game.height)
      this.lastMove = {
        start: start,
        end: end,
      }
    }
  }
}
