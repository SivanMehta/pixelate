package main

import (
	"./pixelate"
)

func main() {
	input := "bliss-4k.jpg"
	pixelate.Pixelate(input, "frames/bliss-240.jpg", 240)
	pixelate.Pixelate(input, "frames/bliss-120.jpg", 120)
	pixelate.Pixelate(input, "frames/bliss-80.jpg", 80)
	pixelate.Pixelate(input, "frames/bliss-60.jpg", 60)
	pixelate.Pixelate(input, "frames/bliss-48.jpg", 48)
	pixelate.Pixelate(input, "frames/bliss-40.jpg", 40)
	pixelate.Pixelate(input, "frames/bliss-20.jpg", 20)
}
