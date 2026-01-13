package main

import (
	"flag"
	"fmt"
	"weather/geo"
	"weather/weather"
)

func main() {
	city := flag.String("city", "", "Город пользователя")
	format := flag.Int("format", 1, "Формат вывода погоды")

	flag.Parse()

	fmt.Println(*city)

	geoData, err := geo.GetMyLocation(*city)
	if err != nil {
		fmt.Println("Ошибка: не удалось определить местоположение:", err)
		return
	}
	weatherData, err := weather.GetWeather(*geoData, *format)
	fmt.Println(weatherData)
}
