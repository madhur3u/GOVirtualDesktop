package main

// THIS IS MADE USING LINUX (UBUNTU 20.04) AND MIGHT NOT WORK WELL WITH WINDOWS

import (
	"encoding/json"
	"fmt"

	"image/color"
	"io/ioutil" // for reading http
	"net/http"  // this is to make http request to get API

	"fyne.io/fyne/v2"
	"fyne.io/fyne/v2/canvas"
	"fyne.io/fyne/v2/container"
	"fyne.io/fyne/v2/layout"
	"fyne.io/fyne/v2/widget"
)

func weatherApp() {

	w := a.NewWindow("Weather App")
	w.Resize(fyne.NewSize(400, 400))

	// api part of the code
	// openweathermap.org  --> API --> create API keys --> https://api.openweathermap.org/data/2.5/weather?q=dehradun&APPID=43a6cd4cffb4ddd6c32017f7a570fffb

	input := widget.NewEntry()
	input.SetPlaceHolder("Enter City Name")

	text := canvas.NewText("https://api.openweathermap.org/data/2.5/weather?q=kyoto&APPID=43a6cd4cffb4ddd6c32017f7a570fffb", color.White)
	text.TextSize = 20

	response, err := http.Get(text.Text)

	if err != nil { // if we encounter error
		fmt.Println("Error, try again")
	}

	// defer response.Body.Close() // close the http part after we have our response, now we will work on response

	// now we have to work with the html body we have, the api will give us a JSON file
	// we are reading that json file in body variable from response, and error
	body, err := ioutil.ReadAll(response.Body)

	if err != nil { // if we encounter error
		fmt.Println("Error, try again")
	}

	// using unmarshalweather structure we will have weather details in weather variable
	// this is the last part in getting data from API
	weather, err := UnmarshalWeather(body)

	if err != nil { // if we encounter error
		fmt.Println("Error, try again")
	}

	// UI part of the code
	img := canvas.NewImageFromFile("bg1.png")
	img.FillMode = canvas.ImageFillStretch

	label1 := canvas.NewText("Weather Information", color.White)
	label1.TextStyle = fyne.TextStyle{Bold: true} // styling label1

	label2 := canvas.NewText(fmt.Sprintf("\t\t\t\tCOUNTRY    \t\t\t:\t %s", weather.Sys.Country), color.White) // taking all these from structures
	label3 := canvas.NewText(fmt.Sprintf("\t\t\t\tWIND SPEED \t\t:\t %.2f mph", weather.Wind.Speed), color.White)
	label4 := canvas.NewText(fmt.Sprintf("\t\t\t\tTEMPRATURE \t\t:\t %.2f °C", weather.Main.Temp-273.15), color.White)
	label5 := canvas.NewText(fmt.Sprintf("\t\t\t\tCITY       \t\t\t\t\t:\t %s", weather.Name), color.White)
	label6 := canvas.NewText(fmt.Sprintf("\t\t\t\tFEELS LIKE \t\t\t:\t %.2f °C", weather.Main.FeelsLike-273.15), color.White)
	label7 := canvas.NewText(fmt.Sprintf("\t\t\t\tDESCRIPTION \t\t:\t %s", weather.Weather[0].Description), color.White)
	label8 := canvas.NewText(fmt.Sprintf("\t\t\t\tHUMIDITY \t\t\t\t:\t %d %%", weather.Main.Humidity), color.White)

	btn1 := widget.NewButton("SEARCH", func() {

		if input.Text == "" {
			input.Text = "Kyoto"
		}

		text.Text = "https://api.openweathermap.org/data/2.5/weather?q=" + input.Text + "&APPID=43a6cd4cffb4ddd6c32017f7a570fffb"
		text.Refresh()

		// fmt.Println(text.Text)

		response, err := http.Get(text.Text)

		if err != nil { // if we encounter error
			fmt.Println("Error, try again")
		}

		// now we have to work with the html body we have, the api will give us a JSON file
		// we are reading that json file in body variable from response, and error
		body, err := ioutil.ReadAll(response.Body)

		if err != nil { // if we encounter error
			fmt.Println("Error, try again")
		}

		// using unmarshalweather structure we will have weather details in weather variable
		// this is the last part in getting data from API
		weather, err := UnmarshalWeather(body)

		if err != nil { // if we encounter error
			fmt.Println("Error, try again")
		}

		label2.Text = fmt.Sprintf("\t\t\t\tCOUNTRY    \t\t\t:\t %s", weather.Sys.Country)
		label3.Text = fmt.Sprintf("\t\t\t\tWIND SPEED \t\t:\t %.2f mph", weather.Wind.Speed)
		label4.Text = fmt.Sprintf("\t\t\t\tTEMPRATURE \t\t:\t %.2f °C", weather.Main.Temp-273.15)
		label5.Text = fmt.Sprintf("\t\t\t\tCITY       \t\t\t\t\t:\t %s", weather.Name)
		label6.Text = fmt.Sprintf("\t\t\t\tFEELS LIKE \t\t\t:\t %.2f °C", weather.Main.FeelsLike-273.15)
		label7.Text = fmt.Sprintf("\t\t\t\tDESCRIPTION \t\t:\t %s", weather.Weather[0].Description)
		label8.Text = fmt.Sprintf("\t\t\t\tHUMIDITY \t\t\t\t:\t %d %%", weather.Main.Humidity)

		label2.Refresh()
		label3.Refresh()
		label4.Refresh()
		label5.Refresh()
		label6.Refresh()
		label7.Refresh()
		label8.Refresh()
		// fmt.Printf(weather.Sys.Country)

	})

	// fmt.Println(text.Text)

	w.SetContent(

		container.New(layout.NewMaxLayout(), img,
			container.NewVBox(
				container.New(layout.NewCenterLayout(), label1),
				input,
				btn1,
				container.NewHBox(layout.NewSpacer()),
				container.NewHBox(layout.NewSpacer()),
				container.NewHBox(layout.NewSpacer()),
				container.NewHBox(layout.NewSpacer()),
				container.NewHBox(layout.NewSpacer()),
				container.NewHBox(layout.NewSpacer()),
				container.NewHBox(layout.NewSpacer()),
				container.NewHBox(layout.NewSpacer()),
				container.NewHBox(layout.NewSpacer()),
				container.NewHBox(layout.NewSpacer()),
				container.NewHBox(layout.NewSpacer()),
				container.NewHBox(layout.NewSpacer()),
				container.NewHBox(layout.NewSpacer()),
				label2,
				label5,
				label4,
				label6,
				label3,
				label7,
				label8,
			),
		),
	)
	w.Show()
}

// in GO we cannot use JSON file directly we have to create a structure from that
// so we make structure using a website which will be used in our program
// quicktype.io --> to make JSON to structure

func UnmarshalWeather(data []byte) (weather, error) {
	var r weather
	err := json.Unmarshal(data, &r)
	return r, err
}

func (r *weather) Marshal() ([]byte, error) {
	return json.Marshal(r)
}

type weather struct {
	Coord      Coord     `json:"coord"`
	Base       string    `json:"base"`
	Weather    []Weather `json:"weather"`
	Main       Main      `json:"main"`
	Visibility int64     `json:"visibility"`
	Wind       Wind      `json:"wind"`
	Clouds     Clouds    `json:"clouds"`
	Dt         int64     `json:"dt"`
	Sys        Sys       `json:"sys"`
	Timezone   int64     `json:"timezone"`
	ID         int64     `json:"id"`
	Name       string    `json:"name"`
	Cod        int64     `json:"cod"`
}

type Clouds struct {
	All int64 `json:"all"`
}

type Coord struct {
	Lon float64 `json:"lon"`
	Lat float64 `json:"lat"`
}

type Main struct {
	Temp      float64 `json:"temp"`
	FeelsLike float64 `json:"feels_like"`
	TempMin   float64 `json:"temp_min"`
	TempMax   float64 `json:"temp_max"`
	Pressure  int64   `json:"pressure"`
	Humidity  int64   `json:"humidity"`
	SeaLevel  int64   `json:"sea_level"`
	GrndLevel int64   `json:"grnd_level"`
}

type Sys struct {
	Country string `json:"country"`
	Sunrise int64  `json:"sunrise"`
	Sunset  int64  `json:"sunset"`
}

type Weather struct {
	ID          int64  `json:"id"`
	Main        string `json:"main"`
	Description string `json:"description"`
	Icon        string `json:"icon"`
}

type Wind struct {
	Speed float64 `json:"speed"`
	Deg   int64   `json:"deg"`
	Gust  float64 `json:"gust"`
}
