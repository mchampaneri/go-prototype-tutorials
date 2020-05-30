package main

import (
	"github.com/gonutz/prototype/draw"
)

func main() {

	// create application window
	height := 300
	width := 500

	draw.RunWindow("Sample - App | Hello world ", width, height,
		func(window draw.Window) {
			window.DrawText("Hello world", 0, 0, draw.White)
		},
	)

}
