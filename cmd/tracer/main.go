package main

import (
	goimage "image"
	gocolor "image/color"
	"image/png"
	"os"
)

func main() {
	println("Hello bob")

	// Image

	image_width := 256
	image_height := 256
	img := goimage.NewRGBA64(goimage.Rect(0, 0, image_width, image_height))

	// Render

	for j := 0; j < image_height; j++ {
		for i := 0; i < image_width; i++ {
			r := float64(i) / float64(image_width-1)
			g := float64(j) / float64(image_height-1)
			b := 0.0

			color := gocolor.RGBA64{
				R: uint16(float64(0xfffe) * r),
				G: uint16(float64(0xfffe) * g),
				B: uint16(float64(0xfffe) * b),
				A: 0xffff,
			}
			img.SetRGBA64(i, j, color)
		}
	}

	f, err := os.Create("out.png")
	if err != nil {
		panic(err)
	}

	err = png.Encode(f, img)
	if err != nil {
		panic("Failed to encode image")
	}

	err = f.Close()
	if err != nil {
		panic("Failed to close file")
	}
}
