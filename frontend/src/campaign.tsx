import { useApp, useErrorHandler } from "./appContext/useApp.tsx"

function Campaign() {
  const { setPage, accessToken } = useApp()
  const { handleError } = useErrorHandler()

  if (accessToken === null) {
    handleError("Not logged in")
    setPage("login")
  }

  return (
    <>
      <div className="flex flex-col ">
        <h1 className="font-bold text-8xl text-gray-50 mb-5">Campaign</h1>
        <h1 className="font-bold text-6xl text-red-500 mb-5">Work in Progress</h1>

        <hr className="border-none my-3" />
        <button
          className="btn w-2xl text-3xl"
          onClick={() => { setPage("home") }}>Back</button>

      </div>
    </>
  );
}
export default Campaign;
