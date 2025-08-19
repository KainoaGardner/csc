function Notif({ notif }: { notif: string }) {


  return (
    <>
      <div
        className={notif.length !== 0 ? "leftAnimation absolute right-5 p-5 bg-neutral-700 border-6 border-neutral-400 animate-bounce" : "text-0 absolute"}>
        <h2 className="text-4xl text-green-400 font-bold"
        >{notif}</h2>
      </div>
    </>
  );
}

export default Notif;


