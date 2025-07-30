import { type Vec2 } from "./util.ts"

export const PieceImages = new Map<string, HTMLImageElement>()
export const PieceImageDimensions = new Map<string, Vec2>()

for (let i = 1; i <= 22; i++) {
  const whiteString = `w${i}`
  const blackString = `b${i}`
  const whitePiece = new Image();
  if (whitePiece instanceof HTMLElement) {
    whitePiece.src = `/images/pieces/${whiteString}.png`;
    PieceImages.set(whiteString, whitePiece)
  }
  const blackPiece = new Image();
  if (blackPiece instanceof HTMLElement) {
    blackPiece.src = `/images/pieces/${blackString}.png`;
    PieceImages.set(blackString, blackPiece)
  }
}




