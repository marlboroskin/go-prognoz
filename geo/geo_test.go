package geo

import (
	"testing"
)

func TestGetMyLocation(t *testing.T) {
	city := "London"
	expectedCity := "London"

	got, err := GetMyLocation(city) // ✅ Здесь всё ок — та же папка
	if err != nil {
		t.Fatalf("GetMyLocation вернул ошибку: %v", err)
	}
	if got.City != expectedCity {
		t.Errorf("Ожидался город %q, получили %q", expectedCity, got.City)
	}
}

func TestGetMyLocationNoCity(t *testing.T) {
	city := "Londonssadww"
	_, err := GetMyLocation(city) // ✅ Не geo.GetMyLocation
	if err != ErrorNoCity {       // ✅ Не geo.ErrorNoCity
		t.Fatalf("Ожидалась %v, получили: %v", ErrorNoCity, err)
	}
}
