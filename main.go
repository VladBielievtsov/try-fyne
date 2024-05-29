package main

import (
	"fyne.io/fyne/v2/app"
	"fyne.io/fyne/v2/canvas"
	"fyne.io/fyne/v2/container"
	"fyne.io/fyne/v2/layout"
	"image/color"
)

func main() {
	a := app.New()
	w := a.NewWindow("Hello")

	text1 := canvas.NewText("1", color.White)
	text2 := canvas.NewText("2", color.White)
	text3 := canvas.NewText("3", color.White)

	grid := container.New(layout.NewGridLayout(2), text1, text2, text3)

	w.SetContent(grid)
	w.ShowAndRun()
}
