package main

import (
	"log"
	"time"

	"github.com/gonutz/prototype/draw"
)

func main() {

	title := "Add png"
	height := 800
	width := 600

	spriteHeight := 210
	spriteWidth := 120

	// let give it motion
	i := 1
	j := 1
	draw.RunWindow(title, height, width, func(window draw.Window) {

		// logic for shifting to next row
		if i%6 == 0 {
			j++
			i = 1
		}
		time.Sleep(100 * time.Millisecond)
		if err := window.DrawImageFilePart("./sample.png",
			(spriteWidth * (i % 6)), spriteHeight*(j%4), spriteWidth, spriteHeight,
			spriteWidth*i, 0, spriteWidth, spriteHeight,
			0,
		); err == nil {
			// nothing
		} else {
			log.Println("could not load image")
		}

		//increase i
		i++

	})

}
