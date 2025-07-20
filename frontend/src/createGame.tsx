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
    console.log(gameConfig)
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
      <h1>Create Game</h1>
      <button onClick={() => { setPage("multiplayer") }}>Back</button>
      <hr />

      <div>
        <label>Width</label>
        <input
          type="number"
          name="width"
          value={gameConfig.width}
          onChange={handleChange}
          placeholder="8"
          min={1}
          max={20}
        />

        <label>Height</label>
        <input
          type="number"
          name="height"
          value={gameConfig.height}
          onChange={handleChange}
          placeholder="8"
          min={1}
          max={20}
        />
      </div>

      <div>
        <label>White Money</label>
        <input
          type="number"
          name="whiteMoney"
          value={gameConfig.whiteMoney}
          onChange={handleChange}
          placeholder="300"
          min={50}
          max={100000}
        />

        <label>Black Money</label>
        <input
          type="number"
          name="blackMoney"
          value={gameConfig.blackMoney}
          onChange={handleChange}
          placeholder="300"
          min={50}
          max={100000}
        />
      </div>

      <div>
        <label>White Time</label>
        <input
          type="number"
          name="whiteTime"
          value={gameConfig.whiteTime}
          onChange={handleChange}
          placeholder="600"
          min={1}
          max={100000}
        />

        <label>Black Time</label>
        <input
          type="number"
          name="blackTime"
          value={gameConfig.blackTime}
          onChange={handleChange}
          placeholder="600"
          min={1}
          max={100000}
        />
      </div>

      <div>
        <label>Place Line</label>
        <input
          type="number"
          name="placeLine"
          value={gameConfig.placeLine}
          onChange={handleChange}
          placeholder={(gameConfig.height / 2).toString()}
          min={1}
          max={gameConfig.height - 1}
        />
      </div>

      <div>
        <label>Public</label>
        <input
          type="checkbox"
          name="public"
          checked={gameConfig.public}
          onChange={handleCheckBoxChange}
        />
      </div>

      <button onClick={postCreateGame}>Create Game</button>

      <p>{gameConfig.public ? "true" : "false"}</p>
    </>
  );
}
export default CreateGame;
