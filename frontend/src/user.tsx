import API_URL from "./env.tsx"
import { useState, useEffect } from "react"
import { useFetchWithAuth } from "./apiCalls/api.tsx"

import { useApp, useErrorHandler, useLogoutHandler } from "./appContext/useApp.tsx"

import { WinnerEnum } from "./game/util.ts"

type UserData = {
  userID: string;
  username: string;
  email: string;
  time: string;
}

type UserStats = {
  gamesPlayed: number;
  gamesWon: number;
  gameLog: string[];
}

type GameLogHistory = {
  id: string;
  date: Date;
  moveCount: number;
  winner: number;
  reason: string;

  whiteID: string;
  blackID: string;
}

const emptyUserData = {
  userID: "",
  username: "",
  email: "",
  time: "",
}

const emptyUserStats = {
  gamesPlayed: 0,
  gamesWon: 0,
  gameLog: [],
}

function User() {
  const { accessToken, setPage, setGameLogID } = useApp()
  const [userData, setUserData] = useState<UserData>(emptyUserData)
  const [userStatsData, setUserStatsData] = useState<UserStats>(emptyUserStats)
  const [gameLogHistoryData, setGameLogHistoryData] = useState<GameLogHistory[]>([])

  const fetchWithAuth = useFetchWithAuth()
  const { handleError } = useErrorHandler()
  const { handleLogout } = useLogoutHandler()

  async function getUser() {
    try {
      const response = await fetchWithAuth(API_URL + "user")
      const data = await response.json();
      if (response.ok) {
        const updatedData = {
          userID: data.data._id,
          username: data.data.username,
          email: data.data.email,
          time: data.data.createdTime.slice(0, 10),
        }
        setUserData(updatedData)

      } else {
        handleError(data.error);
      }
    } catch (error) {
      console.log(error);
    }
  }

  async function getUserStats() {
    try {
      const response = await fetchWithAuth(API_URL + "userStat")
      const data = await response.json();
      if (response.ok) {
        const updatedStats = {
          gamesPlayed: data.data.gamesPlayed,
          gamesWon: data.data.gamesWon,
          gameLog: data.data.gameLog,
        }
        setUserStatsData(updatedStats)

      } else {
        handleError(data.error);
      }
    } catch (error) {
      console.log(error);
    }
  }

  async function getGameLogsHistory() {
    try {
      const response = await fetchWithAuth(API_URL + "userStat/gameLogs")
      const data = await response.json();
      if (response.ok) {
        const result = []
        for (let i = 0; i < data.data.length; i++) {
          const gameLogHistory = {
            id: data.data[i]._id,
            date: data.data[i].date,
            moveCount: data.data[i].moveCount,
            winner: data.data[i].winner,
            reason: data.data[i].reason,
            whiteID: data.data[i].whiteID,
            blackID: data.data[i].blackID,
          }
          result.push(gameLogHistory)
        }
        setGameLogHistoryData(result)
      } else {
        handleError(data.error);
      }
    } catch (error) {
      console.log(error);
    }
  }

  const handleViewGameLog = (gameLogID: string) => {
    setGameLogID(gameLogID)
    setPage("gameLog")
  }


  useEffect(() => {
    if (accessToken === null) {
      handleError("Not logged in")
      setPage("login")
    } else {
      getUser()
      getUserStats()
      getGameLogsHistory()
    }
  }, [])


  return (
    <>
      <h1>User</h1>

      <button onClick={handleLogout}>Logout</button>
      <hr />

      <h2>Username: {userData.username}</h2>
      <h2>Email: {userData.email}</h2>
      <h2>Account Created: {userData.time}</h2>

      <h3>Games Played: {userStatsData.gamesPlayed}</h3>
      <h3>Games Won: {userStatsData.gamesWon}</h3>

      <hr />
      <ul>
        {gameLogHistoryData.map(gameLogs => (
          <li key={gameLogs.id}>
            <p>Date: {gameLogs.date.toString()}</p>
            <p>Move count: {gameLogs.moveCount}</p>
            <p>Winner: {checkWinner(gameLogs.whiteID, gameLogs.blackID, userData.userID, gameLogs.winner)}</p>
            <p>Reason: {gameLogs.reason}</p>
            <button onClick={() => handleViewGameLog(gameLogs.id)}>Join</button>
          </li>
        ))}
      </ul>
    </>
  );
}

function checkWinner(whiteID: string, blackID: string, userID: string, winner: number): string {
  //tie
  if (winner === WinnerEnum.tie) {
    return "Draw"
  }

  if (winner === WinnerEnum.white && userID === whiteID) {
    return "Win"
  }

  if (winner === WinnerEnum.black && userID === blackID) {
    return "Win"
  }

  return "Loss"
}


export default User;
