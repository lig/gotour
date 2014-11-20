package main

import (
	"code.google.com/p/go-tour/pic"
	"image"
	"image/color"
    "math"
)

type Image struct {
	w int
	h int
}

func (i Image) ColorModel() color.Model {
	return color.RGBAModel
}

func (i Image) Bounds() image.Rectangle {
	return image.Rect(0, 0, i.w, i.h)
}

func (i Image) At(x, y int) color.Color {
	var v uint8
	if (y != 0) && (x != 0) {
		v = uint8(256*math.Cos(float64(x/2)) + 256*math.Sin(float64(y/2)))
	}
	return color.RGBA{v, v, 255, 255}
}

func main() {
	m := Image{w: 512, h: 256}
	pic.ShowImage(m)
}
