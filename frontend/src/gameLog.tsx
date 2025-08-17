import API_URL from "./env.tsx"

import { useApp, useErrorHandler, useNotifHandler } from "./appContext/useApp.tsx"
import { useState, useEffect, useRef } from "react"

import { Game } from "./game/game.ts"
import { GameLog } from "./game/gameLog.ts"
import { InputHandler } from "./game/inputHandler.ts"
import { BoardRenderer2D } from "./game/render2d.ts"

import "./css/gameLog.css"

type GameLogData = {
  id: string
  date: Date
  moveCount: number
  moves: string[]
  boardStates: string[]
  boardHeight: number,
  boardWidth: number,
  boardPlaceLine: number,

  winner: number
  reason: string
}

const emptyGameLogData = {
  id: "",
  date: new Date,
  moveCount: 0,
  moves: [],
  boardStates: [],
  boardHeight: 0,
  boardWidth: 0,
  boardPlaceLine: 0,

  winner: 0,
  reason: "",
}

function GameLogPage() {
  const { setPage, accessToken, gameLogID } = useApp()
  const { handleError } = useErrorHandler()
  const { handleNotif } = useNotifHandler()

  const [gameLogData, setGameLogData] = useState<GameLogData>(emptyGameLogData)
  const [viewableMoves, setViewableMoves] = useState<string[]>([])
  const [moveIndex, setMoveIndex] = useState<number>(0)

  const sendMessage = () => {
  }

  if (accessToken === null) {
    handleError("Not logged in")
    setPage("login")
  }

  if (gameLogID === null) {
    handleError("Not in a game")
    setPage("userStats")
  }


  const getGameLog = async () => {
    try {
      const response = await fetch(API_URL + "log/" + gameLogID, {
        method: "Get",
        headers: {
          "Content-Type": "application/json; charset=utf-8",
        },
      });

      const data = await response.json();
      if (response.ok) {
        const updatedGameLog = {
          id: data.data._id,
          date: data.data.date,
          moveCount: data.data.moveCount,
          moves: data.data.moves,
          boardStates: data.data.boardStates,
          boardHeight: data.data.boardHeight,
          boardWidth: data.data.boardWidth,
          boardPlaceLine: data.data.boardPlaceLine,
          winner: data.data.winner,
          reason: data.data.reason,
        }
        setGameLogData(updatedGameLog)
        handleMoveListUpdate(updatedGameLog.moves, 0)
        init(updatedGameLog)
      } else {
        handleError(data.error);
      }
    } catch (error) {
      console.log(error);
    }
  };


  const handleMoveListUpdate = (moves: string[], moveIndex: number) => {
    const result = []
    for (let i = -1; i < 10; i++) {
      if (moveIndex + i < 0) {
        continue
      }
      if (moveIndex + i >= moves.length) {
        break
      }

      const move = moves[moveIndex + i]
      result.push(move)
    }

    setViewableMoves(result)
  }

  const handleNextMove = (gameLog: GameLog | null, game: Game | null, renderer: BoardRenderer2D | null) => {
    if (gameLog === null || game === null || renderer === null) return

    gameLog.nextMove(game, renderer)
    setMoveIndex(gameLog.moveIndex)
    handleMoveListUpdate(gameLog.moves, gameLog.moveIndex)
  }

  const handlePrevMove = (gameLog: GameLog | null, game: Game | null, renderer: BoardRenderer2D | null) => {
    if (gameLog === null || game === null || renderer === null) return

    gameLog.prevMove(game, renderer)
    setMoveIndex(gameLog.moveIndex)
    handleMoveListUpdate(gameLog.moves, gameLog.moveIndex)
  }

  const canvasRef = useRef<HTMLCanvasElement | null>(null)
  const frameRef = useRef<number | null>(null)
  const inputRef = useRef<InputHandler | null>(null)

  const gameRef = useRef<Game | null>(null)
  const gameLogRef = useRef<GameLog | null>(null)
  const rendererRef = useRef<BoardRenderer2D | null>(null)

  const init = (log: GameLogData) => {
    const canvas = canvasRef.current
    if (!canvas)
      return

    const ctx = canvas.getContext("2d")
    if (!ctx)
      return

    if (gameLogID === null) {
      return
    }

    const input = new InputHandler(canvas)
    inputRef.current = input

    const gameLog = new GameLog(
      log.id,
      log.date,
      log.moveCount,
      log.moves,
      log.boardStates,
      log.winner,
      log.reason,
    )

    gameLogRef.current = gameLog

    const money = [300, 300]
    const time = [10000, 10000]
    const game = new Game(
      gameLogID,
      log.boardWidth,
      log.boardHeight,
      log.boardPlaceLine,
      0,
      money,
      time
    )
    game.state = 4
    game.updateOver(log.winner, log.reason, 4)
    if (log.boardStates.length !== 0) {
      game.updateGame(log.boardStates[0])
    }

    const renderer = new BoardRenderer2D(ctx, canvas, game, handleNotif, sendMessage)

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
      rendererRef.current!.drawGameLog(gameRef.current!, gameLogRef.current!, 0, inputRef.current!)
    }

    frameRef.current = requestAnimationFrame(frame)

    return () => {
      cancelAnimationFrame(frameRef.current!)
      input.cleanup()
    }
  }

  useEffect(() => {
    getGameLog()
  }, [])

  return (
    <>
      <button onClick={() => { setPage("userStats") }}>Back</button>
      <canvas ref={canvasRef} width={1000} height={1000}></canvas>

      <div>
        <div>
          <button onClick={() => handlePrevMove(gameLogRef.current, gameRef.current, rendererRef.current)}>Prev</button>
          <button onClick={() => handleNextMove(gameLogRef.current, gameRef.current, rendererRef.current)}>Next</button>
        </div>
      </div>

      <ul>
        {viewableMoves.map((move, index) => (
          <li key={index} className={checkCurrentMove(moveIndex, index) ? "moveSelected" : ""}>{getIndex(moveIndex, index)}: {move}</li>
        ))}
      </ul>

    </>
  );
}

function getIndex(moveIndex: number, index: number): number {
  if (moveIndex === 0) {
    return index
  }

  return moveIndex + index - 1
}

function checkCurrentMove(moveIndex: number, index: number): boolean {
  if (moveIndex === 0) {
    return index === 0
  }

  return index === 1
}

export default GameLogPage;
