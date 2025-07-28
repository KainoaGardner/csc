import { createContext } from "react"

interface SettingContextType {
  boardTheme: number
  setBoardTheme: (boardTheme: number) => void
  pieceTheme: number
  setPieceTheme: (pieceTheme: number) => void
}


export const SettingContext = createContext<SettingContextType | undefined>(undefined)
