import API_URL from "../env.tsx"
import { useContext } from "react"
import { AppContext } from "./appContext.tsx"

import { Sounds, playAudio } from "../sounds.ts"

export const useApp = () => {
  const ctx = useContext(AppContext)
  if (!ctx) throw new Error("only useable in AppProvider")
  return ctx
}

export const useErrorHandler = () => {
  const { setError } = useApp();

  const handleError = (err: string) => {
    playAudio(Sounds.get("error")!)
    setError(err);
    setTimeout(() => { setError(""); }, 5000);
  }

  return { handleError }
}


export const useNotifHandler = () => {
  const { setNotif } = useApp();

  const handleNotif = (err: string) => {
    playAudio(Sounds.get("notif")!)
    setNotif(err);
    setTimeout(() => { setNotif(""); }, 5000);
  }

  return { handleNotif }
}


export const useLogoutHandler = () => {
  const { setAccessToken, setPage } = useApp();

  const handleLogout = () => {
    setAccessToken(null)
    fetch(API_URL + "user/logout", {
      method: "POST",
      credentials: "include"
    })
    setPage("login")
  }

  return { handleLogout }
}
