package posts

import (
	"encoding/json"
	"net/http"
	"time"
)

var myClient = &http.Client{Timeout: 10 * time.Second}

func getJSON(url string, target interface{}) error {
	r, err := myClient.Get(url)
	if err != nil {
		return err
	}
	defer r.Body.Close()
	return json.NewDecoder(r.Body).Decode(target)
}

// WeatherReport represents a typical weather report
type WeatherReport struct {
	ID              int32   `json:"id,omitempty"`
	TemperatureLow  float32 `json:"temperatureLow,omitempty"`
	TemperatureHigh float32 `json:"temperatureHigh,omitempty"`
	Precipitation   float32 `json:"precipitation,omitempty"`
	Humidity        float32 `json:"humidity,omitempty"`
	Wind            float32 `json:"wind,omitempty"`
}

func loadPost2() {
	// `d`, the reference to the DOM is declared in a previous post
	d.GetElementByID("post-2-title").SetInnerHTML("Day 2: Connecting a REST API")
	d.GetElementByID("post-2-attribution").SetInnerHTML("By Joseph Choi under <a class='post-category post-category-js'>Go</a> <a class='post-category post-category-js'>REST API</a> <a class='post-category post-category-js'>Heroku</a>")
	d.GetElementByID("post-2-description-1").SetInnerHTML("In an attempt to get a fullstack Golang project up and running, I decided to deploy a very simple Golang REST API to Heroku, and connect it to the GopherJS frontend. Below is the relevant server-side and frontend code.")

	url := "https://gopherjs-api.herokuapp.com/weatherreport"
	weatherReport := new(WeatherReport) // or &Foo{}
	getJSON(url, weatherReport)
	println(weatherReport.TemperatureLow)
}
