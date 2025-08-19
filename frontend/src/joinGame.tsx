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

      <div className="flex flex-col ">
        <h1 className="font-bold text-8xl text-gray-50 mb-5">Join Game</h1>
        <button
          className="btn w-2xl text-3xl"
          onClick={() => { setPage("public") }}>Game List</button>
        <button
          className="btn w-2xl text-3xl"
          onClick={() => { setPage("private") }}>Join with Code</button>

        <hr className="border-none my-3" />
        <button
          className="btn w-2xl text-3xl"
          onClick={() => { setPage("multiplayer") }}>Back</button>

      </div>
    </>
  );
}
export default JoinGame;
