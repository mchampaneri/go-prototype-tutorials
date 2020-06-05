package main

import "github.com/gonutz/prototype/draw"

func main() {

	title := "Text in motion"
	height := 300
	width := 300

	CurrentX := width / 2
	CurrentY := height / 2
	draw.RunWindow(title, height, width, func(window draw.Window) {

		if window.IsKeyDown(draw.KeyDown) {
			// Increase the Y so it will look like going down
			CurrentY += 1
		} else if window.IsKeyDown(draw.KeyUp) {
			// Decrse the Y so it will look like going up
			CurrentY -= 1
		} else if window.IsKeyDown(draw.KeyLeft) {
			// Decrse the X so it will look like going left
			CurrentX -= 1
		} else if window.IsKeyDown(draw.KeyRight) {
			// Increase the X so it will look like going right
			CurrentX += 1
		}

		// adding a static reactangle
		window.FillRect(CurrentX, CurrentY, 50, 60, draw.Blue)
	})
}
