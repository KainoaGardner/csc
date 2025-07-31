import { Game } from "./game.ts"
import { InputHandler } from "./inputHandler.ts"

type ButtonAction = (...args: unknown[]) => void

interface ButtonConfig {
  x: number,
  y: number,
  width: number,
  height: number,
  strokeSize: number,
  fontSize: number,
  subFontSize: number,
  bgColor: string,
  bgColorHover: string,
  color: string,
  colorHover: string,
  subFontColor: string,
  subFontColorHover: string,
  strokeColor: string,
  strokeColorHover: string,
  text: string,
  subtext: string,
  screen: string,
  onClick: ButtonAction
}

export class Button {
  x: number
  y: number
  width: number
  height: number
  strokeSize: number
  fontSize: number
  subFontSize: number

  bgColor: string
  bgColorHover: string
  color: string
  colorHover: string
  subFontColor: string
  subFontColorHover: string
  strokeColor: string
  strokeColorHover: string

  text: string
  subtext: string
  screen: string
  onClick: ButtonAction

  hovering: boolean = false
  visible: boolean = false

  constructor(buttonConfig: ButtonConfig) {
    this.x = buttonConfig.x
    this.y = buttonConfig.y
    this.width = buttonConfig.width
    this.height = buttonConfig.height
    this.strokeSize = buttonConfig.strokeSize
    this.fontSize = buttonConfig.fontSize
    this.subFontSize = buttonConfig.subFontSize
    this.bgColor = buttonConfig.bgColor
    this.bgColorHover = buttonConfig.bgColorHover
    this.color = buttonConfig.color
    this.colorHover = buttonConfig.colorHover
    this.subFontColor = buttonConfig.subFontColor
    this.subFontColorHover = buttonConfig.subFontColorHover
    this.strokeColor = buttonConfig.strokeColor
    this.strokeColorHover = buttonConfig.strokeColorHover
    this.text = buttonConfig.text
    this.subtext = buttonConfig.subtext
    this.screen = buttonConfig.screen
    this.onClick = buttonConfig.onClick
  }

  draw(ctx: CanvasRenderingContext2D) {
    let increaseSize: number
    let bgColor: string
    let color: string
    let subColor: string
    let strokeColor: string
    if (this.hovering) {
      increaseSize = 1.0
      bgColor = this.bgColorHover
      color = this.colorHover
      subColor = this.subFontColorHover
      strokeColor = this.strokeColorHover
    } else {
      increaseSize = 1.0
      bgColor = this.bgColor
      color = this.color
      subColor = this.subFontColor
      strokeColor = this.strokeColor
    }

    ctx.textAlign = "center"
    ctx.textBaseline = "middle";


    //bg
    ctx.fillStyle = bgColor
    ctx.fillRect(this.x - this.width * (increaseSize - 1.0) / 2, this.y - this.height * (increaseSize - 1.0) / 2, this.width * increaseSize, this.height * increaseSize)

    //color
    ctx.fillStyle = color
    ctx.font = `${this.fontSize}px Arial`
    ctx.fillText(this.text, this.x + this.width / 2, this.y + this.height / 2)

    //subColor
    ctx.fillStyle = subColor
    ctx.font = `${this.subFontSize}px Arial`
    ctx.fillText(this.subtext, this.x + this.width / 2, this.y + this.height / 2 + this.height / 4)

    //stroke
    ctx.strokeStyle = strokeColor
    ctx.lineWidth = this.strokeSize
    ctx.strokeRect(this.x - this.width * (increaseSize - 1.0) / 2, this.y - this.height * (increaseSize - 1.0) / 2, this.width * increaseSize, this.height * increaseSize)
  }

  update(input: InputHandler) {
    this.checkHoveringButton(input)
    this.clickButton(input)
  }

  checkHoveringButton(input: InputHandler): boolean {
    let increaseSize = 1.1
    if (!this.hovering)
      increaseSize = 1.0

    const result = (this.x - this.width * (increaseSize - 1.0) / 2 <= input.mouse.x
      && input.mouse.x <= this.x + this.width + this.width * (increaseSize - 1.0) / 2
      && this.y - this.height * (increaseSize - 1.0) / 2 <= input.mouse.y
      && input.mouse.y <= this.y + this.height + this.height * (increaseSize - 1.0) / 2)

    this.hovering = result
    return result
  }

  clickButton(input: InputHandler) {
    if (input.mouse.justPressed && this.checkHoveringButton(input)) {
      this.onClick()
    }
  }
}


export function createGameButtons(canvas: HTMLCanvasElement, UIRatio: number, game: Game, handleNotif: (err: string) => void, switchShopScreen: () => void): Button[] {
  const result: Button[] = []

  const width = 600 * UIRatio
  const height = 100 * UIRatio

  const copyID = (gameID: string) => {
    navigator.clipboard.writeText(gameID);
    handleNotif("Text Copied")
  }

  const idButtonConfig = {
    x: canvas.width / 2 - width / 2,
    y: canvas.height / 2 - height / 2,
    width: width,
    height: height,
    strokeSize: 5 * UIRatio,
    fontSize: 40 * UIRatio,
    subFontSize: 15 * UIRatio,
    bgColor: "#ecf0f1",
    bgColorHover: "#7f8c8d",
    color: "#000",
    colorHover: "#000",
    subFontColor: "#000",
    subFontColorHover: "#000",
    strokeColor: "#AAA",
    strokeColorHover: "#BBB",
    text: game.id,
    subtext: "Click to copy ID",
    screen: "connect",
    onClick: () => copyID(game.id)
  }

  const idButton = new Button(idButtonConfig)
  result.push(idButton)

  const unitButtonConfig = {
    x: 0,
    y: 900 * UIRatio,
    width: 100 * UIRatio,
    height: 100 * UIRatio,
    strokeSize: 5 * UIRatio,
    fontSize: 25 * UIRatio,
    subFontSize: 13 * UIRatio,
    bgColor: "#ecf0f1",
    bgColorHover: "#7f8c8d",
    color: "#000",
    colorHover: "#000",
    subFontColor: "#000",
    subFontColorHover: "#000",
    strokeColor: "#AAA",
    strokeColorHover: "#BBB",
    text: "Shop",
    subtext: "Click to Switch",
    screen: "place",
    onClick: () => switchShopScreen()
  }


  if (game.userSide === 1) {
    unitButtonConfig.y = 0
  }
  const unitSwitchButton = new Button(unitButtonConfig)
  result.push(unitSwitchButton)

  const readyUp = () => {
    console.log("Ready UP")
  }

  const readyButtonConfig = {
    x: 200 * UIRatio,
    y: 250 * UIRatio,
    width: 600 * UIRatio,
    height: 100 * UIRatio,
    strokeSize: 5 * UIRatio,
    fontSize: 50 * UIRatio,
    subFontSize: 0,
    bgColor: "#ecf0f1",
    bgColorHover: "#7f8c8d",
    color: "#000",
    colorHover: "#000",
    subFontColor: "#000",
    subFontColorHover: "#000",
    strokeColor: "#AAA",
    strokeColorHover: "#BBB",
    text: "Ready",
    subtext: "",
    screen: "place",
    onClick: () => readyUp()
  }

  if (game.userSide === 1) {
    readyButtonConfig.y = 650 * UIRatio
  }

  const readyButton = new Button(readyButtonConfig)
  result.push(readyButton)

  readyButtonConfig.text = "Unready"
  readyButtonConfig.screen = "placeReady"
  const unreadyButton = new Button(readyButtonConfig)
  result.push(unreadyButton)

  return result
}
