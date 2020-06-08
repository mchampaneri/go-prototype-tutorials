package main

import "github.com/gonutz/prototype/draw"

func main() {

	title := "Text in motion"
	height := 300
	width := 300

	draw.RunWindow(title, height, width, func(window draw.Window) {

		window.DrawImageFile("./sample.png", 10, height-50)
	})
}
