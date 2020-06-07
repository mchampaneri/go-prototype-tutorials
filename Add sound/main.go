package main

import "github.com/gonutz/prototype/draw"

func main() {

	title := "Text in motion"
	height := 300
	width := 300

	draw.RunWindow(title, height, width, func(window draw.Window) {

		window.DrawText("Press space to play sound",
			10,
			height/2,
			draw.Cyan)

		if window.IsKeyDown(draw.KeySpace) {
			window.PlaySoundFile("./sound.wav")
		}
	})
}
