import { useApp, useErrorHandler } from "./appContext/useApp.tsx"

function CreateGame() {
  const { setPage, accessToken } = useApp()
  const { handleError } = useErrorHandler()

  if (accessToken === null) {
    handleError("Not logged in")
    setPage("login")
  }

  return (
    <>
      <h1>Create Game</h1>
      <button onClick={() => { setPage("multiplayer") }}>Back</button>
      <hr />

    </>
  );
}
export default CreateGame;
