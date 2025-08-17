import GamePage from "./game.jsx";
import Test from "./test.jsx";

import Login from "./login.jsx";
import Multiplayer from "./multiplayer.jsx";
import CreateGame from "./createGame.jsx";
import JoinGame from "./joinGame.jsx";

import PrivateJoin from "./private.jsx";
import PublicJoin from "./public.jsx";

import Campign from "./campaign.jsx";

import GameLog from "./gameLog.jsx";

import Settings from "./settings.jsx";
import User from "./user.js";
import Error from "./error.jsx";
import Notif from "./notif.jsx"
import Register from "./register.jsx";
import Home from "./home.jsx";

import { useApp } from "./appContext/useApp.tsx"
import type { Page } from "./appContext/appContext.tsx"

function App() {
  const { page, setPage, error, notif } = useApp();

  return (
    <>
      <Error error={error} />
      <Notif notif={notif} />

      <div className="nav">
        <button
          className={page === "home" ? "on navButton" : "navButton"}
          onClick={() => setPage("home")}>
          Home</button>
        <button
          className={page === "userStats" ? "on navButton" : "navButton"}
          onClick={() => setPage("userStats")}>
          User Stats</button>
        <button
          className={page === "login" ? "on navButton" : "navButton"}
          onClick={() => setPage("login")}>
          Login</button>
        <button
          className={page === "register" ? "on navButton" : "navButton"}
          onClick={() => setPage("register")}>
          Register</button>
        <button
          className={page === "test" ? "on navButton" : "navButton"}
          onClick={() => setPage("test")}>
          Test</button>


      </div>
      <Tab page={page} />
    </>
  )
}

function Tab({ page }: { page: Page }) {
  switch (page) {
    case "login":
      return <Login />
    case "register":
      return <Register />
    case "userStats":
      return <User />
    case "multiplayer":
      return <Multiplayer />
    case "createGame":
      return <CreateGame />
    case "joinGame":
      return <JoinGame />
    case "private":
      return <PrivateJoin />
    case "public":
      return <PublicJoin />
    case "campaign":
      return <Campign />
    case "settings":
      return <Settings />
    case "game":
      return <GamePage />
    case "test":
      return <Test />
    case "gameLog":
      return <GameLog />
    default:
      return <Home />
  }
}


export default App
