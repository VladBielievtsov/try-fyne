package main

import (
	"fyne.io/fyne/v2"
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

	check := widget.NewCheck("Optional", func(value bool) {
		log.Println("Check set to: ", value)
	})

	radio := widget.NewRadioGroup([]string{"Option 1", "Option 2"}, func(value string) {
		log.Println("Radio set to: ", value)
	})

	combo := widget.NewSelect([]string{"Option 1", "Option 2"}, func(value string) {
		log.Println("Combo set to: ", value)
	})

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

	tabs := container.NewAppTabs(
		container.NewTabItem("Tab 1", container.NewVBox(input, check, radio, combo, widget.NewButton("Save", func() {
			log.Println("Content was: ", input.Text)
		}))),
		container.NewTabItem("Tab 2", container.NewVBox(widget.NewLabel("Hello"), form)),
	)

	tabs.SetTabLocation(container.TabLocationLeading)

	w.SetContent(tabs)
	w.Resize(fyne.NewSize(450, 400))
	w.ShowAndRun()
}
