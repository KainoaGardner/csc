import { useState } from "react"
import type { ReactNode } from "react"
import { AppContext, type Page } from "./appContext.tsx"


export const AppProvider = ({ children }: { children: ReactNode }) => {
  const [page, setPage] = useState<Page>("home")
  const [error, setError] = useState<string>("")
  const [notif, setNotif] = useState<string>("")
  const [accessToken, setAccessToken] = useState<string | null>(null)
  const [gameID, setGameID] = useState<string | null>(null)
  const [userID, setUserID] = useState<string | null>(null)

  return (
    <AppContext.Provider value={{ page, setPage, error, setError, notif, setNotif, accessToken, setAccessToken, gameID, setGameID, userID, setUserID }}>
      {children}
    </AppContext.Provider>
  )
}
