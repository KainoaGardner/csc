import API_URL from "./env.tsx"
import { useState, useEffect } from "react"
import { useFetchWithAuth } from "./apiCalls/api.tsx"

import { useApp, useErrorHandler, useLogoutHandler } from "./appContext/useApp.tsx"

type UserData = {
  username: string;
  email: string;
  time: string;
}

type UserStats = {
  gamesPlayed: number;
  gamesWon: number;
  gameLog: string[];
}

const emptyUserData = {
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
  const { accessToken, setPage } = useApp()
  const [userData, setUserData] = useState<UserData>(emptyUserData)
  const [userStatsData, setUserStatsData] = useState<UserStats>(emptyUserStats)

  const fetchWithAuth = useFetchWithAuth()
  const { handleError } = useErrorHandler()
  const { handleLogout } = useLogoutHandler()

  async function getUser() {
    try {
      const response = await fetchWithAuth(API_URL + "user")
      const data = await response.json();
      if (response.ok) {
        const updatedData = {
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
        console.log(data)
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


  useEffect(() => {
    if (accessToken === null) {
      handleError("Not logged in")
      setPage("login")
    } else {
      getUser()
      getUserStats()
    }
  }, [])


  return (
    <>
      <h1>User</h1>

      <button onClick={handleLogout}>Logout</button>
      <hr />

      <h2>{userData.username}</h2>
      <h2>{userData.email}</h2>
      <h2>{userData.time}</h2>

      <h3>{userStatsData.gamesPlayed}</h3>
      <h3>{userStatsData.gamesWon}</h3>
      <h3>{userStatsData.gameLog}</h3>

    </>
  );
}
export default User;
