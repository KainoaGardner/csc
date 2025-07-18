import Login from "./login.jsx";
import Error from "./error.jsx";
import Register from "./register.jsx";
import Home from "./home.jsx";
import { useState } from "react"
import "./App.css"

type Page = "home" | "login" | "register"

function App() {
  const [page, setPage] = useState<Page>("home")
  const [error, setError] = useState<string>("")

  const handleError = (err: string) => {
    setError(err);
    setTimeout(() => {
      setError("");
    }, 5000);
  };



  return (
    <>
      <Error error={error} />

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

      <Tab page={page} handleError={handleError} />

    </>
  )
}

function Tab({ page, handleError }: { page: Page, handleError: (err: string) => void }) {
  switch (page) {
    case "login":
      return <Login />
    case "register":
      return <Register handleError={handleError} />
    default:
      return <Home />
  }

}




export default App
