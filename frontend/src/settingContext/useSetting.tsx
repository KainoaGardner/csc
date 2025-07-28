import { useContext } from "react"
import { SettingContext } from "./settingContext.tsx"

export const useApp = () => {
  const ctx = useContext(SettingContext)
  if (!ctx) throw new Error("only useable in Setting Provider")
  return ctx
}
