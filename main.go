package main

import (
	"fyne.io/fyne/v2"
	"fyne.io/fyne/v2/app"
	"fyne.io/fyne/v2/canvas"
	"image/color"
)

func main() {
	a := app.New()
	w := a.NewWindow("Canvas")
	myCanvas := w.Canvas()

	blue := color.NRGBA{B: 255, A: 255}
	rect := canvas.NewRectangle(blue)
	myCanvas.SetContent(rect)

	w.Resize(fyne.NewSize(300, 300))
	w.ShowAndRun()
}
