package main

import (
	"fyne.io/fyne/v2"
	"fyne.io/fyne/v2/app"
	"fyne.io/fyne/v2/canvas"
	"image/color"
)

func main() {
	a := app.New()
	w := a.NewWindow("Hello")

	gradient := canvas.NewHorizontalGradient(color.White, color.Transparent)

	w.SetContent(gradient)
	w.Resize(fyne.NewSize(1280, 720))

	w.ShowAndRun()
}
