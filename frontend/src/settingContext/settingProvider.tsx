import { useState } from "react"
import type { ReactNode } from "react"
import { SettingContext } from "./settingContext.tsx"


export const SettingProvider = ({ children }: { children: ReactNode }) => {
  const [boardTheme, setBoardTheme] = useState<number>(0)
  const [pieceTheme, setPieceTheme] = useState<number>(0)

  return (
    <SettingContext.Provider value={{ boardTheme, setBoardTheme, pieceTheme, setPieceTheme }}>
      {children}
    </SettingContext.Provider>
  )
}
