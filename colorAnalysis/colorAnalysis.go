package colorAnalysis

import (
	"image/color"

	"fyne.io/fyne/v2/canvas"
)

func CalculateAverageColor(img *canvas.Image) (r uint32, g uint32, b uint32, a uint32) {
	//Declaration of size and color variables of the image

	width := int(img.MinSize().Width)
	height := int(img.MinSize().Height)
	var rAverage uint32
	var gAverage uint32
	var bAverage uint32
	var aAverage uint32
	var totalPixel uint32

	//Analyze the color of each pixel
	for y := 0; y < height; y++ {
		for x := 0; x < width; x++ {

			colorPixel := img.Image.At(x, y)
			// Transform the type of color value from color.Color to RGB
			r, g, b, a := colorToRGB(colorPixel)

			//Add the value of r,g,b for one pixel for the average
			rAverage += r
			gAverage += g
			bAverage += b
			aAverage += a
			totalPixel++

		}
	}

	//Calculate the average color
	rAverage = rAverage / totalPixel
	gAverage = gAverage / totalPixel
	bAverage = bAverage / totalPixel
	aAverage = aAverage / totalPixel

	return rAverage, gAverage, bAverage, aAverage

}

func colorToRGB(c color.Color) (r, g, b, a uint32) {

	//Transform a color.Color value into a RGBA one
	r, g, b, a = c.RGBA()
	r = r / 257
	g = g / 257
	b = b / 257
	a = a / 257
	return r, g, b, a
}
