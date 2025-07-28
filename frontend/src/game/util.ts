export const FenStringToPieceInt = new Map<string, number>()

FenStringToPieceInt.set("CP", 1)
FenStringToPieceInt.set("CN", 2)
FenStringToPieceInt.set("CB", 3)
FenStringToPieceInt.set("CR", 4)
FenStringToPieceInt.set("CQ", 5)
FenStringToPieceInt.set("CK", 6)
FenStringToPieceInt.set("SP", 7)
FenStringToPieceInt.set("SL", 8)
FenStringToPieceInt.set("SN", 9)
FenStringToPieceInt.set("SG", 10)
FenStringToPieceInt.set("SC", 11)
FenStringToPieceInt.set("SB", 12)
FenStringToPieceInt.set("SR", 13)
FenStringToPieceInt.set("SK", 14)
FenStringToPieceInt.set("NP", 15)
FenStringToPieceInt.set("NL", 16)
FenStringToPieceInt.set("NN", 17)
FenStringToPieceInt.set("NG", 18)
FenStringToPieceInt.set("NB", 19)
FenStringToPieceInt.set("NR", 20)
FenStringToPieceInt.set("KC", 21)
FenStringToPieceInt.set("KK", 22)

export interface Vec2 {
  x: number;
  y: number;
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
