package main

// THIS IS MADE USING LINUX (UBUNTU 20.04) AND MIGHT NOT WORK WELL WITH WINDOWS

import (
	// "fmt"
	"image/color"
	"math/rand"
	"strconv"
	"time"

	"fyne.io/fyne/v2"
	"fyne.io/fyne/v2/canvas"
	"fyne.io/fyne/v2/container"
	"fyne.io/fyne/v2/layout"
	"fyne.io/fyne/v2/widget"
)

func passwordApp() {

	w := a.NewWindow("Password Generator")
	w.Resize(fyne.NewSize(400, 0))

	blue_color := canvas.NewRectangle(color.NRGBA{R: 0, G: 96, B: 215, A: 255})

	title := canvas.NewText("Welcome to Password Generator", color.White)
	title.TextStyle = fyne.TextStyle{Bold: true}

	body := widget.NewLabel("Enter the length of the password you need \nin the input box below and the app will provide \nyou the best possible random password which will \nbe highly secure and unique.")
	inst1 := widget.NewLabel("Copy the password after it is generated")
	input := widget.NewEntry()
	input.SetPlaceHolder("Enter password length")

	// text := canvas.NewText("", color.White)
	// text.TextSize = 20

	text := widget.NewEntry()
	text.SetPlaceHolder("Password")

	// button to generate password
	btn1 := widget.NewButton("Generate", func() {
		// input
		passlength, _ := strconv.Atoi(input.Text) // convert string to integer
		text.SetText(PasswordGenerator(passlength))
		text.Refresh()
	})

	// show content
	w.SetContent(

		container.NewVBox(

			container.New(layout.NewCenterLayout(), title), // center align
			body,
			inst1,
			input,
			text,
			container.New(layout.NewMaxLayout(), blue_color, btn1),
		),
	)
	w.Show()
}

// function which generate and return password
func PasswordGenerator(passwordLength int) string {
	// Lower case
	lowCase := "abcdefghijklmnopqrstuvxyz"
	// Upper Case
	upCase := "ABCDEFGHIJKLMNOPQRSTUVXYZ"
	// Numbers
	Numbers := "0123456789"
	// Special characters
	SpecialChar := "!@#$&*()_?"

	// variable for storing password
	password := ""
	// loop
	for n := 0; n < passwordLength; n++ {
		// Now random characters
		rand.Seed(time.Now().UnixNano())
		randNum := rand.Intn(4)
		// fmt.Println(randNum)
		// Switch statment
		switch randNum {
		case 0:
			rand.Seed(time.Now().UnixNano())
			randNum := rand.Intn(len(lowCase))
			// len to find lenth of slice/array
			// NOw we will store the generated passowrd character
			password = password + string(lowCase[randNum])

		case 1:
			rand.Seed(time.Now().UnixNano())
			randNum := rand.Intn(len(upCase))
			// len to find lenth of slice/array
			// NOw we will store the generated passowrd character
			password = password + string(upCase[randNum])

		case 2:
			rand.Seed(time.Now().UnixNano())
			randNum := rand.Intn(len(Numbers))
			// len to find lenth of slice/array
			// NOw we will store the generated passowrd character
			password = password + string(Numbers[randNum])

		case 3:
			rand.Seed(time.Now().UnixNano())
			randNum := rand.Intn(len(SpecialChar))
			// len to find lenth of slice/array
			// NOw we will store the generated passowrd character
			password = password + string(SpecialChar[randNum])

		} // end of switch
	} // end of for loop
	// fmt.Println(password)

	return password
}
