// import { useApp } from "./appContext/useApp.tsx"
import { Game } from "./game/game.ts"
import { BoardRenderer2D } from "./game/render2d.ts"

import { useEffect, useRef, useState } from "react"

function Test() {

  const canvasRef = useRef<HTMLCanvasElement | null>(null)
  const frameRef = useRef<number | null>(null)

  const gameRef = useRef<Game | null>(null)
  const rendererRef = useRef<BoardRenderer2D | null>(null)

  useEffect(() => {
    const canvas = canvasRef.current
    if (!canvas)
      return

    const ctx = canvas.getContext("2d")
    if (!ctx)
      return

    const money = { x: 300, y: 300 }
    const time = { x: 10000, y: 10000 }
    const game = new Game("123irngrsa98fradakob", 8, 8, 4, money, time)

    // const fen = "3cq*ck*3/8/8/8/8/8/8/3CQ*CK*3 0/0/0/0/0/0/0/0/0/0/0/0/0/0 w e2 h1 0 0 600/600"

    const renderer = new BoardRenderer2D(ctx, canvas)

    gameRef.current = game
    rendererRef.current = renderer

    let lastFrame = performance.now()

    const frame = (nowFrame: number) => {
      const dt = (nowFrame - lastFrame) / 1000
      lastFrame = nowFrame

      update(dt)
      render()

      frameRef.current = requestAnimationFrame(frame)
    }

    const update = (dt: number) => {
      rendererRef.current!.update(dt)
    }

    const render = () => {
      rendererRef.current!.draw(gameRef.current!, 0, 0)
    }

    frameRef.current = requestAnimationFrame(frame)

    return () => {
      if (frameRef.current !== null) {
        cancelAnimationFrame(frameRef.current)
      }
    }

  }, [])


  return (
    <>
      <canvas ref={canvasRef} width={1200} height={800}></canvas>
    </>
  );
}
export default Test;
