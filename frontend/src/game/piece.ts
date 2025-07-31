import { PieceImages, PieceImageDimensions } from "./images.ts"
import { type Vec2 } from "./util.ts"

export class Piece {
  type: number;
  owner: number;
  moved: boolean;
  imageName: string;
  image: HTMLImageElement | undefined;
  imageSize: Vec2 = { x: 0, y: 0 }

  constructor(type: number, owner: number, moved: boolean) {
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

  draw(ctx: CanvasRenderingContext2D, xTile: number, yTile: number, tileSize: number) {
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

    ctx.drawImage(
      this.image,
      xTile + tileSize / 2 - this.imageSize.x * imageRatio / 2,
      yTile + tileSize / 2 - this.imageSize.y * imageRatio / 2,
      this.imageSize.x * imageRatio,
      this.imageSize.y * imageRatio
    )
  }
}


