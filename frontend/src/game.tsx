import { useApp, useErrorHandler, useNotifHandler } from "./appContext/useApp.tsx"
import { useGameWebSocket } from "./websocket.tsx"
import { useEffect, useRef, useState } from "react"

import { Game } from "./game/game.ts"
import { Button, createGameButtons } from "./game/button.ts"
import { InputHandler } from "./game/inputHandler.ts"
import { BoardRenderer2D } from "./game/render2d.ts"


type Message<T = unknown> = {
  type: string;
  data: T;
}


function GamePage() {
  const { setPage, accessToken, gameID, userID } = useApp()
  const { handleError } = useErrorHandler()
  const { handleNotif } = useNotifHandler()

  if (accessToken === null) {
    handleError("Not logged in")
    setPage("login")
  }

  if (gameID === null) {
    handleError("Not in a game")
    setPage("home")
  }

  const handleMessage = (event: MessageEvent) => {
    const msg = JSON.parse(event.data)
    switch (msg.type) {
      case "join": {

        break
      }
      case "start": {
        //change game settings

        const gameID = msg.data._id
        const whiteID = msg.data.whiteID
        const blackID = msg.data.whiteID
        const width = msg.data.width
        const height = msg.data.height
        const placeLine = msg.data.placeLine
        const money = msg.data.money
        const time = msg.data.startTime

        let userSide = 0
        if (userID === whiteID) {
          userSide = 0
        } else if (userID === blackID) {
          userSide = 1
        } else (
          console.log("BAD USER NOT IN GAME")
        )

        const game = gameRef.current!

        game.updateSettings(gameID, width, height, placeLine, userSide, money, time)
        break
      }
    }

    setMessages((prev) => [...prev, msg])
  }

  const { messages, setMessages, sendMessage } = useGameWebSocket(gameID, accessToken, handleMessage)


  const canvasRef = useRef<HTMLCanvasElement | null>(null)
  const frameRef = useRef<number | null>(null)

  const gameRef = useRef<Game | null>(null)
  const rendererRef = useRef<BoardRenderer2D | null>(null)

  const inputRef = useRef<InputHandler | null>(null)
  const buttonsRef = useRef<Button[]>([])

  useEffect(() => {
    const canvas = canvasRef.current
    if (!canvas)
      return

    const ctx = canvas.getContext("2d")
    if (!ctx)
      return

    if (gameID === null) {
      return
    }

    const money = [300, 300]
    const time = [10000, 10000]
    const game = new Game(gameID, 8, 8, 4, 0, money, time)

    // const fen = "cp*cn*cb*cr*cq*ck*sp*sl*/sn*sg*sc*sb*sr*sk*np*nl*/nn*ng*nb*nr*kc*kk*2/8/8/2KK*KC*NR*NB*NG*NN*/NL*NP*SK*SR*SB*SC*SG*SN*/SL*SP*CK*CQ*CR*CB*CN*CP* 0/0/0/0/0/0/0/0/0/0/0/0/0/0 w e2 h1 0 0 600/600"
    // game.updateGame(fen)

    const renderer = new BoardRenderer2D(ctx, canvas, game)

    const buttons = createGameButtons(canvas,
      renderer.UIRatio,
      game,
      handleNotif,
      renderer.switchShopScreen,
      game.clearBoardPlace,
      game.readyUp,
      game.unreadyUp,
    )
    buttonsRef.current = buttons

    const input = new InputHandler(canvas)
    inputRef.current = input

    gameRef.current = game
    rendererRef.current = renderer


    let lastFrame = performance.now()

    const frame = (nowFrame: number) => {
      const dt = (nowFrame - lastFrame) / 1000
      lastFrame = nowFrame

      update()
      render()

      frameRef.current = requestAnimationFrame(frame)
    }

    const update = () => {
      rendererRef.current!.update(gameRef.current!, inputRef.current!)



      for (const button of buttons) {
        if (!button.visible) {
          continue
        }
        button.update(inputRef.current!)
      }

      inputRef.current!.update()
    }

    const render = () => {
      rendererRef.current!.draw(gameRef.current!, 0, buttonsRef.current!, inputRef.current!)
    }

    frameRef.current = requestAnimationFrame(frame)
  }, [])


  return (
    <>
      <h1>Game</h1>
      <hr />

      <canvas ref={canvasRef} width={1000} height={1000}></canvas>
    </>
  );
}

export default GamePage;
