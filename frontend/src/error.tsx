function Error({ error }: { error: string }) {


  return (
    <>
      <div
        className={error.length !== 0 ? "leftAnimation absolute right-5 p-5 bg-neutral-700 border-6 border-neutral-400 animate-bounce" : "text-0 absolute"}>
        <h2 className="text-4xl text-red-400 font-bold"
        >{error}</h2>
      </div>
    </>
  );
}

export default Error;


