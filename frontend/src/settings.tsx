import { useApp } from "./appContext/useApp.tsx"

function Settings() {
  const { setPage } = useApp()

  return (
    <>
      <div className="flex flex-col ">
        <h1 className="font-bold text-8xl text-gray-50 mb-5">Settings</h1>

        <hr className="border-none my-3" />
        <button
          className="btn w-2xl text-3xl"
          onClick={() => { setPage("home") }}>Back</button>
      </div>
    </>
  );
}
export default Settings;
