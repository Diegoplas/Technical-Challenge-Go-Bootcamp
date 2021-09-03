package main

import (
	"fmt"
	"log"
	"net/http"

	"github.com/gorilla/handlers"
	"github.com/gorilla/mux"
)

const (
	Farenheit = "imperial"
	Celcius   = "metric"
	Kelvin    = "internal"
)

func main() {

	router := getRouter()
	methods := handlers.AllowedMethods([]string{"GET"})
	log.Fatal(http.ListenAndServe(":8000", handlers.CORS(methods)(router)))
}

func getRouter() (router *mux.Router) {
	router = mux.NewRouter()
	router.HandleFunc("/", helloWorldHandler).Methods("GET")
	router.HandleFunc("/get-weather", getWeatherFromCityHandler).Methods("GET")
	return router
}

func getWeatherFromCityHandler(w http.ResponseWriter, r *http.Request) {

	formatedResponse := ""

	//Get the param from the URL
	params := r.URL.Query()
	paramCity := params.Get("city")

	if paramCity == "" {
		http.Error(w, "missing city query param", http.StatusBadRequest)
		return
	}

	cityWeather, _ := GetCurrentCityWeather(paramCity, Celcius)

	//No city name will be obtained in case of an APIError
	if cityWeather.Name == "" {
		formatedResponse = fmt.Sprintf("The city (%s) in not valid. Please introduce a valid one.", paramCity)
	} else {
		formatedResponse = fmt.Sprintf("City: %s \n Weather: %s \n Temperature from %v to %v °С \n Geo Coords [Lon: %v Alt: %v].",
			cityWeather.Name, cityWeather.Weather[0].Description, cityWeather.Main.TempMin, cityWeather.Main.TempMax, cityWeather.Coord.Lon, cityWeather.Coord.Lat)
	}

	w.Write([]byte(formatedResponse))
}

func helloWorldHandler(w http.ResponseWriter, r *http.Request) {

	defaultMessage := ""
	//Get the param from the URL
	param := r.URL.Query()

	helloName := param.Get("name")

	w.Write([]byte(defaultMessage))
	if helloName == "" {
		defaultMessage = "Hello, World!"
	} else {
		defaultMessage = fmt.Sprintf("Hello, %s!", helloName)
	}

	w.Write([]byte(defaultMessage))
}
