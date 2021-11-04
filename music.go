package main

// THIS IS MADE USING LINUX (UBUNTU 20.04) AND MIGHT NOT WORK WELL WITH WINDOWS

import (
	// "fmt"
	"image/color"
	"time"

	"fyne.io/fyne/v2"
	"fyne.io/fyne/v2/canvas"
	"fyne.io/fyne/v2/container"
	"fyne.io/fyne/v2/dialog"
	"fyne.io/fyne/v2/layout"
	"fyne.io/fyne/v2/storage"
	"fyne.io/fyne/v2/theme"
	"fyne.io/fyne/v2/widget"
	"github.com/faiface/beep"
	"github.com/faiface/beep/mp3"
	"github.com/faiface/beep/speaker"
)

func musicPlayer() {

	// var f *os.File
	var format beep.Format
	var streamer beep.StreamSeekCloser
	var pause bool = false

	w := a.NewWindow("MP3 Player")
	w.Resize(fyne.NewSize(600, 400))

	logo := canvas.NewImageFromFile("music.jpg")
	logo.FillMode = canvas.ImageFillStretch

	icon1 := canvas.NewImageFromFile("ico.png")
	icon1.FillMode = canvas.ImageFillContain
	icon1.SetMinSize(fyne.NewSize(100, 100))

	toolbar := widget.NewToolbar(

		widget.NewToolbarSpacer(),
		widget.NewToolbarAction(theme.MediaPlayIcon(), func() {

			speaker.Init(format.SampleRate, format.SampleRate.N(time.Second/10))
			speaker.Play(streamer)
		}),
		widget.NewToolbarAction(theme.MediaPauseIcon(), func() {
			if !pause {
				pause = true
				speaker.Lock()
			} else if pause {
				pause = false
				speaker.Unlock()
				speaker.Play()
			}
		}),
		widget.NewToolbarSpacer(),
	)
	label1 := canvas.NewText("MP3 Audio Player", color.White)
	label1.Alignment = fyne.TextAlignCenter
	label1.TextSize = 20
	label1.TextStyle = fyne.TextStyle{Bold: true}

	text1 := canvas.NewText("Browse MP3 files using the button below and play", color.White)
	text1.TextSize = 12
	text1.TextStyle = fyne.TextStyle{Italic: true}
	text1.Alignment = fyne.TextAlignCenter

	label2 := widget.NewLabel("Audio File Name")
	label2.Alignment = fyne.TextAlignCenter

	browse_files := widget.NewButton("Browse MP3", func() {

		fd := dialog.NewFileOpen(func(uc fyne.URIReadCloser, _ error) {
			streamer, format, _ = mp3.Decode(uc)
			label2.Text = uc.URI().Name()
			label2.Refresh()
		}, w)

		fd.Show()
		fd.SetFilter(storage.NewExtensionFileFilter([]string{".mp3"}))

	})

	blank := canvas.NewText("", color.White)

	// slider := widget.NewSlider(0, 100)
	c := container.NewVBox(blank, blank, blank, icon1, label1, text1, container.NewHBox(layout.NewSpacer(), browse_files, layout.NewSpacer()), label2, toolbar)
	w.SetContent(container.New(layout.NewMaxLayout(), logo, container.NewVBox(c)))

	w.Show()
}
