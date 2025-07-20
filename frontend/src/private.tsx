import API_URL from "./env.tsx"
import { useFetchWithAuth } from "./apiCalls/api.tsx"
import { useApp, useErrorHandler, useNotifHandler } from "./appContext/useApp.tsx"
import { useState } from "react"

function PrivateJoin() {
  const { setPage, accessToken, setGameID } = useApp()
  const { handleError } = useErrorHandler()
  const { handleNotif } = useNotifHandler()
  const fetchWithAuth = useFetchWithAuth()
  const [code, setCode] = useState<string>("")

  if (accessToken === null) {
    handleError("Not logged in")
    setPage("login")
  }

  const handleChange = (e: React.ChangeEvent<HTMLInputElement>) => {
    const { value } = e.target;
    setCode(value)
  }

  async function postJoinGame() {
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

  return (
    <>
      <h1>Join with code</h1>
      <button onClick={() => { setPage("joinGame") }}>Back</button>
      <hr />

      <input
        name="code"
        value={code}
        onChange={handleChange}
        placeholder="Code"
      />

      <button
        onClick={postJoinGame}
      >Submit</button>


    </>
  );
}
export default PrivateJoin;
