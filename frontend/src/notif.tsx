function Notif({ notif }: { notif: string }) {


  return (
    <>
      <div className={notif.length === 0 ? "notifOn" : "notifOff"}>
        <h2>{notif}</h2>
      </div>
    </>
  );
}

export default Notif;


