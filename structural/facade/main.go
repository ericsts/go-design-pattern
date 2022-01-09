package facade

import (
	"encoding/json"
	"fmt"
	"io"
	"io/ioutil"
	"net/http"
)

// CurrentWeatherDataRetriever .,..
type CurrentWeatherDataRetriever interface {
	GetByCityAndCountryCode(city, countryCode string) (Weather, error)
	GetByGeoCoordinates(lat, lon float32) (Weather, error)
}

// CurrentWeatherData ...
type CurrentWeatherData struct {
	APIkey string
}

// Weather ...
type Weather struct {
	ID      int
	Coord   string
	Weather string
	Base    string
	Main    string
	Wind    string
	Clouds  string
	Rain    string
	Dt      string
	Sys     string
}

func (c *CurrentWeatherData) responseParser(body io.Reader) (*Weather, error) {
	// return nil, fmt.Errorf("Not implemented yet")

	w := new(Weather)
	err := json.NewDecoder(body).Decode(w)
	if err != nil {
		return nil, err
	}
	return w, nil
}

// GetByGeoCoordinates ...
func (c *CurrentWeatherData) GetByGeoCoordinates(lat, lon float32) (weather *Weather, err error) {
	return c.doRequest(fmt.Sprintf("http://api.openweathermap.org/data/2.5/weather?q=%f,%f&APPID=%s", lat, lon, c.APIkey))
}

// GetByCityAndCountryCode ...
func (c *CurrentWeatherData) GetByCityAndCountryCode(city, countryCode string) (weather *Weather, err error) {
	return c.doRequest(fmt.Sprintf("http://api.openweathermap.org/data/2.5/weather?lat=%s&lon=%s&APPID=%s", city, countryCode, c.APIkey))
}

func (c *CurrentWeatherData) doRequest(uri string) (weather *Weather, err error) {
	client := &http.Client{}
	req, err := http.NewRequest("GET", uri, nil)
	if err != nil {
		return
	}
	req.Header.Set("Content-Type", "application/json")
	resp, err := client.Do(req)
	if err != nil {
		return
	}
	if resp.StatusCode != 200 {
		byt, errMsg := ioutil.ReadAll(resp.Body)
		if errMsg == nil {
			errMsg = fmt.Errorf("%s", string(byt))
		}
		err = fmt.Errorf("status code was %d, aborting. error message was:%s", resp.StatusCode, errMsg)
		return
	}
	weather, err = c.responseParser(resp.Body)
	resp.Body.Close()
	return
}

/*
to use it
weatherMap := CurrentWeatherData{*apiKey}
weather, err := weatherMap.GetByCityAndCountryCode("Madrid", "ES")
if err != nil {
t.Fatal(err)
}
fmt.Printf("Temperature in Madrid is %f celsius\n",
weather.Main.Temp-273.15)
*/
