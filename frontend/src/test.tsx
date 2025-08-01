// import { useApp } from "./appContext/useApp.tsx"

import { Game } from "./game/game.ts"
import { Button, createGameButtons } from "./game/button.ts"
import { InputHandler } from "./game/inputHandler.ts"
import { BoardRenderer2D } from "./game/render2d.ts"

import { useNotifHandler } from "./appContext/useApp.tsx"
import { useEffect, useRef } from "react"

function Test() {
  const { handleNotif } = useNotifHandler()
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

    const money = [300, 300]
    const time = [10000, 10000]
    const game = new Game("123irngrsa98fradakob", 8, 8, 4, 0, money, time)
    game.state = 1

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
      inputRef.current!.update()
      for (const button of buttons) {
        if (!button.visible) {
          continue
        }
        button.update(inputRef.current!)
      }
    }

    const render = () => {
      rendererRef.current!.draw(gameRef.current!, 0, buttonsRef.current!, inputRef.current!)
    }

    frameRef.current = requestAnimationFrame(frame)

    return () => {
      if (frameRef.current !== null) {
        cancelAnimationFrame(frameRef.current)
      }
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
