package main

import (
	"fyne.io/fyne/v2"
	"fyne.io/fyne/v2/app"
	"fyne.io/fyne/v2/canvas"
	"image/jpeg"
	"log"
	"os"
)

func main() {
	a := app.New()
	w := a.NewWindow("Hello")

	//raster := canvas.NewRasterWithPixels(func(_, _, w, h int) color.Color {
	//	return color.RGBA{R: uint8(rand.Intn(255)), G: uint8(rand.Intn(255)), B: uint8(rand.Intn(255)), A: 0xff}
	//})

	catFile, err := os.Open("url")
	if err != nil {
		log.Fatal(err)
	}

	img, err := jpeg.Decode(catFile)
	if err != nil {
		log.Fatal(err)
	}

	raster := canvas.NewRasterFromImage(img)
	w.SetContent(raster)
	w.Resize(fyne.NewSize(1280, 720))

	w.ShowAndRun()
}
