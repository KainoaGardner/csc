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



export interface Vec2 {
  x: number;
  y: number;
}

export interface Mouse {
  x: number;
  y: number;
  pressed: boolean;
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



