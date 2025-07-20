import { useApp, useErrorHandler } from "./appContext/useApp.tsx"

function Multiplayer() {
  const { setPage, accessToken } = useApp()
  const { handleError } = useErrorHandler()

  if (accessToken === null) {
    handleError("Not logged in")
    setPage("login")
  }

  return (
    <>
      <h1>Multiplayer</h1>
      <button onClick={() => { setPage("home") }}>Back</button>
      <hr />


      <button onClick={() => { setPage("createGame") }}>Create Game</button>
      <button onClick={() => { setPage("joinGame") }}>Join Game</button>
    </>
  );
}
export default Multiplayer;
