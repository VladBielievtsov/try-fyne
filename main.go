package main

import (
	"fyne.io/fyne/v2"
	"fyne.io/fyne/v2/app"
	"fyne.io/fyne/v2/container"
	"fyne.io/fyne/v2/layout"
	"fyne.io/fyne/v2/theme"
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

	// Tab 4

	toolBar := widget.NewToolbar(
		widget.NewToolbarAction(theme.DocumentCreateIcon(), func() {
			log.Println("New Document")
		}),
		widget.NewToolbarSeparator(),
		widget.NewToolbarAction(theme.ContentCutIcon(), func() {}),
		widget.NewToolbarAction(theme.ContentCopyIcon(), func() {}),
		widget.NewToolbarAction(theme.ContentPasteIcon(), func() {}),
		widget.NewToolbarSpacer(),
		widget.NewToolbarAction(theme.HelpIcon(), func() {
			log.Println("Display help")
		}),
	)

	// Tab 5

	data := [][]string{[]string{"top left", "top right"},
		[]string{"bottom left", "bottom right"}}

	list := widget.NewTable(
		func() (int, int) {
			return len(data), len(data[0])
		},
		func() fyne.CanvasObject {
			return widget.NewLabel("template")
		},
		func(i widget.TableCellID, o fyne.CanvasObject) {
			o.(*widget.Label).SetText(data[i.Row][i.Col])
		},
	)

	tabs := container.NewAppTabs(
		container.NewTabItem("Tab 1", container.NewVBox(input, check, radio, combo, widget.NewButton("Save", func() {
			log.Println("Content was: ", input.Text)
		}))),
		container.NewTabItem("Tab 2", container.NewVBox(widget.NewLabel("Hello"), form)),
		container.NewTabItem("Tab 3", container.NewVBox(progress, infinite)),
		container.NewTabItem("Tab 4", container.NewBorder(toolBar, nil, nil, nil, widget.NewLabel("Content"))),
		container.NewTabItem("Tab 5", container.New(layout.NewMaxLayout(), list)),
	)

	tabs.SetTabLocation(container.TabLocationLeading)

	w.SetContent(tabs)
	w.Resize(fyne.NewSize(450, 400))
	w.ShowAndRun()
}
