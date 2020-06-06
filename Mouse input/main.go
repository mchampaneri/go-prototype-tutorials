package main

import "github.com/gonutz/prototype/draw"

func main() {

	title := "Caputre mouse input"
	height := 300
	width := 300

	// Defining rect
	rectXCord := 0
	rectYCord := 0
	rectHeight := 30
	rectWidth := 60

	draw.RunWindow(title, height, width, func(window draw.Window) {

		// lets move this rect with the
		// movement of mouse pointer

		rectXCord, rectYCord = window.MousePosition()

		// let's try to put poiner in middle of
		// rect
		rectXCord = rectXCord - 30
		rectYCord = rectYCord - 15

		// render rect on frame
		window.FillRect(rectXCord, rectYCord, rectWidth, rectHeight, draw.Cyan)
		// Awesome  .......
	})
}
