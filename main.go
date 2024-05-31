package main

import (
	"fmt"
	"image/color"
	"strings"

	"fyne.io/fyne/v2"
	"fyne.io/fyne/v2/app"
	"fyne.io/fyne/v2/canvas"
	"fyne.io/fyne/v2/container"
	"fyne.io/fyne/v2/dialog"
	"fyne.io/fyne/v2/theme"
	"fyne.io/fyne/v2/widget"
)

func main() {
	a := app.New()
	a.Settings().SetTheme(theme.DarkTheme())
	w := a.NewWindow("Image Converter")
	w.Resize(fyne.NewSize(500, 500))
	w.CenterOnScreen()

	labelFile := widget.NewLabel("File:")
	entryFile := widget.NewEntry()
	buttonSelect := widget.NewButton("Select image", func() {
		dialog.ShowFileOpen(func(read fyne.URIReadCloser, err error) {
			if err != nil {
				dialog.ShowError(err, w)
				return
			}
			if read == nil {
				return
			}
			selectedFile := read.URI().String()
			filePath := strings.TrimPrefix(selectedFile, "file://")
			entryFile.SetText(selectedFile)
			labelFile.SetText(fmt.Sprintf("File: \"%s\"", filePath))
		}, w)
	})

	chooseImage := container.NewVBox(labelFile, buttonSelect)

	settings := container.NewCenter(widget.NewLabel("Setting"))

	content := container.NewVBox(chooseImage, canvas.NewLine(color.White), settings)

	w.SetContent(content)

	w.Show()
	a.Run()
}
