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
      <h1>Campaign</h1>
      <button onClick={() => { setPage("home") }}>Back</button>
      <hr />

    </>
  );
}
export default Campaign;
