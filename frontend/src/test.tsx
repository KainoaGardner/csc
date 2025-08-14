import { useApp, useNotifHandler } from "./appContext/useApp.tsx"
import { useGameWebSocket } from "./websocket.tsx"
import { useEffect, useRef } from "react"

import { Game } from "./game/game.ts"
import { InputHandler } from "./game/inputHandler.ts"
import { BoardRenderer2D } from "./game/render2d.ts"

function Test() {
  const { accessToken, userID } = useApp()
  // const { handleError } = useErrorHandler()
  const { handleNotif } = useNotifHandler()

  const gameID = "12039871209837"

  // if (accessToken === null) {
  //   handleError("Not logged in")
  //   setPage("login")
  // }

  // if (gameID === null) {
  //   handleError("Not in a game")
  //   setPage("home")
  // }

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
        renderer.updateButtonScreen(game)

        break
      }
      case "place": {
        break
      }
      case "ready": {
        const game = gameRef.current!
        const ready = msg.data.ready
        const state = msg.data.state
        game.updateReady(ready, state)

        const renderer = rendererRef.current!
        renderer.updateButtonScreen(game)
        break
      }
      case "move": {
        // const game = gameRef.current!

        console.log(msg)
        // const fen = msg.data.fen
        //
        // game.state = 2
        // game.updateGame(fen)

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
    const game = new Game(gameID, 8, 8, 4, 1, money, time)
    game.state = 2

    // const fen = "4ck*3/8/8/8/8/8/8/3CQ*CK*3 5/5/5/5/5/5/5/0/0/0/0/1/0/0 w e2 h1 0 0 600/600"
    // const fen = "3ck*cr*3/CP-4CP-1CP-/8/8/8/8/4SC*3/4CK*1SR*CR* 5/5/5/5/5/5/5/0/0/0/0/1/0/0 w e2 h1 0 0 600/600"
    const fen = "3ck*cr*3/CP-4CP-1CP-/8/8/8/4sb-3/cp-3SC*2cp-/4CK*1SR*1 5/5/5/5/5/5/5/0/0/0/0/1/0/0 b e2 h1 0 0 600/600"
    // const fen = "cp*cn*cb*cr*cq*ck*sp*sl*/sn*sg*sc*sb*sr*sk*np*nl*/nn*ng*nb*nr*kc*kk*2/8/8/2KK*KC*NR*NB*NG*NN*/NL*NP*SK*SR*SB*SC*SG*SN*/SL*SP*CK*CQ*CR*CB*CN*CP* 0/0/0/0/0/0/0/0/0/0/0/0/0/0 w e2 h1 0 0 600/600"
    game.updateGame(fen)

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
export default Test;
