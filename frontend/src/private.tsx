import API_URL from "./env.tsx"
import { useApp, useErrorHandler, useNotifHandler } from "./appContext/useApp.tsx"
import { useState } from "react"

function PrivateJoin() {
  const { setPage, accessToken, setGameID } = useApp()
  const { handleError } = useErrorHandler()
  const { handleNotif } = useNotifHandler()
  const [code, setCode] = useState<string>("")

  if (accessToken === null) {
    handleError("Not logged in")
    setPage("login")
  }

  const handleChange = (e: React.ChangeEvent<HTMLInputElement>) => {
    const { value } = e.target;
    setCode(value)
  }

  async function getJoinablePrivateGame() {
    try {
      const response = await fetch(API_URL + "game/" + code + "/private", {
        method: "GET",
      })

      const data = await response.json();
      if (response.ok) {
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
        onClick={getJoinablePrivateGame}
      >Submit</button>


    </>
  );
}
export default PrivateJoin;
