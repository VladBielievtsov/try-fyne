package main

import (
	"fyne.io/fyne/v2/app"
	"fyne.io/fyne/v2/canvas"
	"fyne.io/fyne/v2/container"
	"fyne.io/fyne/v2/layout"
	"fyne.io/fyne/v2/theme"
	"image/color"
)

func main() {
	a := app.New()
	w := a.NewWindow("Hello")

	img := canvas.NewImageFromResource(theme.FyneLogo())
	img.FillMode = canvas.ImageFillOriginal
	text := canvas.NewText("Overlay", color.White)
	content := container.New(layout.NewCenterLayout(), img, text)

	w.SetContent(content)
	w.ShowAndRun()
}
