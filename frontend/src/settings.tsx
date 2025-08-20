import { useApp } from "./appContext/useApp.tsx"

import { updateVolume } from "./sounds.ts"

function Settings() {
  const { setPage, volume, setVolume } = useApp()

  const handleVolumeChange = (vol: number) => {
    setVolume(vol)
    updateVolume(vol)
  }

  return (
    <>
      <div className="flex flex-col ">
        <h1 className="font-bold text-8xl text-gray-50 mb-5">Settings</h1>

        <h2 className="font-bold text-3xl text-gray-50">Volume</h2>
        <h2 className="font-bold text-3xl text-neutral-400 mb-4">{volume}</h2>
        <input
          className="w-2xl px-4"
          type="range"
          min="0"
          max="100"
          value={volume}
          onChange={(e) => handleVolumeChange(parseInt(e.target.value))}
        />


        <hr className="border-none my-3" />
        <button
          className="btn w-2xl text-3xl"
          onClick={() => { setPage("home") }}>Back</button>
      </div>
    </>
  );
}
export default Settings;
