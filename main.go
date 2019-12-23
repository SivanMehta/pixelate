package main

import (
	"strconv"

	"./pixelate"
)

func main() {
	input := "bliss-4k.jpg"
	sizes := [...]int{20, 40, 48, 60, 80, 120, 240}

	for size := 0; size < len(sizes); size++ {
		pixelSize := sizes[size]
		output := "frames/bliss-" + strconv.Itoa(pixelSize) + ".jpg"
		pixelate.Pixelate(input, output, pixelSize)
	}
}
