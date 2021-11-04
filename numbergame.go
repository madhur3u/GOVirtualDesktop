package main

// THIS IS MADE USING LINUX (UBUNTU 20.04) AND MIGHT NOT WORK WELL WITH WINDOWS

import (
	"image/color"
	"strconv"

	"fyne.io/fyne/v2"
	"fyne.io/fyne/v2/canvas"
	"fyne.io/fyne/v2/container"
	"fyne.io/fyne/v2/layout"
	"fyne.io/fyne/v2/widget"

	"fmt"
	"math/rand"
	"time"
)

func GuessNumber() {

	min, max := 1, 100
	rand.Seed(time.Now().UnixNano())
	secretNumber := rand.Intn(max-min) + min
	fmt.Println("The secret number is", secretNumber)

	bg := canvas.NewImageFromFile("game.jpg")
	bg.FillMode = canvas.ImageFillStretch

	w := a.NewWindow("Guess The Number")
	w.Resize(fyne.NewSize(500, 300))

	codes := widget.NewLabel(
		"Rules\nYou have to choose a number between 1 and 100\nAccording to the instructions below reach the\ngiven number in minimum possible attempts\n\nHow to read instructions :\n'd' is difference between your number and answer\nYou are so far right now\t\t\t\t: d > 50\nYou are quite near.\t\t\t\t\t\t: 50 > d > 25\nYou are so close. Keep going.\t: 25 > d > 10\nJust there.\t\t\t\t\t\t\t\t\t\t: 10 > d > 5\nOnly few steps away.\t\t\t\t\t: d < 5\n",
	)

	rules := widget.NewButton("Rules", func() {
		w2 := a.NewWindow("Rules")
		w2.Resize(fyne.NewSize(200, 200))
		w2.SetContent(codes)
		w2.Show()
	})

	top := canvas.NewText("Enter any number between 1 - 100", color.White)
	top.Alignment = fyne.TextAlignCenter
	top.TextStyle = fyne.TextStyle{Bold: true}
	top.TextSize = 20

	attempt := 0
	att := canvas.NewText("   No. of Attempts : "+strconv.Itoa(attempt), color.White)
	att.TextStyle = fyne.TextStyle{Italic: true}
	input := widget.NewEntry()
	input.SetPlaceHolder("Enter your guess here")
	result := canvas.NewText("", color.White)
	result.TextSize = 15

	btn1 := widget.NewButton("Check", func() {

		input.SetPlaceHolder(input.Text)

		if attempt == 0 {
			result.TextStyle = fyne.TextStyle{Bold: false}
			result.TextSize = 15
		}

		attempt++
		ans, _ := strconv.Atoi(input.Text)
		att.Text = "   No. of Attempts : " + strconv.Itoa(attempt)

		diff := Abs(secretNumber - ans)

		if ans == secretNumber {
			result.Text = "  You did it. The number was " + input.Text + "."
			result.TextStyle = fyne.TextStyle{Bold: true}
			result.TextSize = 18
			attempt = 0
			input.Text = ""
			input.Refresh()
		} else if diff > 50 {
			result.Text = "  Try again. You are so far right now."
		} else if diff < 50 && diff > 25 {
			result.Text = "  Try again. You are quite near."
		} else if diff < 25 && diff > 10 {
			result.Text = "  Try again. You are so close. Keep going."
		} else if diff < 10 && diff > 5 {
			result.Text = "  Try again. Just there."
		} else if diff < 5 {
			result.Text = "  Try again. Only few steps away."
		}

		input.Text = ""
		att.Refresh()
		input.Refresh()
		result.Refresh()
	})
	btn2 := widget.NewButton("New Game", func() {

		secretNumber = rand.Intn(max-min) + min
		fmt.Println("The secret number is", secretNumber)

		result.Text = ""
		result.TextStyle = fyne.TextStyle{Bold: false}
		result.TextSize = 15
		attempt = 0
		att.Text = "   No. of Attempts : " + strconv.Itoa(attempt)
		att.Refresh()
		result.Refresh()
		input.Text = ""
		input.SetPlaceHolder("Enter your guess here")
		input.Refresh()

	})

	btn3 := widget.NewButton("Give Up", func() {

		ans1 := strconv.Itoa(secretNumber)
		result.Text = "  Failed. The number was " + ans1 + "."
		result.TextStyle = fyne.TextStyle{Bold: true}
		result.TextSize = 18
		att.Text = "   No. of Attempts : " + strconv.Itoa(attempt)
		attempt = 0
		att.Refresh()
		result.Refresh()
		input.Text = ""
		input.SetPlaceHolder("Enter your guess here")
		input.Refresh()
	})

	content := container.NewVBox(rules, top, input, container.NewHBox(btn1, btn2, btn3), result, att)

	w.SetContent(container.New(layout.NewMaxLayout(), bg, content))
	w.Show()
}

func Abs(x int) int {
	if x < 0 {
		return -x
	}
	return x
}
