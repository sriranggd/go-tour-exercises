package main

import (
	"image"
	"image/color"
)

type Image struct {
	height int
	width  int
	pixels [][]uint8
}

func (i Image) Bounds() image.Rectangle {
	return image.Rect(0, 0, i.width, i.height)
}

func (i Image) ColorModel() color.Model {
	return color.RGBAModel
}

func (i Image) At(x, y int) color.Color {
	if x > i.width || y > i.height {
		return nil
	}

	return color.RGBA{R: i.pixels[x][y], G: i.pixels[x][y], B: 255, A: 255}
}
