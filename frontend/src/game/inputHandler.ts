import { type Mouse } from "./util.ts"

export class InputHandler {
  mouse: Mouse = {
    x: 0,
    y: 0,
    pressed: [false, false, false],
    prevPressed: [false, false, false],
    justPressed: [false, false, false],
    justReleased: [false, false, false]
  }

  constructor(canvas: HTMLCanvasElement) {
    const handleMouseMove = (e: MouseEvent) => {
      const rect = canvas.getBoundingClientRect()
      this.mouse.x = e.clientX - rect.left
      this.mouse.y = e.clientY - rect.top
    }

    const handleMouseDown = (event: MouseEvent) => {
      if (event.button <= 2) {
        this.mouse.pressed[event.button] = true
      }
    }

    const handleMouseUp = (event: MouseEvent) => {
      if (event.button <= 2) {
        this.mouse.pressed[event.button] = false
      }
    }

    const handleContextMenu = (event: MouseEvent) => {
      event.preventDefault()
    }


    canvas.addEventListener("mousemove", handleMouseMove)
    canvas.addEventListener("mousedown", handleMouseDown)
    canvas.addEventListener("mouseup", handleMouseUp)
    document.addEventListener("contextmenu", handleContextMenu)

    this.cleanup = () => {
      canvas.removeEventListener("mousemove", handleMouseMove)
      canvas.removeEventListener("mousedown", handleMouseDown)
      window.removeEventListener("mouseup", handleMouseUp)
      document.removeEventListener("contextmenu", handleContextMenu)
    }
  }

  update() {
    for (let i = 0; i < 3; i++) {
      this.mouse.justPressed[i] = this.mouse.pressed[i] && !this.mouse.prevPressed[i]
      this.mouse.justReleased[i] = !this.mouse.pressed[i] && this.mouse.prevPressed[i]

      this.mouse.prevPressed[i] = this.mouse.pressed[i]

    }
  }

  cleanup: () => void;
}
