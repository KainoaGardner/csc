import { useApp } from "./appContext/useApp.tsx"

function Settings() {
  const { setPage } = useApp()

  return (
    <>
      <h1>Settings</h1>
      <button onClick={() => { setPage("home") }}>Back</button>
      <hr />

    </>
  );
}
export default Settings;
