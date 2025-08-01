import { PieceImages, PieceImageDimensions } from "./images.ts"
import { PieceTypeToPrice } from "./util.ts"
import { type Vec2 } from "./util.ts"
import { InputHandler } from "./inputHandler.ts"
import { Game } from "./game.ts"

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

  placeUpdate(game: Game, tileSize: number, input: InputHandler) {
    const x = (this.x + 1) * tileSize
    const y = (this.y + 1) * tileSize

    const result = x <= input.mouse.x && input.mouse.x <= x + tileSize && y <= input.mouse.y && input.mouse.y <= y + tileSize
    if (result && input.mouse.justPressed) {
      this.selected = true
    }

    if (this.selected && input.mouse.justReleased) {
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
          } else {
            game.board[this.y][this.x] = null
            game.board[placeY][placeX] = this
            this.x = placeX
            this.y = placeY
          }
        }
      }

      if (!this.checkPieceOnBoard(placeX, placeY, game)) {
        if (!shopPiece) {
          const price = PieceTypeToPrice.get(this.type)
          if (price === undefined) { return }

          game.money[game.userSide] += price
          game.board[this.y][this.x] = null
        }
      }
    }
  }

  checkPieceOnBoard(x: number, y: number, game: Game): boolean {
    const result = 0 <= x && x < game.width && 0 <= y && y < game.height
    return result
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
}

// if ((game.userSide === 0 && placeY < game.placeLine) ||
//   (game.userSide === 1 && placeY >= game.placeLine)
// ) {
//   return
// }

