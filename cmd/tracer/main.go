package main

import (
	goimage "image"
	// gocolor "image/color"
	"image/png"
	"os"
)

func main() {
	println("Hello bob")

	width := 200
	height := 300
	img := goimage.NewRGBA64(goimage.Rect(0, 0, width, height))

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
