import API_URL from "./env.tsx"
import { useApp, useErrorHandler, useNotifHandler } from "./appContext/useApp.tsx"
import { useFetchWithAuth } from "./apiCalls/api.tsx"
import { useState, useEffect } from "react"

type JoinableGame = {
  id: string;
  width: number;
  height: number;
  placeLine: number;
  whiteID: string;
  whiteMoney: number;
  blackMoney: number;
  whiteTime: number;
  blackTime: number;
}

function PublicJoin() {
  const { setPage, accessToken, setGameID } = useApp()
  const { handleError } = useErrorHandler()
  const { handleNotif } = useNotifHandler()

  const fetchWithAuth = useFetchWithAuth()

  const [joinGames, setJoinGames] = useState<JoinableGame[]>([])

  if (accessToken === null) {
    handleError("Not logged in")
    setPage("login")
  }

  async function getJoinableGames() {
    try {
      const response = await fetchWithAuth(API_URL + "game/join/all")
      const data = await response.json();
      if (response.ok) {
        const result = []
        console.log(data)
        for (let i = 0; i < data.data.length; i++) {
          const joinableGameData = {
            id: data.data[i]._id,
            width: data.data[i].width,
            height: data.data[i].height,
            placeLine: data.data[i].placeLine,
            whiteID: data.data[i].whiteID,
            whiteMoney: data.data[i].money[0],
            blackMoney: data.data[i].money[1],
            whiteTime: data.data[i].time[0],
            blackTime: data.data[i].time[1],
          }
          result.push(joinableGameData)
        }
        setJoinGames(result)
      } else {
        handleError(data.error);
      }
    } catch (error) {
      console.log(error);
    }
  }

  async function postJoinGame(code: string) {
    try {
      const url = API_URL + "game/" + code + "/join"
      console.log(url)
      const response = await fetchWithAuth(API_URL + "game/" + code + "/join", {
        method: "POST",
      })

      const data = await response.json();
      if (response.ok) {
        handleNotif(data.message)
        setGameID(data.data._id)
        setPage("game")
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
      getJoinableGames()
    }
  }, [])


  return (
    <>
      <h1>Game List</h1>
      <button onClick={() => { setPage("joinGame") }}>Back</button>
      <hr />

      <ul>
        {joinGames.map(games => (
          <li key={games.id}>
            <p>Size: {games.width}x{games.height} Place Line: {games.placeLine}</p>
            <p>Money: {games.whiteMoney}/{games.blackMoney}</p>
            <p>Time: {games.whiteTime}/{games.blackTime}</p>
            <button onClick={() => postJoinGame(games.id)}>Join</button>
          </li>
        ))}
      </ul>

    </>
  );
}
export default PublicJoin;
