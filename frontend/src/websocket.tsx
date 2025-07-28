import API_URL from "./env.tsx"
import { useEffect, useRef, useState } from "react"
import { useApp } from "./appContext/useApp.tsx"

type Message<T = unknown> = {
  type: string;
  data: T;
}

const joinRequest: Message<null> = {
  type: "join",
  data: null,
}

export function useGameWebSocket(gameID: string | null, accessToken: string | null, onMessage?: (msg: Message) => void) {
  // const { handleError } = useErrorHandler()
  // const { handleNotif } = useNotifHandler()


  // const { setPage } = useApp()


  const [messages, setMessages] = useState<Message[]>([])
  const ws = useRef<WebSocket | null>(null)

  useEffect(() => {
    if (gameID === null || accessToken === null) {
      return
    }

    const socket = new WebSocket(API_URL + "ws/" + gameID + "/" + accessToken)
    ws.current = socket

    socket.onopen = () => {
      console.log("connected to websocket")
      sendMessage(joinRequest)
    }

    socket.onmessage = (event) => {
      const data = JSON.parse(event.data)
      if (onMessage) onMessage(data)
    }

    socket.onclose = (event) => {
      console.log("Websocket closed", event.code, event.reason)
      // setPage("home")
    }

    socket.onerror = (error) => {
      console.error("Websocket error: ", error)
    }

    return () => {
      socket.close()
    }

  }, [gameID, accessToken])


  const sendMessage = (msg: Message) => {
    if (ws.current?.readyState === WebSocket.OPEN) {
      ws.current.send(JSON.stringify(msg))
    }
  }

  return { messages, setMessages, sendMessage }
}
