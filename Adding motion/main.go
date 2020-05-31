package main

import "github.com/gonutz/prototype/draw"

func main() {

	title := "Text in motion"
	height := 300
	width := 300
	i := 0
	draw.RunWindow(title, height, width, func(window draw.Window) {
		// let make this ellipse living
		i++
		// with every redraw parameter of ellipse in changing which gets
		// refelected in redraw result

		// let's change colour as well
		// adding filled ellipse
		window.FillEllipse(i%height, i%width, i%width, i%height, draw.Color{float32(i % 7),
			float32(i % 11),
			float32(i % 17),
			1,
		})
	})
}
