function Error({ error }: { error: string }) {


  return (
    <>
      <div className={error.length === 0 ? "errorOn" : "errorOff"}>
        <h2>{error}</h2>
      </div>
    </>
  );
}

export default Error;


