import { useState } from "react"
import type { ReactNode } from "react"
import { AppContext, type Page } from "./appContext.tsx"


export const AppProvider = ({ children }: { children: ReactNode }) => {
  const [page, setPage] = useState<Page>("home")
  const [error, setError] = useState<string>("")
  const [notif, setNotif] = useState<string>("")
  const [accessToken, setAccessToken] = useState<string | null>(null)

  return (
    <AppContext.Provider value={{ page, setPage, error, setError, notif, setNotif, accessToken, setAccessToken }}>
      {children}
    </AppContext.Provider>
  )
}
