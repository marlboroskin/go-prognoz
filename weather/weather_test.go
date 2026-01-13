package weather

import (
	"strings"
	"testing"
	"weather/geo"
)

func TestGetWeather(t *testing.T) {
	expected := "Stavropol"
	geoData := geo.Geodata{
		City: expected,
	}
	format := 3
	result, err := GetWeather(geoData, format)
	if err != nil {
		t.Errorf("Пришла ошибка %v", err)
	}
	if !strings.Contains(result, expected) {
		t.Errorf("Ожидалось %v, получение %v", expected, result)
	}
}

func TestGetWeatherWrongFormat(t *testing.T) {
	testCases := []struct {
		name   string
		format int
	}{
		{name: "Big format", format: 147},
		{name: "Minus format", format: -1},
	}

	for _, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) {
			geoData := geo.Geodata{City: "Stavropol"}
			_, err := GetWeather(geoData, tc.format)
			if err != ErrWrongFormat {
				t.Errorf("Ожидалось %v, получено %v", ErrWrongFormat, err)
			}
		})
	}
}
