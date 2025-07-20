import { useApp, useErrorHandler } from "./appContext/useApp.tsx"

function JoinGame() {
  const { setPage, accessToken } = useApp()
  const { handleError } = useErrorHandler()

  if (accessToken === null) {
    handleError("Not logged in")
    setPage("login")
  }

  return (
    <>
      <h1>Join Game</h1>
      <button onClick={() => { setPage("multiplayer") }}>Back</button>
      <hr />

      <button onClick={() => { setPage("public") }}>Game List</button>
      <button onClick={() => { setPage("private") }}>Join with Code</button>

    </>
  );
}
export default JoinGame;
