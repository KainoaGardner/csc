import { useApp } from "./appContext/useApp.tsx"

function Home() {
  const { setPage } = useApp()

  return (
    <>
      <h1>Home</h1>
      <hr />

      <button onClick={() => { setPage("multiplayer") }}>Multiplayer</button>
      <button onClick={() => { setPage("campaign") }}>Campaign</button>
      <button onClick={() => { setPage("settings") }}>Settings</button>
    </>
  );
}
export default Home;
