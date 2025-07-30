import { type Mouse } from "./util.ts"

type ButtonAction = (...args: unknown[]) => void

export class Button {
  x: number
  y: number
  width: number
  height: number
  strokeSize: number
  fontSize: number

  bgColor: string
  color: string
  strokeColor: string

  text: string
  screen: string
  onClick: ButtonAction


  hovering: boolean = false
  visible: boolean = false

  constructor(
    x: number,
    y: number,
    width: number,
    height: number,
    strokeSize: number,
    fontSize: number,
    bgColor: string,
    color: string,
    strokeColor: string,
    text: string,
    screen: string,
    onClick: ButtonAction) {
    this.x = x
    this.y = y
    this.width = width
    this.height = height
    this.strokeSize = strokeSize
    this.fontSize = fontSize
    this.bgColor = bgColor
    this.color = color
    this.strokeColor = strokeColor
    this.text = text
    this.screen = screen
    this.onClick = onClick
  }

  draw(ctx: CanvasRenderingContext2D) {
    let increaseSize = 1.1
    if (!this.hovering)
      increaseSize = 1.0

    ctx.textAlign = "center"
    ctx.textBaseline = "middle";

    ctx.fillStyle = this.bgColor
    ctx.fillRect(this.x - this.width * (increaseSize - 1.0) / 2, this.y - this.height * (increaseSize - 1.0) / 2, this.width * increaseSize, this.height * increaseSize)

    ctx.fillStyle = this.color
    ctx.font = `${this.fontSize}px Arial`
    ctx.fillText(this.text, this.x + this.width / 2, this.y + this.height / 2)

    ctx.strokeStyle = this.strokeColor
    ctx.lineWidth = this.strokeSize
    ctx.strokeRect(this.x - this.width * (increaseSize - 1.0) / 2, this.y - this.height * (increaseSize - 1.0) / 2, this.width * increaseSize, this.height * increaseSize)
  }

  checkHoveringButton(mouse: Mouse): boolean {
    let increaseSize = 1.1
    if (!this.hovering)
      increaseSize = 1.0

    const result = (this.x - this.width * (increaseSize - 1.0) / 2 <= mouse.x
      && mouse.x <= this.x + this.width + this.width * (increaseSize - 1.0) / 2
      && this.y - this.height * (increaseSize - 1.0) / 2 <= mouse.y
      && mouse.y <= this.y + this.height + this.height * (increaseSize - 1.0) / 2)

    this.hovering = result
    return result
  }

  clickButton(mouse: Mouse) {
    if (this.checkHoveringButton(mouse)) {
      this.onClick()
      mouse.pressed = false
    }
  }
}


export function createGameButtons(canvas: HTMLCanvasElement, UIRatio: number, gameID: string, handleNotif: (err: string) => void): Button[] {
  const result: Button[] = []

  const width = 600 * UIRatio
  const height = 100 * UIRatio
  const x = canvas.width / 2 - width / 2
  const y = canvas.height / 2 - height / 2
  const strokeSize = 5 * UIRatio
  const fontSize = 40 * UIRatio
  const bgColor = "#ecf0f1"
  const color = "#000"
  const strokeColor = "#AAA"
  const text = gameID
  const copyID = (gameID: string) => {
    navigator.clipboard.writeText(gameID);
    handleNotif("Text Copied")
  }

  const idButton = new Button(
    x,
    y,
    width,
    height,
    strokeSize,
    fontSize,
    bgColor,
    color,
    strokeColor,
    text,
    "connect",
    () => copyID(gameID),
  )

  result.push(idButton)

  return result
}
