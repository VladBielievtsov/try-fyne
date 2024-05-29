package main

import (
	"fyne.io/fyne/v2/app"
	"fyne.io/fyne/v2/canvas"
	"fyne.io/fyne/v2/container"
	"image/color"
)

func main() {
	a := app.New()
	w := a.NewWindow("Hello")

	top := canvas.NewText("top bar", color.White)
	left := canvas.NewText("left", color.White)
	middle := canvas.NewText("center", color.White)

	content := container.NewBorder(top, nil, left, nil, middle)

	w.SetContent(content)
	w.ShowAndRun()
}
