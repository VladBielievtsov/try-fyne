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

	circle := canvas.NewCircle(color.White)
	circle.StrokeColor = color.Gray{Y: 0x99}
	w.SetContent(circle)
	w.Resize(fyne.NewSize(1280, 720))

	w.ShowAndRun()
}
