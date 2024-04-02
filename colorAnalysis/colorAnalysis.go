package colorAnalysis

import (
	"image/color"

	"fyne.io/fyne/v2/canvas"
)

func CalculateAverageColor(img *canvas.Image) (r uint32, g uint32, b uint32, a uint32) {

	width := int(img.MinSize().Width)
	height := int(img.MinSize().Height)
	var rMoyenne uint32
	var gMoyenne uint32
	var bMoyenne uint32
	var aMoyenne uint32
	var j uint32

	for y := 0; y < height; y++ {
		for x := 0; x < width; x++ {

			colorrr := img.Image.At(x, y)
			r, g, b, a := colorToRGB(colorrr)
			rMoyenne += r
			gMoyenne += g
			bMoyenne += b
			aMoyenne += a
			j++

		}
	}
	rMoyenne = rMoyenne / j
	gMoyenne = gMoyenne / j
	bMoyenne = bMoyenne / j
	aMoyenne = aMoyenne / j

	return rMoyenne, gMoyenne, bMoyenne, aMoyenne

}

func colorToRGB(c color.Color) (r, g, b, a uint32) {
	r, g, b, a = c.RGBA()
	r = r / 257
	g = g / 257
	b = b / 257
	a = a / 257
	return r, g, b, a
}
