import API_URL from "./env.tsx"
import { useState, useEffect } from "react"
import { useFetchWithAuth } from "./apiCalls/api.tsx"

import { useApp, useErrorHandler } from "./appContext/useApp.tsx"

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
      <div className="flex items-start">
        <div className="flex-none">
          <h1 className="font-bold text-8xl text-gray-50 mb-10">User</h1>

          <h2 className="font-bold text-4xl text-gray-50">Username</h2>
          <h2 className="font-bold text-3xl text-neutral-400 mb-4">{userData.username}</h2>

          <h2 className="font-bold text-4xl text-gray-50">Email</h2>
          <h2 className="font-bold text-3xl text-neutral-400 mb-4">{userData.email}</h2>

          <h2 className="font-bold text-4xl text-gray-50">Account Created</h2>
          <h2 className="font-bold text-3xl text-neutral-400 mb-4">{userData.time}</h2>

          <h2 className="font-bold text-4xl text-gray-50">Games Played</h2>
          <h2 className="font-bold text-4xl text-neutral-400 mb-4">{userStatsData.gamesPlayed}</h2>

          <h2 className="font-bold text-4xl text-gray-50">Games Won</h2>
          <h2 className="font-bold text-4xl text-neutral-400 mb-4">{userStatsData.gamesWon}</h2>
        </div>

        <div className="flex-1 flex justify-end">
          <div>
            <h1 className="font-bold text-8xl text-gray-50 mb-10 text-right">Recent Games</h1>
            <ul className="flex flex-col items-end">
              {gameLogHistoryData.map(gameLogs => (
                <li
                  className="bg-neutral-700 border-5 border-neutral-400 py-4 px-4 flex items-start justify-between my-4 w-200"
                  key={gameLogs.id}
                >
                  <div>
                    <h2 className=
                      {getWinnerCSS(checkWinner(gameLogs.whiteID, gameLogs.blackID, userData.userID, gameLogs.winner))}
                    >{checkWinner(gameLogs.whiteID, gameLogs.blackID, userData.userID, gameLogs.winner)}</h2>

                    <h2 className="font-bold text-3xl text-gray-50">Date</h2>
                    <h2 className="font-bold text-2xl text-neutral-400 mb-4">{new Date(gameLogs.date).toLocaleString()}</h2>

                    <h2 className="font-bold text-3xl text-gray-50">Move Count</h2>
                    <h2 className="font-bold text-3xl text-neutral-400 mb-4">{gameLogs.moveCount}</h2>

                    {/* <h2 className="font-bold text-3xl text-gray-50">Winner</h2> */}



                    <h2 className="font-bold text-3xl text-gray-50">Reason</h2>
                    <h2 className="font-bold text-3xl text-neutral-400 mb-4">{gameLogs.reason}</h2>
                  </div>

                  <button
                    className="hover:bg-neutral-600 mt-0.5 text-2xl py-2 px-4 text-gray-50 border-neutral-400 bg-neutral-700 border-4"
                    onClick={() => handleViewGameLog(gameLogs.id)}>View</button>
                </li>
              ))}
            </ul>
          </div>
        </div>
      </div>
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

function getWinnerCSS(winner: string): string {
  if (winner === "Win") {
    return "font-bold text-4xl text-green-300 mb-2"
  }

  if (winner === "Loss") {
    return "font-bold text-4xl text-red-300 mb-2"
  }

  return "font-bold text-4xl text-yellow-300 mb-2"
}


export default User;
