import { useApp, useErrorHandler, useNotifHandler } from "./appContext/useApp.tsx"
import { useGameWebSocket } from "./websocket.tsx"
import { useEffect, useRef } from "react"

import { Game } from "./game/game.ts"
import { InputHandler } from "./game/inputHandler.ts"
import { BoardRenderer2D } from "./game/render2d.ts"

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

  let lastFen: string | null = null
  const handleMessage = (event: MessageEvent) => {
    const msg = JSON.parse(event.data)
    switch (msg.type) {
      case "error": {
        const game = gameRef.current!
        if (lastFen !== null) {
          game.updateGame(lastFen)
        }
        break
      }
      case "start": {
        //change game settings

        const gameID = msg.data._id
        const whiteID = msg.data.whiteID
        const blackID = msg.data.blackID
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
        game.state = 1

        const renderer = rendererRef.current!
        renderer.updateButtons(canvasRef.current!, game, handleNotif, sendMessage)
        renderer.updateButtonScreen(game)

        break
      }
      case "ready": {
        const game = gameRef.current!
        const ready = msg.data.ready
        const state = msg.data.state
        const fen = msg.data.fen
        game.updateReady(ready, state)
        game.updateGame(fen)
        lastFen = fen

        const renderer = rendererRef.current!
        renderer.updateButtonScreen(game)
        break
      }
      case "move": {
        const game = gameRef.current!

        game.state = 2
        const fen = msg.data.fen
        game.updateGame(fen)
        lastFen = fen

        break
      }
      case "draw": {
        const game = gameRef.current!
        const draw = msg.data.draw

        game.updateDraw(draw)
        break
      }
      case "over": {
        const game = gameRef.current!

        const fen = msg.data.fen
        const state = msg.data.state
        const reason = msg.data.reason
        const winner = msg.data.winner

        game.updateGame(fen)
        game.updateOver(winner, reason, state)
        lastFen = fen
        break
      }
    }
  }

  const { sendMessage } = useGameWebSocket(gameID, accessToken, handleMessage)

  const canvasRef = useRef<HTMLCanvasElement | null>(null)
  const frameRef = useRef<number | null>(null)

  const gameRef = useRef<Game | null>(null)
  const rendererRef = useRef<BoardRenderer2D | null>(null)

  const inputRef = useRef<InputHandler | null>(null)

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

    const renderer = new BoardRenderer2D(ctx, canvas, game, handleNotif, sendMessage)

    const input = new InputHandler(canvas)
    inputRef.current = input

    gameRef.current = game
    rendererRef.current = renderer

    const frame = () => {
      update()
      render()

      frameRef.current = requestAnimationFrame(frame)
    }

    const update = () => {
      rendererRef.current!.update(gameRef.current!, inputRef.current!, sendMessage)
      inputRef.current!.update()
    }

    const render = () => {
      rendererRef.current!.draw(gameRef.current!, 0, inputRef.current!)
    }

    frameRef.current = requestAnimationFrame(frame)

    return () => {
      cancelAnimationFrame(frameRef.current!)
      input.cleanup()
    }
  }, [])

  return (
    <>
      <canvas ref={canvasRef} width={1000} height={1000}></canvas>
    </>
  );
}

export default GamePage;
