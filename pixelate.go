package main

import (
	"image"
	"log"
	"os"

	// Package image/jpeg is not used explicitly in the code below,
	// but is imported for its initialization side-effect, which allows
	// image.Decode to understand JPEG formatted images. Uncomment these
	// two lines to also understand GIF and PNG images:
	// _ "image/gif"
	// _ "image/png"
	_ "image/jpeg"
)

const pixelSize = 10

func main() {
	// Decode the JPEG data. If reading from file, create a reader with
	//
	reader, err := os.Open("bliss-4k.jpg")
	if err != nil {
		log.Fatal(err)
	}
	defer reader.Close()
	m, _, err := image.Decode(reader)
	if err != nil {
		log.Fatal(err)
	}
	bounds := m.Bounds()

	height := bounds.Max.Y / pixelSize
	width := bounds.Max.X / pixelSize
	reds := make([]uint32, height*width)
	greens := make([]uint32, height*width)
	blues := make([]uint32, height*width)
	alphas := make([]uint32, height*width)
	pixelsPerCell := uint32(pixelSize * pixelSize)

	for row := bounds.Min.Y; row < bounds.Max.Y; row++ {
		for col := bounds.Min.X; col < bounds.Max.X; col++ {
			// returns 32 bit color
			newRow := row / pixelSize
			newCol := col / pixelSize
			cell := newCol*width + newRow

			r, g, b, a := m.At(col, row).RGBA()

			reds[cell] += r
			greens[cell] += g
			blues[cell] += b
			alphas[cell] += a
		}
	}

	for pixel := 0; pixel < height*width; pixel++ {
		reds[pixel] = reds[pixel] / pixelsPerCell
		greens[pixel] = greens[pixel] / pixelsPerCell
		blues[pixel] = blues[pixel] / pixelsPerCell
		alphas[pixel] = alphas[pixel] / pixelsPerCell
	}
}
