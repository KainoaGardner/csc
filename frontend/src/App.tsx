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


import { useApp, useLogoutHandler } from "./appContext/useApp.tsx"
import type { Page } from "./appContext/appContext.tsx"

function App() {
  const { accessToken, page, setPage, error, notif } = useApp();
  const { handleLogout } = useLogoutHandler()

  return (
    <>
      <div className="flex flex-col h-full">
        <Error error={error} />
        <Notif notif={notif} />

        <div className="nav flex  flex-1 mx-1 mb-8">
          <button
            className={page === "home" ? "on navButton text-2xl" : "navButton text-2xl"}
            onClick={() => setPage("home")}>
            Home</button>
          <button
            className={page === "userStats" ? "on navButton text-2xl" : "navButton text-2xl"}
            onClick={() => setPage("userStats")}>
            User</button>
          <LogInOutButton page={page} setPage={setPage} accessToken={accessToken} handleLogout={handleLogout} />
          <button
            className={page === "register" ? "on navButton text-2xl" : "navButton text-2xl"}
            onClick={() => setPage("register")}>
            Register</button>
          <button
            className={page === "test" ? "on navButton text-2xl" : "navButton text-2xl"}
            onClick={() => setPage("test")}>
            Test</button>
        </div>

        <div className="flex-8">
          <Tab page={page} />
        </div>
      </div>
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

function LogInOutButton(
  {
    page,
    setPage,
    accessToken,
    handleLogout
  }: {
    page: string,
    setPage: (page: Page) => void,
    accessToken: string | null,
    handleLogout: () => void
  }) {
  if (accessToken === null) {
    return <button
      className={page === "login" ? "on navButton text-2xl" : "navButton text-2xl"}
      onClick={() => setPage("login")}>
      Login</button>
  } else {
    return <button
      className={page === "logout" ? "on navButton text-2xl" : "navButton text-2xl"}
      onClick={handleLogout}>
      Logout</button>
  }
}


export default App
