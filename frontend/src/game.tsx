import { useApp, useErrorHandler } from "./appContext/useApp.tsx"


function Game() {
  const { setPage, accessToken, gameID } = useApp()
  const { handleError } = useErrorHandler()



  if (accessToken === null) {
    handleError("Not logged in")
    setPage("login")
  }

  if (gameID === null) {
    handleError("Not in a game")
    setPage("home")
  }

  return (
    <>
      <h1>Game</h1>
      <hr />

      <p>code: {gameID}</p>
      <canvas></canvas>

    </>
  );
}
export default Game;
