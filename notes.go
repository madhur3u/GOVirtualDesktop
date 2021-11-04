package main

// THIS IS MADE USING LINUX (UBUNTU 20.04) AND MIGHT NOT WORK WELL WITH WINDOWS

import (
	"image/color"
	"io/ioutil"
	"strconv"

	"fyne.io/fyne/v2"
	"fyne.io/fyne/v2/canvas"
	"fyne.io/fyne/v2/container"
	"fyne.io/fyne/v2/dialog"
	"fyne.io/fyne/v2/layout"
	"fyne.io/fyne/v2/storage"
	"fyne.io/fyne/v2/widget"
)

func notesApp() {

	w := a.NewWindow("NOTEPAD")
	w.Resize(fyne.NewSize(600, 400))

	blue_color := canvas.NewRectangle(color.NRGBA{R: 0, G: 96, B: 215, A: 255})

	count := 1

	head := widget.NewEntry()
	head.SetPlaceHolder("Enter file name (otherwise default name will be taken)")

	input := widget.NewMultiLineEntry()
	input.SetPlaceHolder("Enter your text here")

	list := container.NewVBox()

	save := widget.NewButton("Save File", func() {

		saveFileDialog := dialog.NewFileSave(
			func(uc fyne.URIWriteCloser, _ error) {

				textData := []byte(input.Text)
				uc.Write(textData)
			}, w)

		saveFileDialog.SetFileName(head.Text + ".txt")
		saveFileDialog.Show()

	})

	open := widget.NewButton("Open File", func() {

		openFileDialog := dialog.NewFileOpen(
			func(rf fyne.URIReadCloser, _ error) {

				readData, _ := ioutil.ReadAll(rf)
				output := fyne.NewStaticResource("Open File", readData)

				viewData := widget.NewMultiLineEntry()
				viewData.SetText(string(output.StaticContent))

				input.SetText(viewData.Text)
				head.SetText(output.Name())
			}, w)

		openFileDialog.SetFilter(storage.NewExtensionFileFilter([]string{".txt"}))
		openFileDialog.Show()

	})

	list.Add(container.New(layout.NewMaxLayout(), blue_color, save))
	list.Add(open)

	right := container.New(layout.NewBorderLayout(head, nil, nil, nil), head, input)

	list.Add(widget.NewButton("New File", func() {

		head.SetText("New File " + strconv.Itoa(count))
		count++
	}))

	split := container.NewHSplit(list, right)
	split.Offset = 0.20

	w.SetContent(split)

	w.Show()
}
