package main

// THIS IS MADE USING LINUX (UBUNTU 20.04) AND MIGHT NOT WORK WELL WITH WINDOWS

import (
	"image/color"
	"strconv"

	"fyne.io/fyne/v2"
	"fyne.io/fyne/v2/canvas"
	"fyne.io/fyne/v2/container"
	"fyne.io/fyne/v2/layout"
	"fyne.io/fyne/v2/theme"
	"fyne.io/fyne/v2/widget"
	"github.com/Knetic/govaluate"
)

func calculator() {

	w := a.NewWindow("Claculator") // the window name will have calculator
	// a.Settings().SetTheme(theme.LightTheme())

	w.Resize(fyne.NewSize(350, 0)) // setting size of window

	// output is what will be shown in the screen, this will change after each button pressed

	// help := widget.NewLabel("Input and answer will be shown below\nYou can input using mouse or keyboard\nClick below to input using keyboard")

	output := ""
	// screen := widget.NewLabel(output)
	screen := widget.NewEntry()
	screen.SetPlaceHolder("Input your values")

	// take care of history, make a string and display that in screen
	historyString := ""
	historyOutput := widget.NewLabel(historyString)
	isHistory := false

	var historyArray []string

	// these are all buttons which we need to make in our calculator with their functions
	// history button will show us history of all calculations done before
	history := widget.NewButton("History", func() {

		if isHistory {
			historyString = ""

		} else {
			for i := len(historyArray) - 1; i >= 0; i-- {
				historyString += historyArray[i] + "\n"
			}
		}
		historyOutput.SetText(historyString)
		isHistory = !isHistory
	})
	clear := widget.NewButton("C", func() {
		output = ""
		screen.SetText(output)
		historyString = ""
		historyOutput.SetText(historyString)
		screen.SetPlaceHolder("Input your values")

	})
	open := widget.NewButton("(", func() {
		output += "("
		screen.SetText(output)
	})
	close := widget.NewButton(")", func() {
		output += ")"
		screen.SetText(output)
	})
	modulo := widget.NewButton("%", func() {
		output += "%"
		screen.SetText(output)
	})
	// back button will remove last element from string if its length is > 0
	back := widget.NewButton("Back", func() {
		if len(output) > 0 {
			output = output[:len(output)-1]
			screen.SetText(output)
		}
		screen.SetPlaceHolder("Input your values")

	})
	divide := widget.NewButton("/", func() {
		output += "/"
		screen.SetText(output)
	})
	seven := widget.NewButton("7", func() {
		output += "7"
		screen.SetText(output)
	})
	eight := widget.NewButton("8", func() {
		output += "8"
		screen.SetText(output)
	})
	nine := widget.NewButton("9", func() {
		output += "9"
		screen.SetText(output)
	})
	multiply := widget.NewButton("x", func() {
		output += "*"
		screen.SetText(output)
	})
	four := widget.NewButton("4", func() {
		output += "4"
		screen.SetText(output)
	})
	five := widget.NewButton("5", func() {
		output += "5"
		screen.SetText(output)
	})
	six := widget.NewButton("6", func() {
		output += "6"
		screen.SetText(output)
	})
	minus := widget.NewButton("-", func() {
		output += "-"
		screen.SetText(output)
	})
	one := widget.NewButton("1", func() {
		output += "1"
		screen.SetText(output)
	})
	two := widget.NewButton("2", func() {
		output += "2"
		screen.SetText(output)
	})
	three := widget.NewButton("3", func() {
		output += "3"
		screen.SetText(output)
	})
	plus := widget.NewButton("+", func() {
		output += "+"
		screen.SetText(output)
	})
	zero2 := widget.NewButton("00", func() {
		output += "00"
		screen.SetText(output)
	})
	zero := widget.NewButton("0", func() {
		output += "0"
		screen.SetText(output)
		theme.BackgroundColor()
	})
	decimal := widget.NewButton(".", func() {
		output += "."
		screen.SetText(output)
	})

	// govaluate library is used to perform calculations
	// it takes a string with mathematical expressions and give desired answer
	equal := widget.NewButton("=", func() {

		expression, err := govaluate.NewEvaluableExpression(screen.Text)
		if err == nil {
			result, err := expression.Evaluate(nil)
			if err == nil {
				ans := strconv.FormatFloat(result.(float64), 'f', -1, 64)
				current_history := screen.Text + " = " + ans
				historyArray = append(historyArray, current_history)
				output = ans
			} else {
				output = ""
				screen.SetPlaceHolder("Input Invalid")
			}
		} else {
			output = ""
			screen.SetPlaceHolder("Input Invalid")
		}

		screen.SetText(output)
	})

	// color modules which is used to color buttons
	blue_color := canvas.NewRectangle(color.NRGBA{R: 0, G: 96, B: 215, A: 255})
	black_color := canvas.NewRectangle(color.NRGBA{R: 100, G: 100, B: 90, A: 255})

	// this will show us the buttons in calculator
	w.SetContent(container.NewVBox(

		screen,
		historyOutput,
		container.NewGridWithColumns(2,
			container.NewGridWithColumns(1, history),
			container.NewGridWithColumns(2,
				container.New(layout.NewMaxLayout(), black_color, clear),
				container.New(layout.NewMaxLayout(), black_color, back),
			),
		),
		container.NewGridWithColumns(4,
			open, close, modulo, divide,
			seven, eight, nine, multiply,
			four, five, six, minus,
			one, two, three, plus,
			zero2, zero, decimal,
			container.New(
				layout.NewMaxLayout(),
				blue_color,
				equal,
			),
		),
	))
	w.Show()
}
