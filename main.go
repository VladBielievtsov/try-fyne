package main

import (
	"fyne.io/fyne/v2"
	"fyne.io/fyne/v2/app"
	"fyne.io/fyne/v2/canvas"
	"fyne.io/fyne/v2/container"
	"image/color"
	"time"
)

func main() {
	a := app.New()
	w := a.NewWindow("Hello")

	obj := canvas.NewRectangle(color.Black)
	obj.Resize(fyne.NewSize(50, 50))
	w.SetContent(container.NewWithoutLayout(obj))

	red := color.NRGBA{R: 0xff, A: 0xff}
	blue := color.NRGBA{B: 0xff, A: 0xff}
	canvas.NewColorRGBAAnimation(red, blue, time.Second*2, func(c color.Color) {
		obj.FillColor = c
		canvas.Refresh(obj)
	}).Start()

	move := canvas.NewPositionAnimation(fyne.NewPos(0, 0), fyne.NewPos(200, 0), time.Second, obj.Move)
	move.AutoReverse = true
	move.Start()

	w.Resize(fyne.NewSize(250, 50))
	w.SetPadded(false)
	w.ShowAndRun()
}
