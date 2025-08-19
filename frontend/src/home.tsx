import { useApp } from "./appContext/useApp.tsx"

function Home() {
  const { setPage } = useApp()

  return (
    <div className="flex flex-col ">
      <h1 className="font-bold text-8xl text-gray-50">CSC</h1>
      <h2 className="font-bold text-4xl text-neutral-400 ml-1 mb-4">Chess Shogi Checkers</h2>
      <img className="my-5" src="" width="500" height="500"></img>

      <div className="flex flex-col">
        <button className="btn w-2xl text-3xl" onClick={() => { setPage("multiplayer") }}>Multiplayer</button>
        <button className="btn w-2xl text-3xl" onClick={() => { setPage("campaign") }}>Campaign</button>
        <button className="btn w-2xl text-3xl" onClick={() => { setPage("settings") }}>Settings</button>
      </div>

    </div>
  );
}
export default Home;
