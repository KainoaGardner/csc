export const Sounds = new Map<string, HTMLAudioElement>()

Sounds.set("error", new Audio("sounds/error.mp3"))
Sounds.set("notif", new Audio("sounds/notif.mp3"))
Sounds.set("place", new Audio("sounds/place.mp3"))

export function playAudio(audio: HTMLAudioElement) {
  audio.pause()
  audio.currentTime = 0
  audio.play()
}

export function updateVolume(volume: number, k = 4) {
  const x = Math.min(Math.max(volume / 100, 0), 1)

  let y: number
  if (k === 0) {
    y = x
  } else {
    y = (Math.exp(k * x) - 1) / (Math.exp(k) - 1)
  }

  for (const [, value] of Sounds) {
    value.volume = y
  }
}
