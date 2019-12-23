package draw

import (
	"image"
	"image/color"
)

// Bounds of a rectangle on which we want to draw
type Bounds struct {
	X1 int
	Y1 int
	X2 int
	Y2 int
}

// Rect draws a rectangle utilizing HLine() and VLine()
func Rect(bounds Bounds, img *image.Rectangle, col color.Color) {
	for row := bounds.Y1; row < bounds.Y2; row++ {
		for column := bounds.X1; column < bounds.X2; column++ {
			img.Set(row, column, col)
		}
	}
}
