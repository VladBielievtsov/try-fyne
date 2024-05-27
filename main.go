package main

import (
	"fyne.io/fyne/v2"
	"fyne.io/fyne/v2/app"
	"fyne.io/fyne/v2/canvas"
	"image/color"
)

func setContentToText(c fyne.Canvas) {
	green := color.NRGBA{R: 0, G: 255, B: 0, A: 255}
	text := canvas.NewText("Text", green)
	text.TextStyle.Bold = true
	c.SetContent(text)
}

func setContentToCircle(c fyne.Canvas) {
	red := color.NRGBA{R: 255, G: 0, B: 0, A: 255}
	circle := canvas.NewCircle(color.White)
	circle.StrokeWidth = 4
	circle.StrokeColor = red
	c.SetContent(circle)
}

func main() {
	a := app.New()
	w := a.NewWindow("Canvas")
	myCanvas := w.Canvas()

	//blue := color.NRGBA{B: 255, A: 255}
	//rect := canvas.NewRectangle(blue)
	//myCanvas.SetContent(rect)

	//setContentToText(myCanvas)
	setContentToCircle(myCanvas)

	w.Resize(fyne.NewSize(300, 300))
	w.ShowAndRun()
}
