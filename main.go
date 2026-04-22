package main

import(
	"encoding/json"
	"io/util"
	"net/http"
	"strings"
)

type apiConfigData struct{
	OpenWeatherMapApiKey string `json:"OpenWeatherMapApiKey"`
}
type weatherData struct{
	Name string `json:"name"`
	Main struct{
		Kelvin float64 `json:"temp"`
	}`json:"main"`
}


func loadApiConfig(filename string)(apiConfigData, error){
	bytes, err := io.util.ReadFile(filename)
	if err != nil{
		return apiConfigData{},err
	}

	var c apiConfigData
	err = json.Unmarshal(bytes, &c)
	if err != nil{
		return apiConfigData{},err
	}

	return c, nil
}



