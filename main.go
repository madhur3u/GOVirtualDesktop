package main

// use the following line to execute
// go build main.go calc.go notes.go gallery.go password.go weather.go music.go numbergame.go getTime.go
// go run main.go calc.go notes.go gallery.go password.go weather.go music.go numbergame.go getTime.go
// THIS IS MADE USING LINUX (UBUNTU 20.04) AND MIGHT NOT WORK WELL WITH WINDOWS
// PLEASE CONNECT TO WIFI / INTERNET BEFORE EXECUTION

// all packages needed by main.go file are imported here
import (
	"image/color"
	"io/ioutil"
	"net/http"

	"fyne.io/fyne/v2"
	"fyne.io/fyne/v2/app"
	"fyne.io/fyne/v2/canvas"
	"fyne.io/fyne/v2/container"
	"fyne.io/fyne/v2/theme"

	"fyne.io/fyne/v2/layout"
	"fyne.io/fyne/v2/widget"
)

// initialize new app and window
var a fyne.App = app.New()
var w fyne.Window = a.NewWindow("OS")

func main() {

	w.SetFullScreen(true) // open the current window in fullscreen

	// setting home screen wallapaper
	img1, _ := http.Get("https://user-images.githubusercontent.com/89251393/140249554-a867f149-5e27-4c53-b56e-419c029864c5.jpg")
	body, _ := ioutil.ReadAll(img1.Body)
	img := canvas.NewImageFromResource(fyne.NewStaticResource("", body))
	img.FillMode = canvas.ImageFillStretch

	//-------------------------------------------DATE TIME-----------------------
	// this is used to set DATE and TIME in home screen and change its style and font
	// function getTime is made in file --> getTime.go
	datetime := canvas.NewText(getTime(), color.White)
	datetime.TextSize = 15
	datetime.Alignment = fyne.TextAlignCenter
	datetime.TextStyle = fyne.TextStyle{Bold: true}

	// since we need to update time every second
	// so we need to call that function in background and update our datetime Text
	// datetime.Refresh() to show updated value in app
	go func() {
		for {
			datetime.Text = getTime()
			datetime.Refresh()
		}
	}()

	//-----------------------------------------APPS STARTS HERE-----------------------------------
	// from here we are setting up our apps and icons
	// icon variables hold the icon of the particular app
	// and then a button is made which call a function having code of that app
	img2, _ := http.Get("https://user-images.githubusercontent.com/89251393/140249548-0fc343c7-0da6-4586-a39e-8fbb13f5dd44.png")
	body1, _ := ioutil.ReadAll(img2.Body)
	icon1 := canvas.NewImageFromResource(fyne.NewStaticResource("", body1))
	icon1.FillMode = canvas.ImageFillContain
	calc := widget.NewButton("", func() {
		calculator()
	})

	img3, _ := http.Get("https://user-images.githubusercontent.com/89251393/140249614-adc807e1-0d93-417f-bb13-b8903c12ee9d.png")
	body3, _ := ioutil.ReadAll(img3.Body)
	icon2 := canvas.NewImageFromResource(fyne.NewStaticResource("", body3))
	icon2.FillMode = canvas.ImageFillContain
	notepad := widget.NewButton("", func() {
		notesApp()
	})

	img4, _ := http.Get("https://user-images.githubusercontent.com/89251393/140249556-e51b6917-2805-4086-8d31-e70eeac4c708.png")
	body4, _ := ioutil.ReadAll(img4.Body)
	icon3 := canvas.NewImageFromResource(fyne.NewStaticResource("", body4))
	icon3.FillMode = canvas.ImageFillContain
	gallery := widget.NewButton("", func() {
		galleryApp()
	})

	img5, _ := http.Get("https://user-images.githubusercontent.com/89251393/140250230-cdddb4bd-8403-4b1a-a586-78bc3b9ad24d.png")
	body5, _ := ioutil.ReadAll(img5.Body)
	icon4 := canvas.NewImageFromResource(fyne.NewStaticResource("", body5))
	icon4.FillMode = canvas.ImageFillContain
	weather := widget.NewButton("", func() {
		weatherApp()
	})

	img6, _ := http.Get("https://user-images.githubusercontent.com/89251393/140249591-17c40d54-869c-4d74-a89b-6c57b8e41053.png")
	body6, _ := ioutil.ReadAll(img6.Body)
	icon5 := canvas.NewImageFromResource(fyne.NewStaticResource("", body6))
	icon5.FillMode = canvas.ImageFillContain
	password := widget.NewButton("", func() {
		passwordApp()
	})

	img7, _ := http.Get("https://user-images.githubusercontent.com/89251393/140249608-9c1f5282-3172-42ca-bb8c-c5f7d7a45730.png")
	body7, _ := ioutil.ReadAll(img7.Body)
	icon7 := canvas.NewImageFromResource(fyne.NewStaticResource("", body7))
	icon7.FillMode = canvas.ImageFillContain
	music := widget.NewButton("", func() {
		musicPlayer()
	})

	img8, _ := http.Get("https://user-images.githubusercontent.com/89251393/140249587-1d0981cf-88d3-457d-93d2-e488213f8fdf.png")
	body8, _ := ioutil.ReadAll(img8.Body)
	icon8 := canvas.NewImageFromResource(fyne.NewStaticResource("", body8))
	icon8.FillMode = canvas.ImageFillContain
	numberGame := widget.NewButton("", func() {
		GuessNumber()
	})

	// a quit button is added to close the app
	img9, _ := http.Get("https://user-images.githubusercontent.com/89251393/140249619-cb66d5e9-0f2e-439c-a222-431cba8e5cac.png")
	body9, _ := ioutil.ReadAll(img9.Body)
	icon6 := canvas.NewImageFromResource(fyne.NewStaticResource("", body9))
	icon6.FillMode = canvas.ImageFillContain
	quit := widget.NewButtonWithIcon("", theme.CancelIcon(), func() {
		w.Close()
	})
	//-----------------------------------------APPS ENDS HERE-----------------------------------

	// this has all apps and their icons
	apps := container.New(layout.NewHBoxLayout(),

		layout.NewSpacer(),
		container.NewGridWrap(fyne.NewSize(50, 50), container.NewPadded(calc, icon1)),
		container.NewGridWrap(fyne.NewSize(50, 50), container.NewPadded(notepad, icon2)),
		container.NewGridWrap(fyne.NewSize(50, 50), container.NewPadded(gallery, icon3)),
		container.NewGridWrap(fyne.NewSize(50, 50), container.NewPadded(weather, icon4)),
		container.NewGridWrap(fyne.NewSize(50, 50), container.NewPadded(password, icon5)),
		container.NewGridWrap(fyne.NewSize(50, 50), container.NewPadded(music, icon7)),
		container.NewGridWrap(fyne.NewSize(50, 50), container.NewPadded(numberGame, icon8)),
		container.NewGridWrap(fyne.NewSize(50, 50), container.NewPadded(quit, icon6)),
		layout.NewSpacer(),
	)

	// making desktop using set content
	w.SetContent(
		container.New(layout.NewMaxLayout(), img,
			container.New(layout.NewBorderLayout(datetime, apps, nil, nil), datetime, apps),
		),
	)

	w.ShowAndRun()
}
