package main

import (
	"encoding/json"
	"fmt"
	"net/http"
)

const apiKey = "2a1a925a8aac909f5bd03e7f19b41ca1"

type WeatherData struct {
    Main struct {
        Temp float64 `json:"temp"`
    } `json:"main"`
    Weather []struct {
        Description string `json:"description"`
    } `json:"weather"`
}

func getWeather(city string) (string, error) {
    url := fmt.Sprintf("http://api.openweathermap.org/data/2.5/weather?q=%s&appid=%s&units=metric", city, apiKey)
    resp, err := http.Get(url)
    if err != nil {
        return "", err
    }
    defer resp.Body.Close()

    var data WeatherData
    if err := json.NewDecoder(resp.Body).Decode(&data); err != nil {
        return "", err
    }

    weather := fmt.Sprintf("        현재 %s의 온도는 %.1f°C이며, 상태는 %s입니다.", city, data.Main.Temp, data.Weather[0].Description)
    return weather, nil
}
