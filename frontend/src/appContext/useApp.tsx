import { useContext } from "react"
import { AppContext } from "./appContext.tsx"

export const useApp = () => {
  const ctx = useContext(AppContext)
  if (!ctx) throw new Error("only useable in AppProvider")
  return ctx
}

export const useErrorHandler = () => {
  const { setError } = useApp();

  const handleError = (err: string) => {
    setError(err);
    setTimeout(() => { setError(""); }, 5000);
  }

  return { handleError }
}

export const useNotifHandler = () => {
  const { setNotif } = useApp();

  const handleNotif = (err: string) => {
    setNotif(err);
    setTimeout(() => { setNotif(""); }, 5000);
  }

  return { handleNotif }
}
