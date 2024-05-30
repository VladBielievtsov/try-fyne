package main

import (
	"fyne.io/fyne/v2/app"
	"fyne.io/fyne/v2/container"
	"fyne.io/fyne/v2/widget"
	"log"
)

func main() {
	a := app.New()
	w := a.NewWindow("Hello")

	input := widget.NewEntry()
	input.SetPlaceHolder("Enter text...")

	tabs := container.NewAppTabs(
		container.NewTabItem("Tab 1", widget.NewLabel("Hello")),
		container.NewTabItem("Tab 2", container.NewVBox(input, widget.NewButton("Save", func() {
			log.Println("Content was: ", input.Text)
		}))),
	)

	tabs.SetTabLocation(container.TabLocationLeading)

	w.SetContent(tabs)
	w.ShowAndRun()
}
