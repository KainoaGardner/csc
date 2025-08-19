import API_URL from "./env.tsx"
import { useFetchWithAuth } from "./apiCalls/api.tsx"
import { useApp, useErrorHandler, useNotifHandler } from "./appContext/useApp.tsx"
import { useState } from "react"

type FormData = {
  width: number;
  height: number;
  whiteMoney: number;
  blackMoney: number;
  whiteTime: number;
  blackTime: number;
  placeLine: number;
  public: boolean;
}

const emptyFormData = {
  width: 8,
  height: 8,
  whiteMoney: 300,
  blackMoney: 300,
  whiteTime: 600,
  blackTime: 600,
  placeLine: 4,
  public: true,
}

function CreateGame() {
  const { setPage, accessToken, setGameID } = useApp()
  const { handleError } = useErrorHandler()
  const { handleNotif } = useNotifHandler()

  const fetchWithAuth = useFetchWithAuth()

  const [gameConfig, setGameConfig] = useState<FormData>(emptyFormData);

  if (accessToken === null) {
    handleError("Not logged in")
    setPage("login")
  }

  const handleChange = (e: React.ChangeEvent<HTMLInputElement>) => {
    const { name, value } = e.target;
    setGameConfig((prev) => ({
      ...prev,
      [name]: parseInt(value),
    }))
  }

  const handleCheckBoxChange = (e: React.ChangeEvent<HTMLInputElement>) => {
    const { name, checked } = e.target;
    setGameConfig((prev) => ({
      ...prev,
      [name]: checked,
    }))
  }

  async function postCreateGame() {
    const bodyData = {
      width: gameConfig.width,
      height: gameConfig.height,
      money: [gameConfig.whiteMoney, gameConfig.blackMoney],
      startTime: [gameConfig.whiteTime, gameConfig.blackTime],
      placeLine: gameConfig.placeLine,
      public: gameConfig.public,
    }

    try {
      const response = await fetchWithAuth(API_URL + "game", {
        method: "POST",
        body: JSON.stringify(bodyData),
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
      <div className="flex flex-col ">
        <h1 className="font-bold text-8xl text-gray-50 mb-5">Create Game</h1>

        <div className="flex items-center">
          <label className="font-bold text-3xl text-gray-50 mr-5 mb-2.5">Width</label>
          <input
            className="textInput text-3xl"
            type="number"
            name="width"
            value={gameConfig.width}
            onChange={handleChange}
            placeholder="8"
            min={1}
            max={20}
          />

          <label className="font-bold text-3xl text-gray-50 mx-5 mb-2.5">Height</label>
          <input
            className="textInput text-3xl"
            type="number"
            name="height"
            value={gameConfig.height}
            onChange={handleChange}
            placeholder="8"
            min={1}
            max={20}
          />
        </div>

        <div className="flex items-center">
          <label className="font-bold text-3xl text-gray-50 mr-5 mb-2.5">White Money</label>
          <input
            className="textInput text-3xl"
            type="number"
            name="whiteMoney"
            value={gameConfig.whiteMoney}
            onChange={handleChange}
            placeholder="300"
            min={50}
            max={100000}
          />

          <label className="font-bold text-3xl text-gray-50 mx-5 mb-2.5">Black Money</label>
          <input
            className="textInput text-3xl"
            type="number"
            name="blackMoney"
            value={gameConfig.blackMoney}
            onChange={handleChange}
            placeholder="300"
            min={50}
            max={100000}
          />
        </div>

        <div className="flex items-center">
          <label className="font-bold text-3xl text-gray-50 mr-5 mb-2.5">White Time</label>
          <input
            className="textInput text-3xl"
            type="number"
            name="whiteTime"
            value={gameConfig.whiteTime}
            onChange={handleChange}
            placeholder="600"
            min={1}
            max={100000}
          />

          <label className="font-bold text-3xl text-gray-50 mx-5 mb-2.5">Black Time</label>
          <input
            className="textInput text-3xl"
            type="number"
            name="blackTime"
            value={gameConfig.blackTime}
            onChange={handleChange}
            placeholder="600"
            min={1}
            max={100000}
          />
        </div>

        <div className="flex items-center">
          <label className="font-bold text-3xl text-gray-50 mr-5 mb-2.5">Place Line</label>
          <input
            className="textInput text-3xl"
            type="number"
            name="placeLine"
            value={gameConfig.placeLine}
            onChange={handleChange}
            placeholder={(gameConfig.height / 2).toString()}
            min={1}
            max={gameConfig.height - 1}
          />
        </div>

        <div className="flex items-center">
          <label className="font-bold text-3xl text-gray-50 mr-5 mb-2.5">Public</label>
          <input
            className="p-2 w-10 h-20 text-3xl"
            type="checkbox"
            name="public"
            checked={gameConfig.public}
            onChange={handleCheckBoxChange}
          />
        </div>

        <button
          className="btn w-2xl text-3xl"
          onClick={postCreateGame}>Create</button>

        <hr className="border-none my-4" />
        <button
          className="btn w-2xl text-3xl"
          onClick={() => { setPage("multiplayer") }}>Back</button>


      </div>


    </>
  );
}
export default CreateGame;
