import { useApp, useErrorHandler } from "./appContext/useApp.tsx"
import { useGameWebSocket } from "./websocket.tsx"
import { useEffect, useRef, useState } from "react"

type Message<T = unknown> = {
  type: string;
  data: T;
}


function Game() {
  const { setPage, accessToken, gameID } = useApp()
  const { handleError } = useErrorHandler()

  if (accessToken === null) {
    handleError("Not logged in")
    setPage("login")
  }

  if (gameID === null) {
    handleError("Not in a game")
    setPage("home")
  }

  const handleMessage = (msg: Message) => {
    setMessages((prev) => [...prev, msg])
  }

  const { messages, setMessages, sendMessage } = useGameWebSocket(gameID, accessToken, handleMessage)

  const canvasRef = useRef<HTMLCanvasElement | null>(null)
  useEffect(() => {
    const canvas = canvasRef.current
    if (!canvas)
      return

    const ctx = canvas.getContext("2d")
  })

  return (
    <>
      <h1>Game</h1>
      <hr />

      <p>code: {gameID}</p>
      {messages.map((msg, index) => (
        <div key={index}>
          <h3>{msg.type}</h3>
          {/* <p>{msg.data}</p> */}
        </div>
      ))}

      <canvas ref={canvasRef} width={512} height={512}></canvas>

    </>
  );
}
export default Game;
