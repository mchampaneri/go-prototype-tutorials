package main

import "github.com/gonutz/prototype/draw"

func main() {

	title := "Add png"
	height := 300
	width := 300

	draw.RunWindow(title, height, width, func(window draw.Window) {

		// add the sample.png at bottom of screen.
		window.DrawImageFile("./sample.png", 0, height-54)
	})
}
