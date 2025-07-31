import { type Mouse } from "./util.ts"

export class InputHandler {
  mouse: Mouse = { x: 0, y: 0, pressed: false, prevPressed: false, justPressed: false, justReleased: false }

  constructor(canvas: HTMLCanvasElement) {
    const handleMouseMove = (e: MouseEvent) => {
      const rect = canvas.getBoundingClientRect()
      this.mouse.x = e.clientX - rect.left
      this.mouse.y = e.clientY - rect.top
    }

    const handleMouseDown = () => {
      this.mouse.pressed = true

    }

    const handleMouseUp = () => {
      this.mouse.pressed = false
    }

    canvas.addEventListener("mousemove", handleMouseMove)
    canvas.addEventListener("mousedown", handleMouseDown)
    canvas.addEventListener("mouseup", handleMouseUp)

    this.cleanup = () => {
      canvas.removeEventListener("mousemove", handleMouseMove)
      canvas.removeEventListener("mousedown", handleMouseDown)
      window.removeEventListener("mouseup", handleMouseUp)
    }
  }

  update() {
    this.mouse.justPressed = this.mouse.pressed && !this.mouse.prevPressed
    this.mouse.justReleased = !this.mouse.pressed && this.mouse.prevPressed
    this.mouse.prevPressed = this.mouse.pressed
  }


  cleanup: () => void;
}
