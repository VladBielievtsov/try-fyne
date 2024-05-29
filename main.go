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

	line := canvas.NewLine(color.White)
	line.StrokeWidth = 5
	w.SetContent(line)
	w.Resize(fyne.NewSize(1280, 720))

	w.ShowAndRun()
}
