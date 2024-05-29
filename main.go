package main

import (
	"fyne.io/fyne/v2"
	"fyne.io/fyne/v2/app"
	"fyne.io/fyne/v2/canvas"
	"fyne.io/fyne/v2/storage"
	"log"
)

func main() {
	a := app.New()
	w := a.NewWindow("Hello")

	url := "https://avatars.githubusercontent.com/u/40768241?v=4"

	uri, err := storage.ParseURI(url)
	if err != nil {
		log.Fatal(err)
		return
	}

	image := canvas.NewImageFromURI(uri)

	image.FillMode = canvas.ImageFillOriginal
	w.SetContent(image)
	w.Resize(fyne.NewSize(1280, 720))

	w.ShowAndRun()
}
