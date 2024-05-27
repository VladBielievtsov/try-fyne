package main

import (
	"fyne.io/fyne/v2"
	"fyne.io/fyne/v2/app"
	"fyne.io/fyne/v2/widget"
	"time"
)

func updateTime(clock *widget.Label) {
	formatted := time.Now().Format("Time: 15:04:05")
	clock.SetText(formatted)
}

func main() {
	a := app.New()
	w := a.NewWindow("Clock")
	w.Resize(fyne.NewSize(1280, 720))
	clock := widget.NewLabel("")
	updateTime(clock)
	w.SetContent(clock)

	w2 := a.NewWindow("Second")
	w2.Resize(fyne.NewSize(450, 300))
	w2.SetContent(clock)

	go func() {
		for range time.Tick(time.Second) {
			updateTime(clock)
		}
	}()

	w.Show()
	w2.Show()
	a.Run()
}
