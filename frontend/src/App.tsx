import Login from "./login.jsx";
import Error from "./error.jsx";
import Notif from "./notif.jsx"
import Register from "./register.jsx";
import Home from "./home.jsx";
import "./App.css"


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
          className={page === "login" ? "on navButton" : "navButton"}
          onClick={() => setPage("login")}>
          Login</button>
        <button
          className={page === "register" ? "on navButton" : "navButton"}
          onClick={() => setPage("register")}>
          Register</button>

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
    default:
      return <Home />
  }

}




export default App
