import { createContext } from "react"

interface AppContextType {
  page: Page
  setPage: (page: Page) => void
  error: string
  setError: (error: string) => void
  notif: string
  setNotif: (notif: string) => void
  accessToken: string | null;
  setAccessToken: (token: string | null) => void
  gameID: string | null;
  setGameID: (id: string | null) => void
  gameLogID: string | null;
  setGameLogID: (id: string | null) => void
  userID: string | null;
  setUserID: (id: string | null) => void
  volume: number;
  setVolume: (volume: number) => void
}

export type Page =
  "home" |
  "login" |
  "register" |
  "userStats" |
  "multiplayer" |
  "campaign" |
  "settings" |
  "createGame" |
  "joinGame" |
  "game" |
  "private" |
  "public" |
  "test" |
  "gameLog"

export const AppContext = createContext<AppContextType | undefined>(undefined)
