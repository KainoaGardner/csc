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
      <div className="flex flex-col items-start">
        <h1 className="font-bold text-8xl text-gray-50 mb-5">Join With Code</h1>

        <h2 className="font-bold text-3xl text-gray-50">Game Code</h2>
        <input
          className="textInput"
          name="code"
          value={code}
          onChange={handleChange}
          placeholder="Code"
        />

        <button
          className="hover:bg-neutral-600 mt-3.5 text-2xl py-2 px-4 text-gray-50 border-neutral-400 bg-neutral-700 border-5"
          onClick={getJoinablePrivateGame}
        >Submit</button>


        <hr className="border-none my-4" />
        <button
          className="btn w-2xl text-3xl"
          onClick={() => { setPage("joinGame") }}>Back</button>
      </div>


    </>
  );
}
export default PrivateJoin;
