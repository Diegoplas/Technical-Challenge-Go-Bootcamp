package main

type OpenWeatherMap struct {
	API_KEY string
}

type APIError struct {
	Message string `json:"message"`
	COD     string `json:"cod"`
}

type Main struct {
	Temp    float64 `json:"temp"`
	TempMin float64 `json:"temp_min"`
	TempMax float64 `json:"temp_max"`
}
type Coord struct {
	Lon float64 `json:"lon"`
	Lat float64 `json:"lat"`
}

type Weather struct {
	Description string `json:"description"`
}

type CurrentWeatherResponse struct {
	Coord   Coord     `json:"coord"`
	Main    Main      `json:"main"`
	Weather []Weather `json:"weather"`
	Name    string    `json:"name"`
}
