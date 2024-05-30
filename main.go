package main

import (
	"fyne.io/fyne/v2"
	"fyne.io/fyne/v2/app"
	"fyne.io/fyne/v2/container"
	"fyne.io/fyne/v2/widget"
	"log"
	"time"
)

func main() {
	a := app.New()
	w := a.NewWindow("Hello")

	// Tab 1

	input := widget.NewEntry()
	input.SetPlaceHolder("Enter text...")

	check := widget.NewCheck("Optional", func(value bool) {
		log.Println("Check set to: ", value)
	})

	radio := widget.NewRadioGroup([]string{"Option 1", "Option 2"}, func(value string) {
		log.Println("Radio set to: ", value)
	})

	combo := widget.NewSelect([]string{"Option 1", "Option 2"}, func(value string) {
		log.Println("Combo set to: ", value)
	})

	// Tab 2

	entry := widget.NewEntry()
	textArea := widget.NewMultiLineEntry()

	form := &widget.Form{
		Items: []*widget.FormItem{
			{Text: "Input", Widget: entry},
		},
		OnSubmit: func() {
			log.Println("Form submitted: ", entry.Text)
			log.Println("Multiline: ", textArea.Text)
		},
	}

	form.Append("Text", textArea)

	// Tab 3

	progress := widget.NewProgressBar()
	infinite := widget.NewProgressBarInfinite()

	go func() {
		for i := 0.0; i <= 1.0; i += 0.1 {
			time.Sleep(time.Millisecond * 250)
			progress.SetValue(i)
		}
	}()

	tabs := container.NewAppTabs(
		container.NewTabItem("Tab 1", container.NewVBox(input, check, radio, combo, widget.NewButton("Save", func() {
			log.Println("Content was: ", input.Text)
		}))),
		container.NewTabItem("Tab 2", container.NewVBox(widget.NewLabel("Hello"), form)),
		container.NewTabItem("Tab 3", container.NewVBox(progress, infinite)),
	)

	tabs.SetTabLocation(container.TabLocationLeading)

	w.SetContent(tabs)
	w.Resize(fyne.NewSize(450, 400))
	w.ShowAndRun()
}
