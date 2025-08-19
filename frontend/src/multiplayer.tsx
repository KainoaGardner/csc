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
      <div className="flex flex-col ">
        <h1 className="font-bold text-8xl text-gray-50 mb-5">Multiplayer</h1>

        <button
          className="btn w-2xl text-3xl"
          onClick={() => { setPage("createGame") }}>Create Game</button>
        <button
          className="btn w-2xl text-3xl"
          onClick={() => { setPage("joinGame") }}>Join Game</button>

        <hr className="border-none my-3" />
        <button
          className="btn w-2xl text-3xl"
          onClick={() => { setPage("home") }}>Back</button>


      </div>
    </>
  );
}
export default Multiplayer;
