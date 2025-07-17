// import Login from "./login.jsx";
import { useState } from "react"

function App() {
  const [page, setPage] = useState("home")

  return (
    <>
      <div className="nav">
        <button
          className={page === "home" ? "on navButton" : "navButton"}
          onClick={() => setPage("home")}>
          Home </button>
        <button
          className={page === "login" ? "on navButton" : "navButton"}
          onClick={() => setPage("login")}>
          Login </button>

      </div>
      {/* <Tab page={page} /> */}
    </>
  )
}

// function Tab(page) {
//
// }

export default App
