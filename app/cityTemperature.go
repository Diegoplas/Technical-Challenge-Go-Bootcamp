package main

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
	"os"
	"time"
)

const (
	API_URL string = "api.openweathermap.org"
)

func GetCurrentCityWeather(city, unit string) (*CurrentWeatherResponse, error) {

	var weatherResponse CurrentWeatherResponse
	var invalidCityError APIError
	var errUnmarshal error
	errUnmarshal = nil

	openWeatherMapKey := OpenWeatherMap{API_KEY: os.Getenv("OWM_API_KEY")}
	if openWeatherMapKey.API_KEY == "" {
		// No API keys present, return error
		return nil, fmt.Errorf("credentials file not found, please run cerebro-cli creds")
	}

	getWeatherURL := fmt.Sprintf("http://%s/data/2.5/weather?q=%s&units=%s&APPID=%s", API_URL, city, unit, openWeatherMapKey.API_KEY)

	body, err := makeExternalApiRequest(getWeatherURL)
	if err != nil {
		return nil, fmt.Errorf("error making the external api request %v", err.Error())
	}

	if len(body) == 40 {
		errUnmarshal = json.Unmarshal(body, &invalidCityError)
	} else {
		errUnmarshal = json.Unmarshal(body, &weatherResponse)
	}

	if errUnmarshal != nil {
		return nil, fmt.Errorf("err = %s", errUnmarshal)
	}
	return &weatherResponse, nil
}

func makeExternalApiRequest(url string) ([]byte, error) {
	// Build an http client so we can have control over timeout
	client := &http.Client{
		Timeout: time.Second * 10,
	}
	response, err := client.Get(url)
	if err != nil {
		return nil, fmt.Errorf("error making Get request to the client: %v", err.Error())
	}

	// defer the closing of the res body
	defer response.Body.Close()
	// read the http response body into a byte stream
	body, readErr := ioutil.ReadAll(response.Body)
	if readErr != nil {
		return nil, readErr
	}

	return body, nil
}
