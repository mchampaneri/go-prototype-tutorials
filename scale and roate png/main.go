package main

import "github.com/gonutz/prototype/draw"

func main() {

	title := "Scale position and rotate image"
	height := 300
	width := 300

	imageHeight := 100
	imageWidth := 100
	rotate := 0 // in deg
	draw.RunWindow(title, height, width, func(window draw.Window) {

		// do increase size and rotate till image size is less then screen height
		if imageHeight < 300 {
			// add image on stage
			// change height and width with frame
			imageWidth = imageWidth + 1
			imageHeight = imageHeight + 1

			// roate image
			rotate = rotate + 3
		}
		window.DrawImageFileTo("./sample.png", 0, 0, imageWidth, imageHeight, rotate)

	})
}
