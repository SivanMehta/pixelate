package pixelate

import (
	"image"
	"log"
	"os"

	"image/color"
	"image/draw"
	"image/jpeg"
)

// Pixelate a given image down to squares of a given size
func Pixelate(input string, output string, pixelSize int) {
	// Decode the JPEG data. If reading from file, create a reader with

	reader, err := os.Open(input)
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
			cell := newRow*width + newCol

			r, g, b, a := m.At(col, row).RGBA()

			reds[cell] += r
			greens[cell] += g
			blues[cell] += b
			alphas[cell] += a
		}
	}

	img := image.NewRGBA(image.Rect(0, 0, bounds.Max.X, bounds.Max.Y))

	for pixel := 0; pixel < height*width; pixel++ {
		col := color.RGBA{
			uint8(((reds[pixel] / pixelsPerCell) >> 8)),
			uint8(((greens[pixel] / pixelsPerCell) >> 8)),
			uint8(((blues[pixel] / pixelsPerCell) >> 8)),
			uint8(((alphas[pixel] / pixelsPerCell) >> 8)),
		}

		x1 := (pixel % width) * pixelSize
		y1 := (pixel / width) * pixelSize
		x2 := x1 + pixelSize
		y2 := y1 + pixelSize

		rectangle := image.Rect(x1, y1, x2, y2)
		draw.Draw(img, rectangle, &image.Uniform{col}, image.ZP, draw.Src)
	}

	myfile, err := os.Create(output) // ... now lets save imag
	if err != nil {
		panic(err)
	}
	jpeg.Encode(myfile, img, nil)
}
