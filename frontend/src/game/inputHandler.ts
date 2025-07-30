import { Button } from "./button.ts"
import { type Mouse } from "./util.ts"

export class InputHandler {
  mouse: Mouse = { x: 0, y: 0, pressed: false }

  constructor(canvas: HTMLCanvasElement, buttons: Button[]) {
    const handleMouseMove = (e: MouseEvent) => {
      const rect = canvas.getBoundingClientRect()
      this.mouse.x = e.clientX - rect.left
      this.mouse.y = e.clientY - rect.top

      for (const button of buttons) {
        if (!button.visible) {
          return
        }
        button.checkHoveringButton(this.mouse)
      }
    }

    const handleMouseDown = () => {
      this.mouse.pressed = true

      for (const button of buttons) {
        if (!button.visible) {
          return
        }
        button.clickButton(this.mouse)
      }
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


  cleanup: () => void;
}
