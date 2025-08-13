import { Game } from "./game.ts"
import { type Message } from "./util.ts"
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

  update(input: InputHandler, game: Game, updateButtonScreen: (game: Game) => void) {

    this.checkHoveringButton(input)
    this.clickButton(input, game, updateButtonScreen)
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

  clickButton(input: InputHandler, game: Game, updateButtonScreen: (game: Game) => void) {
    if (input.mouse.justPressed[0] && this.checkHoveringButton(input)) {
      this.onClick()
      input.mouse.justPressed[0] = false
      updateButtonScreen(game)
    }
  }
}

export function createGameButtons(canvas: HTMLCanvasElement,
  UIRatio: number,
  game: Game,
  handleNotif: (err: string) => void,
  sendMessage: (msg: Message<unknown>) => void,
  switchShopScreen: () => void,
  clearBoard: () => void,
  readyUp: (sendMessage: (msg: Message<unknown>) => void) => void,
  unreadyUp: (sendMessage: (msg: Message<unknown>) => void) => void,
  drawUp: (sendMessage: (msg: Message<unknown>) => void) => void,
  undrawUp: (sendMessage: (msg: Message<unknown>) => void) => void,
  pressResign: () => void,
  cancelResign: () => void,
  confirmResign: (sendMessage: (msg: Message<unknown>) => void) => void,

): Map<string, Button> {
  if (game.userSide === 0) {
    return createWhiteButtons(
      canvas,
      UIRatio,
      game,
      handleNotif,
      sendMessage,
      switchShopScreen,
      clearBoard,
      readyUp,
      unreadyUp,
      drawUp,
      undrawUp,
      pressResign,
      cancelResign,
      confirmResign,
    )
  } else {
    return createBlackButtons(
      canvas,
      UIRatio,
      game,
      handleNotif,
      sendMessage,
      switchShopScreen,
      clearBoard,
      readyUp,
      unreadyUp,
      drawUp,
      undrawUp,
      pressResign,
      cancelResign,
      confirmResign,
    )
  }
}

function createWhiteButtons(canvas: HTMLCanvasElement,
  UIRatio: number,
  game: Game,
  handleNotif: (err: string) => void,
  sendMessage: (msg: Message<unknown>) => void,
  switchShopScreen: () => void,
  clearBoard: () => void,
  readyUp: (sendMessage: (msg: Message<unknown>) => void) => void,
  unreadyUp: (sendMessage: (msg: Message<unknown>) => void) => void,
  drawUp: (sendMessage: (msg: Message<unknown>) => void) => void,
  undrawUp: (sendMessage: (msg: Message<unknown>) => void) => void,
  pressResign: () => void,
  cancelResign: () => void,
  confirmResign: (sendMessage: (msg: Message<unknown>) => void) => void,


): Map<string, Button> {
  const result: Map<string, Button> = new Map<string, Button>()

  const copyID = (gameID: string) => {
    navigator.clipboard.writeText(gameID);
    console.log("COPY")
    handleNotif("Text Copied")
  }

  const defaultButtonConfig = {
    x: 0,
    y: 0,
    width: 100 * UIRatio,
    height: 100 * UIRatio,
    strokeSize: 0 * UIRatio,
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
    text: "",
    subtext: "",
    onClick: () => { },
  }

  const idConfig = { ...defaultButtonConfig }
  const width = 600 * UIRatio
  const height = 100 * UIRatio

  idConfig.x = canvas.width / 2 - width / 2
  idConfig.y = canvas.height / 2 - height / 2
  idConfig.width = width
  idConfig.height = height
  idConfig.strokeSize = 5 * UIRatio
  idConfig.fontSize = 40 * UIRatio
  idConfig.text = game.id
  idConfig.subtext = "Click to copy ID"
  idConfig.onClick = () => copyID(game.id)
  const idButton = new Button(idConfig)
  result.set("id", idButton)

  const unitSwitchConfig = { ...defaultButtonConfig }
  unitSwitchConfig.x = 0
  unitSwitchConfig.y = 900 * UIRatio
  unitSwitchConfig.text = "Shop"
  unitSwitchConfig.subtext = "Click to switch"
  unitSwitchConfig.onClick = () => switchShopScreen()
  const unitSwitchButton = new Button(unitSwitchConfig)
  result.set("shop", unitSwitchButton)

  const readyConfig = { ...defaultButtonConfig }
  readyConfig.x = 0
  readyConfig.y = 800 * UIRatio
  readyConfig.text = "Ready"
  readyConfig.onClick = () => readyUp(sendMessage)
  const readyButton = new Button(readyConfig)
  result.set("ready", readyButton)

  const unreadyConfig = { ...readyConfig }
  unreadyConfig.text = "Unready"
  unreadyConfig.onClick = () => unreadyUp(sendMessage)
  const unreadyButton = new Button(unreadyConfig)
  result.set("unready", unreadyButton)

  const clearConfig = { ...defaultButtonConfig }
  clearConfig.x = 0
  clearConfig.y = 700 * UIRatio
  clearConfig.text = "Clear"
  clearConfig.onClick = () => clearBoard()
  const clearButton = new Button(clearConfig)
  result.set("clear", clearButton)

  //change clearBoard function
  const resignConfig = { ...defaultButtonConfig }
  resignConfig.x = 0
  resignConfig.y = 900 * UIRatio
  resignConfig.text = "Resign"
  resignConfig.onClick = () => pressResign()
  const resignButton = new Button(resignConfig)
  result.set("resign", resignButton)

  const confirmConfig = { ...defaultButtonConfig }
  confirmConfig.x = 0
  confirmConfig.y = 900 * UIRatio
  confirmConfig.height = 50 * UIRatio
  confirmConfig.text = "Confirm"
  confirmConfig.fontSize = 15 * UIRatio
  confirmConfig.onClick = () => confirmResign(sendMessage)
  const confirmButton = new Button(confirmConfig)
  result.set("confirm", confirmButton)

  const cancelConfig = { ...confirmConfig }
  cancelConfig.y = 950 * UIRatio
  cancelConfig.height = 50 * UIRatio
  cancelConfig.text = "Cancel"
  cancelConfig.onClick = () => cancelResign()
  const cancelButton = new Button(cancelConfig)
  result.set("cancel", cancelButton)

  //change clearBoard function
  const drawConfig = { ...defaultButtonConfig }
  drawConfig.x = 0
  drawConfig.y = 800 * UIRatio
  drawConfig.text = "Draw"
  drawConfig.onClick = () => drawUp(sendMessage)
  const drawButton = new Button(drawConfig)
  result.set("draw", drawButton)

  //change to undraw
  const undrawConfig = { ...drawConfig }
  undrawConfig.text = "Undraw"
  undrawConfig.onClick = () => undrawUp(sendMessage)
  const undrawButton = new Button(undrawConfig)
  result.set("undraw", undrawButton)

  return result
}


function createBlackButtons(
  canvas: HTMLCanvasElement,
  UIRatio: number,
  game: Game,
  handleNotif: (err: string) => void,
  sendMessage: (msg: Message<unknown>) => void,
  switchShopScreen: () => void,
  clearBoard: () => void,
  readyUp: (sendMessage: (msg: Message<unknown>) => void) => void,
  unreadyUp: (sendMessage: (msg: Message<unknown>) => void) => void,
  drawUp: (sendMessage: (msg: Message<unknown>) => void) => void,
  undrawUp: (sendMessage: (msg: Message<unknown>) => void) => void,
  pressResign: () => void,
  cancelResign: () => void,
  confirmResign: (sendMessage: (msg: Message<unknown>) => void) => void,


): Map<string, Button> {
  const result: Map<string, Button> = new Map<string, Button>()

  const copyID = (gameID: string) => {
    navigator.clipboard.writeText(gameID);
    console.log("COPY")
    handleNotif("Text Copied")
  }

  const defaultButtonConfig = {
    x: 0,
    y: 0,
    width: 100 * UIRatio,
    height: 100 * UIRatio,
    strokeSize: 0 * UIRatio,
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
    text: "",
    subtext: "",
    onClick: () => { },
  }

  const idConfig = { ...defaultButtonConfig }
  const width = 600 * UIRatio
  const height = 100 * UIRatio

  idConfig.x = canvas.width / 2 - width / 2
  idConfig.y = canvas.height / 2 - height / 2
  idConfig.width = width
  idConfig.height = height
  idConfig.strokeSize = 5 * UIRatio
  idConfig.fontSize = 40 * UIRatio
  idConfig.text = game.id
  idConfig.subtext = "Click to copy ID"
  idConfig.onClick = () => copyID(game.id)
  const idButton = new Button(idConfig)
  result.set("id", idButton)

  const unitSwitchConfig = { ...defaultButtonConfig }
  unitSwitchConfig.x = 900 * UIRatio
  unitSwitchConfig.y = 0
  unitSwitchConfig.text = "Shop"
  unitSwitchConfig.subtext = "Click to switch"
  unitSwitchConfig.onClick = () => switchShopScreen()
  const unitSwitchButton = new Button(unitSwitchConfig)
  result.set("shop", unitSwitchButton)

  const readyConfig = { ...defaultButtonConfig }
  readyConfig.x = 900 * UIRatio
  readyConfig.y = 100 * UIRatio
  readyConfig.text = "Ready"
  readyConfig.onClick = () => readyUp(sendMessage)
  const readyButton = new Button(readyConfig)
  result.set("ready", readyButton)

  const unreadyConfig = { ...readyConfig }
  unreadyConfig.text = "Unready"
  unreadyConfig.onClick = () => unreadyUp(sendMessage)
  const unreadyButton = new Button(unreadyConfig)
  result.set("unready", unreadyButton)

  const clearConfig = { ...defaultButtonConfig }
  clearConfig.x = 900 * UIRatio
  clearConfig.y = 200 * UIRatio
  clearConfig.text = "Clear"
  clearConfig.onClick = () => clearBoard()
  const clearButton = new Button(clearConfig)
  result.set("clear", clearButton)

  //change clearBoard function
  const resignConfig = { ...defaultButtonConfig }
  resignConfig.x = 900 * UIRatio
  resignConfig.y = 0
  resignConfig.text = "Resign"
  resignConfig.onClick = () => pressResign()
  const resignButton = new Button(resignConfig)
  result.set("resign", resignButton)

  const confirmConfig = { ...defaultButtonConfig }
  confirmConfig.x = 900 * UIRatio
  confirmConfig.y = 0 * UIRatio
  confirmConfig.height = 50 * UIRatio
  confirmConfig.text = "Confirm"
  confirmConfig.fontSize = 15 * UIRatio
  confirmConfig.onClick = () => confirmResign(sendMessage)
  const confirmButton = new Button(confirmConfig)
  result.set("confirm", confirmButton)

  const cancelConfig = { ...confirmConfig }
  cancelConfig.y = 50 * UIRatio
  cancelConfig.height = 50 * UIRatio
  cancelConfig.text = "Cancel"
  cancelConfig.onClick = () => cancelResign()
  const cancelButton = new Button(cancelConfig)
  result.set("cancel", cancelButton)

  //change clearBoard function
  const drawConfig = { ...defaultButtonConfig }
  drawConfig.x = 900 * UIRatio
  drawConfig.y = 100 * UIRatio
  drawConfig.text = "Draw"
  drawConfig.onClick = () => drawUp(sendMessage)
  const drawButton = new Button(drawConfig)
  result.set("draw", drawButton)

  //change to undraw
  const undrawConfig = { ...drawConfig }
  undrawConfig.text = "Undraw"
  undrawConfig.onClick = () => undrawUp(sendMessage)
  const undrawButton = new Button(undrawConfig)
  result.set("undraw", undrawButton)

  return result
}

