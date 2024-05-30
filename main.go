package main

import (
	"fmt"
	"fyne.io/fyne/v2"
	"fyne.io/fyne/v2/app"
	"fyne.io/fyne/v2/canvas"
	"fyne.io/fyne/v2/container"
	"fyne.io/fyne/v2/dialog"
	"fyne.io/fyne/v2/layout"
	"fyne.io/fyne/v2/theme"
	"fyne.io/fyne/v2/widget"
	"gorm.io/driver/sqlite"
	"gorm.io/gorm"
	"image/color"
)

type Task struct {
	ID          uint
	Title       string
	Description string
}

func main() {
	a := app.New()
	a.Settings().SetTheme(theme.DarkTheme())
	w := a.NewWindow("Todo")
	w.Resize(fyne.NewSize(500, 600))
	w.CenterOnScreen()

	var tasks []Task
	var createContent *fyne.Container
	var tasksList *widget.List
	var content *fyne.Container

	DB, _ := gorm.Open(sqlite.Open("todo.db"), &gorm.Config{})
	DB.AutoMigrate(&Task{})
	DB.Find(&tasks)

	noTasksLabel := widget.NewLabel("No Tasks")
	if len(tasks) != 0 {
		noTasksLabel.Hide()
	}
	tasksBar := container.NewHBox(
		widget.NewLabel("Your Tasks"),
		layout.NewSpacer(),
		widget.NewButtonWithIcon("", theme.ContentAddIcon(), func() {
			w.SetContent(createContent)
		}),
	)

	tasksList = widget.NewList(
		func() int {
			return len(tasks)
		}, func() fyne.CanvasObject {
			return widget.NewLabel("Default")
		}, func(id widget.ListItemID, object fyne.CanvasObject) {
			object.(*widget.Label).SetText(tasks[id].Title)
		},
	)

	tasksList.OnSelected = func(id widget.ListItemID) {
		detailsBar := container.NewHBox(
			widget.NewLabel(
				fmt.Sprintf(
					"Details about \"%s\"",
					tasks[id].Title,
				),
			),
			layout.NewSpacer(),
			widget.NewButtonWithIcon("", theme.CancelIcon(), func() {
				w.SetContent(content)
				tasksList.Unselect(id)
			}),
		)

		taskTitle := widget.NewLabel(tasks[id].Title)
		taskTitle.TextStyle = fyne.TextStyle{Bold: true}
		taskDescription := widget.NewLabel(tasks[id].Description)
		taskDescription.TextStyle = fyne.TextStyle{Italic: true}
		taskDescription.Wrapping = fyne.TextWrapWord

		buttonsBox := container.NewHBox(
			// Delete
			widget.NewButtonWithIcon("", theme.DeleteIcon(), func() {
				dialog.ShowConfirm(
					"Delete task",
					fmt.Sprintf("Delete task \"%s\"?", tasks[id].Title),
					func(ok bool) {
						if ok {
							DB.Delete(&Task{}, "id = ?", tasks[id].ID)
							DB.Find(&tasks)
							if len(tasks) == 0 {
								noTasksLabel.Show()
							} else {
								noTasksLabel.Hide()
							}
						}
						tasksList.UnselectAll()
						w.SetContent(content)
					},
					w,
				)
			}),
			// Edit
			widget.NewButtonWithIcon("", theme.DocumentCreateIcon(), func() {
				editBar := container.NewHBox(
					widget.NewLabel(
						fmt.Sprintf(
							"Editing \"%s\"",
							tasks[id].Title,
						),
					),
					layout.NewSpacer(),
					widget.NewButtonWithIcon("", theme.CancelIcon(), func() {
						w.SetContent(content)
						tasksList.Unselect(id)
					}),
				)
				editTitle := widget.NewEntry()
				editTitle.SetText(tasks[id].Title)
				editDescription := widget.NewMultiLineEntry()
				editDescription.SetText(tasks[id].Description)

				editButton := widget.NewButtonWithIcon("Save task", theme.DocumentSaveIcon(), func() {
					DB.Model(&Task{}).Where("id = ?", tasks[id].ID).Updates(Task{
						Title:       editTitle.Text,
						Description: editDescription.Text,
					})
					DB.Find(&tasks)
					w.SetContent(content)
					tasksList.UnselectAll()
				})

				editContent := container.NewVBox(
					editBar,
					canvas.NewLine(color.White),
					editTitle,
					editDescription,
					editButton,
				)

				w.SetContent(editContent)
			}),
		)

		detailsVBox := container.NewVBox(detailsBar, canvas.NewLine(color.White), taskTitle, taskDescription, buttonsBox)
		w.SetContent(detailsVBox)
	}

	tasksScroll := container.NewScroll(tasksList)
	tasksScroll.SetMinSize(fyne.NewSize(500, 500))

	content = container.NewVBox(
		tasksBar,
		canvas.NewLine(color.White),
		noTasksLabel,
		tasksScroll,
	)

	titleEntry := widget.NewEntry()
	titleEntry.SetPlaceHolder("Task title...")

	descriptionEntry := widget.NewMultiLineEntry()
	descriptionEntry.SetPlaceHolder("Task description...")
	descriptionEntry.Wrapping = fyne.TextWrapWord

	saveTaskButton := widget.NewButtonWithIcon("", theme.DocumentSaveIcon(), func() {
		task := Task{
			Title:       titleEntry.Text,
			Description: descriptionEntry.Text,
		}
		DB.Create(&task)
		DB.Find(&tasks)

		titleEntry.Text = ""
		titleEntry.Refresh()
		descriptionEntry.Text = ""
		descriptionEntry.Refresh()

		w.SetContent(content)
		tasksList.UnselectAll()

		if len(tasks) == 0 {
			noTasksLabel.Show()
		} else {
			noTasksLabel.Hide()
		}
	})

	createBar := container.NewHBox(
		widget.NewLabel("Create new task"),
		layout.NewSpacer(),
		widget.NewButtonWithIcon("", theme.CancelIcon(), func() {
			w.SetContent(content)
			tasksList.UnselectAll()
		}),
	)

	createContent = container.NewVBox(
		createBar,
		canvas.NewLine(color.White),
		container.NewVBox(
			titleEntry,
			descriptionEntry,
			saveTaskButton,
		),
	)

	w.SetContent(content)
	w.Show()
	a.Run()
}
