export const PieceEnum = {
  Pawn: 1,
  Knight: 2,
  Bishop: 3,
  Rook: 4,
  Queen: 5,
  King: 6,
  Fu: 7,
  Kyou: 8,
  Kei: 9,
  Gin: 10,
  Kin: 11,
  Kaku: 12,
  Hi: 13,
  Ou: 14,
  To: 15,
  NariKyou: 16,
  NariKei: 17,
  NariGin: 18,
  Uma: 19,
  Ryuu: 20,
  Checker: 21,
  CheckerKing: 22,
}

export const MochigomaEnum = {
  MochiFu: 0,
  MochiKyou: 1,
  MochiKei: 2,
  MochiGin: 3,
  MochiKin: 4,
  MochiKaku: 5,
  MochiHi: 6,
}

export const PlaceEnum = {
  create: 0,
  delete: 1,
  move: 2,
}

export const PromoteTypeEnum = {
  chess: 0,
  shogi: 1,
  checkers: 2,
}

export const FenStringToPieceInt = new Map<string, number>()
FenStringToPieceInt.set("CP", PieceEnum.Pawn)
FenStringToPieceInt.set("CN", PieceEnum.Knight)
FenStringToPieceInt.set("CB", PieceEnum.Bishop)
FenStringToPieceInt.set("CR", PieceEnum.Rook)
FenStringToPieceInt.set("CQ", PieceEnum.Queen)
FenStringToPieceInt.set("CK", PieceEnum.King)
FenStringToPieceInt.set("SP", PieceEnum.Fu)
FenStringToPieceInt.set("SL", PieceEnum.Kyou)
FenStringToPieceInt.set("SN", PieceEnum.Kei)
FenStringToPieceInt.set("SG", PieceEnum.Gin)
FenStringToPieceInt.set("SC", PieceEnum.Kin)
FenStringToPieceInt.set("SB", PieceEnum.Kaku)
FenStringToPieceInt.set("SR", PieceEnum.Hi)
FenStringToPieceInt.set("SK", PieceEnum.Ou)
FenStringToPieceInt.set("NP", PieceEnum.To)
FenStringToPieceInt.set("NL", PieceEnum.NariKyou)
FenStringToPieceInt.set("NN", PieceEnum.NariKei)
FenStringToPieceInt.set("NG", PieceEnum.NariGin)
FenStringToPieceInt.set("NB", PieceEnum.Uma)
FenStringToPieceInt.set("NR", PieceEnum.Ryuu)
FenStringToPieceInt.set("KC", PieceEnum.Checker)
FenStringToPieceInt.set("KK", PieceEnum.CheckerKing)

export const ShogiMochiPieceToChar = new Map<number, string>()
ShogiMochiPieceToChar.set(MochigomaEnum.MochiFu, "P")
ShogiMochiPieceToChar.set(MochigomaEnum.MochiKyou, "L")
ShogiMochiPieceToChar.set(MochigomaEnum.MochiKei, "N")
ShogiMochiPieceToChar.set(MochigomaEnum.MochiGin, "S")
ShogiMochiPieceToChar.set(MochigomaEnum.MochiKin, "G")
ShogiMochiPieceToChar.set(MochigomaEnum.MochiKaku, "B")
ShogiMochiPieceToChar.set(MochigomaEnum.MochiHi, "R")

export const ChessPromotePieceToChar = new Map<number, string>()
ChessPromotePieceToChar.set(PieceEnum.Pawn, "P")
ChessPromotePieceToChar.set(PieceEnum.Knight, "N")
ChessPromotePieceToChar.set(PieceEnum.Bishop, "B")
ChessPromotePieceToChar.set(PieceEnum.Rook, "R")
ChessPromotePieceToChar.set(PieceEnum.Queen, "Q")

export const PieceTypeToPrice = new Map<number, number>()
PieceTypeToPrice.set(PieceEnum.Pawn, 3)
PieceTypeToPrice.set(PieceEnum.Knight, 20)
PieceTypeToPrice.set(PieceEnum.Bishop, 25)
PieceTypeToPrice.set(PieceEnum.Rook, 30)
PieceTypeToPrice.set(PieceEnum.Queen, 50)
PieceTypeToPrice.set(PieceEnum.King, 50)
PieceTypeToPrice.set(PieceEnum.Fu, 3)
PieceTypeToPrice.set(PieceEnum.Kyou, 8)
PieceTypeToPrice.set(PieceEnum.Kei, 12)
PieceTypeToPrice.set(PieceEnum.Gin, 15)
PieceTypeToPrice.set(PieceEnum.Kin, 20)
PieceTypeToPrice.set(PieceEnum.Kaku, 28)
PieceTypeToPrice.set(PieceEnum.Hi, 35)
PieceTypeToPrice.set(PieceEnum.Ou, 45)
PieceTypeToPrice.set(PieceEnum.Checker, 10)

export type Message<T = unknown> = {
  type: string;
  data: T;
}

export type PlaceMessage = {
  position: string,
  fromPosition: string,
  type: number,
  place: number,
}

export type ReadyMessage = {
  ready: boolean,
}

export type MoveMessage = {
  move: string,
}

export type DrawMessage = {
  draw: boolean,
}

export interface Vec2 {
  x: number;
  y: number;
}

export interface Move {
  start: Vec2,
  end: Vec2,
  Drop: number | null
  Promote: number | null
}

export interface PendingMove {
  move: Move,
  type: number,
}


export interface Annotation {
  start: Vec2 | null
  end: Vec2 | null
}

export interface Mouse {
  x: number;
  y: number;
  pressed: boolean[];
  prevPressed: boolean[];
  justPressed: boolean[];
  justReleased: boolean[];
}

export function sendPlaceMessage(position: string, from: string, type: number, place: number, sendMessage: (msg: Message<unknown>) => void) {
  const placeRequest: Message<PlaceMessage> = {
    type: "place",
    data: {
      position: position,
      fromPosition: from,
      type: type,
      place: place,
    },
  }

  sendMessage(placeRequest)
}

export function sendMoveMessage(move: string, sendMessage: (msg: Message<unknown>) => void) {
  const moveRequest: Message<MoveMessage> = {
    type: "move",
    data: {
      move: move,
    },
  }

  sendMessage(moveRequest)
}




export function sendDrawMessage(draw: boolean, sendMessage: (msg: Message<unknown>) => void) {
  const drawRequest: Message<DrawMessage> = {
    type: "draw",
    data: {
      draw: draw,
    },
  }

  sendMessage(drawRequest)
}

export function sendResignMessage(sendMessage: (msg: Message<unknown>) => void) {
  const resignRequest: Message<null> = {
    type: "resign",
    data: null,
  }

  sendMessage(resignRequest)
}



export function isCharDigit(char: string): boolean {
  if (char.length !== 1) {
    return false
  }

  const c = char[0]
  return /[0-9]/.test(c)
}

export function isCharUppercase(char: string): boolean {
  if (char.length !== 1) {
    return false
  }

  const c = char[0]
  return /[A-Z]/.test(c)
}

export function isCharLowercase(char: string): boolean {
  if (char.length !== 1) {
    return false
  }

  const c = char[0]
  return /[a-z]/.test(c)
}

export function convertPositionToString(pos: Vec2, boardHeight: number): string {
  let result = ""

  const startWidthStr = convertNumberToLowercase(pos.x + 1)

  const startHeightStr = (boardHeight - pos.y).toString()

  result = startWidthStr + startHeightStr

  return result
}

export function convertStringToPosition(posString: string, boardHeight: number): Vec2 {
  const result: Vec2 = { x: 0, y: 0 }

  let xStr = ""
  let yStr = ""

  for (let i = 0; i < posString.length; i++) {
    const c = posString[i]
    if (isCharDigit(c)) {
      yStr += c
    } else if (isCharLowercase(c)) {
      xStr += c
    }
  }

  const x = convertLowercaseToNumber(xStr) - 1
  const y = parseInt(yStr)

  result.x = x
  result.y = boardHeight - y

  return result
}

export function convertNumberToLowercase(x: number): string {
  let result = ""

  while (x > 0) {
    const amount = (x - 1) % 26
    x = Math.floor((x - 1) / 26)
    const char = String.fromCharCode(amount + 97)
    result += char
  }

  return result
}

export function convertLowercaseToNumber(str: string): number {
  let result = 0

  for (let i = str.length - 1; i >= 0; i--) {
    if (!isCharLowercase(str[i])) {
      return -1
    }

    const charCode = str[i].codePointAt(0)
    if (!charCode)
      return -1

    const amount = charCode - 97 + 1
    const y = str.length - 1 - i
    result += amount * Math.pow(26, y)
  }

  return result
}


export function fitTextToWidth(ctx: CanvasRenderingContext2D, text: string, maxWidth: number, maxFontSize: number, minFontSize: number): number {
  let currFontSize = maxFontSize

  while (currFontSize >= minFontSize) {
    ctx.font = `${currFontSize}px Arial Black`
    const textWidth = ctx.measureText(text).width
    if (textWidth <= maxWidth) {
      break
    }
    currFontSize--
  }

  return currFontSize
}


export function convertSecondsToTimeString(time: number): string {
  const sec = Math.floor(time % 60)
  time = Math.floor(time / 60)
  const min = Math.floor(time % 60)
  time = Math.floor(time / 60)
  const hour = Math.floor(time % 60)

  let secString = sec.toString()
  let minString = min.toString()
  let hourString = hour.toString()
  if (sec === 0) {
    secString = "00"
  }
  if (min === 0) {
    minString = "00"
  }
  if (hour === 0) {
    hourString = "00"
  }

  if (hour === 0 && min === 0) {
    return secString
  } else if (hour === 0) {
    return `${minString}:${secString}`
  } else {
    return `${hourString}:${minString}:${secString}`
  }
}

export function checkVec2Equal(a: Vec2, b: Vec2): boolean {
  return (a.x === b.x && a.y === b.y)
}


export function checkEqualAnnotation(a: Annotation, b: Annotation): boolean {
  if (a.start === null || a.end === null || b.start === null || b.end === null) {
    return false
  }

  return checkVec2Equal(a.start, b.start) && checkVec2Equal(a.end, b.end)
}


export const AnnotationEnum = {
  singleSpace: 0,
  straightArrow: 1,
  diagonalArrow: 2,
  turnArrow: 3,
}


export function getAnnotationType(anno: Annotation): number {
  if (anno.start === null || anno.end === null) {
    return -1
  }

  //single spacce
  if (checkVec2Equal(anno.start, anno.end)) {
    return AnnotationEnum.singleSpace
  }

  //straightArrow
  if (anno.start.x === anno.end.x || anno.start.y === anno.end.y) {
    return AnnotationEnum.straightArrow
  }

  //turnArrow
  const dx = Math.abs(anno.start.x - anno.end.x)
  const dy = Math.abs(anno.start.y - anno.end.y)

  if (dx === 1 && dy === 2 || dx === 2 && dy === 1) {
    return AnnotationEnum.turnArrow
  }

  return AnnotationEnum.diagonalArrow
}

export function convertMoveToString(move: Move, boardHeight: number): string {
  let result = ""

  const startStr = convertStartMoveToString(move, boardHeight)

  result += startStr

  const endStr = convertEndMoveToString(move, boardHeight)

  result += "," + endStr

  return result
}

function convertStartMoveToString(move: Move, boardHeight: number): string {
  let result = convertPositionToString(move.start, boardHeight)

  if (move.Drop !== null) {
    const dropPiece = move.Drop
    const pieceChar = ShogiMochiPieceToChar.get(dropPiece)
    if (pieceChar === undefined) return ""

    result = pieceChar + "*"
  }

  return result
}

function convertEndMoveToString(move: Move, boardHeight: number): string {
  let result = ""

  const endWidthStr = convertNumberToLowercase(move.end.x + 1)

  const endHeightStr = (boardHeight - move.end.y).toString()
  result = endWidthStr + endHeightStr

  if (move.Promote !== null) {
    if (move.Promote !== 0) {
      const promotePiece = ChessPromotePieceToChar.get(move.Promote)
      if (promotePiece === undefined) { return "" }

      result += promotePiece
    } else {
      result += "+"
    }
  }

  return result
}
