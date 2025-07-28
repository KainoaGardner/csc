import { useApp } from "./appContext/useApp.tsx"
import { Game } from "./game/game.ts"
import { type Vec2 } from "./game/util.ts"

function Test() {
  const money = { x: 300, y: 300 }
  const time = { x: 10000, y: 10000 }
  const board = new Game(8, 8, 4, money, time)
  const fen = "3cq*ck*3/8/8/8/8/8/8/3CQ*CK*3 0/0/0/0/0/0/0/0/0/0/0/0/0/0 w - - 0 0 600/600"

  return (
    <>
      <h1>Test</h1>
      <button onClick={() => { board.updateGame(fen) }}>fen</button>
      <hr />
    </>
  );
}
export default Test;
