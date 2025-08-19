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


  const handleJoinGame = (gameID: string) => {
    setGameID(gameID)
    setPage("game")
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

      <div className="flex flex-col items-start">
        <h1 className="font-bold text-8xl text-gray-50 mb-5">Game List</h1>

        <hr className="border-none my-4" />
        <button
          className="btn w-2xl text-3xl"
          onClick={() => { setPage("joinGame") }}>Back</button>

        <hr className="border-none my-4" />

        <ul className="flex flex-col items-end">
          {joinGames.map(games => (
            <li
              className="bg-neutral-700 border-5 border-neutral-400 py-4 px-4 flex items-start justify-between my-4 w-200"
              key={games.id}>

              <div>
                <h2 className="font-bold text-3xl text-gray-50">Board Size</h2>
                <h2 className="font-bold text-2xl text-neutral-400 mb-4">{games.placeLine}</h2>

                <h2 className="font-bold text-3xl text-gray-50">Money</h2>
                <h2 className="font-bold text-2xl text-neutral-400 mb-4">${games.whiteMoney}/${games.blackMoney}</h2>

                <h2 className="font-bold text-3xl text-gray-50">Time</h2>
                <h2 className="font-bold text-2xl text-neutral-400 mb-4">{games.whiteTime / 1000}/{games.blackTime / 1000}</h2>
              </div>

              <button
                className="hover:bg-neutral-600 mt-0.5 text-2xl py-2 px-4 text-gray-50 border-neutral-400 bg-neutral-700 border-4"
                onClick={() => handleJoinGame(games.id)}>Join</button>
            </li>
          ))}
        </ul>

      </div>


    </>
  );
}
export default PublicJoin;
