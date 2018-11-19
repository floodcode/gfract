package main

import (
	"image"
	"image/color"
	"image/png"
	"os"
)

const filename = "fract.png"

const width = 512
const height = 512

func main() {
	img := makeImage()
	file, err := os.Create(filename)
	if err != nil {
		panic(err.Error())
	}

	png.Encode(file, img)
}

func makeImage() *image.RGBA {
	img := image.NewRGBA(image.Rectangle{
		image.Point{0, 0},
		image.Point{width, height},
	})

	for x := 0; x < width; x++ {
		for y := 0; y < height; y++ {
			pix := gen(uint8(x), uint8(y))
			img.SetRGBA(x, y, color.RGBA{0, pix, 0, 0xff})
		}
	}

	return img
}

// The place where the magic begins.
// Just try to modify this function in the way you want.
func gen(x uint8, y uint8) uint8 {
	return x ^ y
}
