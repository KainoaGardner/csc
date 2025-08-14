import { PieceImages, PieceImageDimensions } from "./images.ts"
import { PieceTypeToPrice, PlaceEnum, type Message, sendPlaceMessage, sendMoveMessage, convertPositionToString } from "./util.ts"
import { type Vec2, PieceEnum, convertMoveToString, PromoteTypeEnum, type Move } from "./util.ts"
import { InputHandler } from "./inputHandler.ts"
import { Game } from "./game.ts"
import { checkValidPieceMove, checkPieceOnBoard, checkValidDropMove, getEnemyTurnInt, checkPromote, checkMustPromote } from "./engine.ts"

export class Piece {
  x: number
  y: number

  type: number;
  owner: number;
  moved: boolean;
  imageName: string;
  image: HTMLImageElement | undefined;
  imageSize: Vec2 = { x: 0, y: 0 }

  selected: boolean = false

  constructor(x: number, y: number, type: number, owner: number, moved: boolean) {
    this.x = x
    this.y = y
    this.type = type
    this.owner = owner
    this.moved = moved

    let imageName = ""
    if (this.owner === 0) {
      imageName = "w"
    } else {
      imageName = "b"
    }

    imageName += this.type.toString()
    this.imageName = imageName
    this.image = PieceImages.get(imageName)
  }

  draw(ctx: CanvasRenderingContext2D, tileSize: number, input: InputHandler) {
    if (this.image === undefined) {
      return
    }

    if (this.imageSize.x === 0 || this.imageSize.y === 0) {
      const imageSize = PieceImageDimensions.get(this.imageName)
      if (imageSize === undefined) {
        const dimensions = { x: this.image.naturalWidth, y: this.image.naturalHeight }
        PieceImageDimensions.set(this.imageName, dimensions)
        this.imageSize = dimensions
      } else {
        this.imageSize = imageSize
      }
    }

    let imageRatio = tileSize / this.imageSize.y * 0.9
    if (this.type >= 0 && this.type <= 6) {
      imageRatio *= 0.9
    }

    if (this.selected) {
      ctx.drawImage(this.image, input.mouse.x - this.imageSize.x * imageRatio / 2, input.mouse.y - this.imageSize.y * imageRatio / 2, this.imageSize.x * imageRatio, this.imageSize.y * imageRatio)
    } else {
      ctx.drawImage(
        this.image,
        (this.x + 1) * tileSize + tileSize / 2 - this.imageSize.x * imageRatio / 2,
        (this.y + 1) * tileSize + tileSize / 2 - this.imageSize.y * imageRatio / 2,
        this.imageSize.x * imageRatio,
        this.imageSize.y * imageRatio
      )
    }
  }

  placeUpdate(game: Game, tileSize: number, input: InputHandler, sendMessage: (msg: Message<unknown>) => void) {
    if (game.ready[game.userSide]) return

    if (this.checkHovering(tileSize, input) && input.mouse.justPressed[0] && game.userSide === this.owner) {
      this.selected = true
    }

    if (this.selected && input.mouse.justReleased[0]) {
      this.selected = false
      const placeX = Math.floor(input.mouse.x / tileSize) - 1
      const placeY = Math.floor(input.mouse.y / tileSize) - 1

      const shopPiece = !this.checkPieceOnBoardSide(this.x, this.y, game)

      if (this.checkPieceOnBoardSide(placeX, placeY, game)) {
        if (game.board[placeY][placeX] === null) {
          if (shopPiece) {
            const price = PieceTypeToPrice.get(this.type)
            if (price === undefined) { return }

            if (game.money[game.userSide] - price < 0) {
              console.log("Out of money")
              return
            }

            game.money[game.userSide] -= price

            const placePiece = new Piece(placeX, placeY, this.type, this.owner, false)
            game.board[placeY][placeX] = placePiece

            const positionString = convertPositionToString({ x: placeX, y: placeY }, game.height)
            sendPlaceMessage(positionString, "", this.type, PlaceEnum.create, sendMessage)

          } else {
            const positionString = convertPositionToString({ x: placeX, y: placeY }, game.height)
            const fromString = convertPositionToString({ x: this.x, y: this.y }, game.height)

            game.board[this.y][this.x] = null
            game.board[placeY][placeX] = this
            this.x = placeX
            this.y = placeY

            sendPlaceMessage(positionString, fromString, this.type, PlaceEnum.move, sendMessage)
          }
        }
      }

      if (!checkPieceOnBoard(placeX, placeY, game)) {
        if (!shopPiece) {
          const price = PieceTypeToPrice.get(this.type)
          if (price === undefined) { return }

          game.money[game.userSide] += price
          game.board[this.y][this.x] = null

          const positionString = convertPositionToString({ x: this.x, y: this.y }, game.height)
          sendPlaceMessage(positionString, "", this.type, PlaceEnum.delete, sendMessage)
        }
      }
    }
  }


  moveUpdate(
    game: Game,
    tileSize: number,
    input: InputHandler,
    sendMessage: (msg: Message<unknown>) => void,
    setPendingMove: (move: Move, type: number, game: Game) => void
  ) {
    if (this.checkHovering(tileSize, input) && input.mouse.justPressed[0] && game.userSide == this.owner) {
      this.selected = true
    }

    if (this.selected && input.mouse.justReleased[0]) {
      this.selected = false
      const placeX = Math.floor(input.mouse.x / tileSize) - 1
      const placeY = Math.floor(input.mouse.y / tileSize) - 1

      const start = { x: this.x, y: this.y }
      const end = { x: placeX, y: placeY }

      if (checkValidPieceMove(start, end, this, game) && game.turn === game.userSide) {
        //send Move Message

        const promoteType = checkPromote(start, end, this, game.height)
        if (promoteType !== null) {
          const mustPromote = checkMustPromote(end, this, game.height)

          if (mustPromote) {
            if (this.type === PieceEnum.Checker || (this.type >= PieceEnum.Fu && this.type <= PieceEnum.Hi)) { //shogi checkers must promote no choice
              const move: Move = {
                start: { x: this.x, y: this.y },
                end: { x: placeX, y: placeY },
                Drop: null,
                Promote: 0,
              }
              const moveString = convertMoveToString(move, game.height)
              sendMoveMessage(moveString, sendMessage)
              game.turn = getEnemyTurnInt(game)
            } else { //pawn chess must promote with choice
              const move: Move = {
                start: { x: this.x, y: this.y },
                end: { x: placeX, y: placeY },
                Drop: null,
                Promote: null,
              }

              setPendingMove(move, PromoteTypeEnum.chess, game)
            }
          } else { //shogi promote not required can choose
            const move: Move = {
              start: { x: this.x, y: this.y },
              end: { x: placeX, y: placeY },
              Drop: null,
              Promote: null,
            }

            setPendingMove(move, PromoteTypeEnum.shogi, game)
          }
        } else {
          const move: Move = {
            start: { x: this.x, y: this.y },
            end: { x: placeX, y: placeY },
            Drop: null,
            Promote: null,
          }

          const moveString = convertMoveToString(move, game.height)
          sendMoveMessage(moveString, sendMessage)
          game.turn = getEnemyTurnInt(game)
        }
        game.board[this.y][this.x] = null
        game.board[placeY][placeX] = this
        this.x = placeX
        this.y = placeY

      }
    }
  }

  moveMochigomaUpdate(game: Game, tileSize: number, input: InputHandler, sendMessage: (msg: Message<unknown>) => void) {
    if (this.checkHovering(tileSize, input) && input.mouse.justPressed[0] && game.userSide == this.owner) {
      this.selected = true
    }

    if (this.selected && input.mouse.justReleased[0]) {
      this.selected = false
      const placeX = Math.floor(input.mouse.x / tileSize) - 1
      const placeY = Math.floor(input.mouse.y / tileSize) - 1

      const start = { x: this.x, y: this.y }
      const end = { x: placeX, y: placeY }

      if (checkValidDropMove(end, this, game) && game.turn === game.userSide) {
        const move: Move = {
          start: start,
          end: end,
          Drop: this.type - PieceEnum.Fu,
          Promote: null,
        }

        const moveString = convertMoveToString(move, game.height)
        sendMoveMessage(moveString, sendMessage)

        const mochigomaIndex = game.userSide * 7 + this.type - PieceEnum.Fu
        game.mochigoma[mochigomaIndex]--

        const dropPiece = new Piece(placeX, placeY, this.type, this.owner, false)
        game.board[placeY][placeX] = dropPiece
        game.turn = getEnemyTurnInt(game)
      }
    }
  }

  checkValidMove(): boolean {
    return false
  }


  checkPieceOnBoardSide(x: number, y: number, game: Game): boolean {
    if (game.userSide === 0) {
      const result = 0 <= x && x < game.width && game.placeLine <= y && y < game.height
      return result
    } else {
      const result = 0 <= x && x < game.width && 0 <= y && y < game.placeLine
      return result
    }
  }

  checkHovering(tileSize: number, input: InputHandler) {
    const x = (this.x + 1) * tileSize
    const y = (this.y + 1) * tileSize

    return x <= input.mouse.x && input.mouse.x <= x + tileSize && y <= input.mouse.y && input.mouse.y <= y + tileSize
  }

}


