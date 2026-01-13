package weather

import (
	"errors"
	"fmt"
	"io"
	"net/http"
	"net/url"
	"weather/geo"
)

var ErrWrongFormat = errors.New("ERROR_FORMAT")

func GetWeather(geo geo.Geodata, format int) (string, error) {
	if format < 0 || format > 4 {
		return "", ErrWrongFormat
	}
	baseUrl, err := url.Parse("https://wttr.in/" + url.PathEscape(geo.City))
	if err != nil {
		fmt.Println(err.Error())
		return "", errors.New("ERROR_URL")
	}
	params := url.Values{}
	params.Set("format", fmt.Sprintf("%d", format))
	baseUrl.RawQuery = params.Encode()

	resp, err := http.Get(baseUrl.String())
	if err != nil {
		fmt.Println(err.Error())
		return "", errors.New("ERROR_HTTP")
	}
	defer resp.Body.Close()
	body, err := io.ReadAll(resp.Body)
	if err != nil {
		fmt.Println(err.Error())
		return "", errors.New("ERROR_READBODY")
	}
	return string(body), nil
}
