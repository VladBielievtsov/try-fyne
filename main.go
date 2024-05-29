package main

import (
	"fyne.io/fyne/v2"
	"fyne.io/fyne/v2/app"
	"fyne.io/fyne/v2/canvas"
	"image/color"
	"math/rand"
)

func main() {
	a := app.New()
	w := a.NewWindow("Hello")

	raster := canvas.NewRasterWithPixels(func(_, _, w, h int) color.Color {
		return color.RGBA{R: uint8(rand.Intn(255)), G: uint8(rand.Intn(255)), B: uint8(rand.Intn(255)), A: 0xff}
	})

	w.SetContent(raster)
	w.Resize(fyne.NewSize(1280, 720))

	w.ShowAndRun()
}
