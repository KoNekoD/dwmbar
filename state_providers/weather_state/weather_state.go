package weather_state

import (
	"encoding/json"
	"fmt"
	"io"
	"net/http"
	"sync"
	"time"
)

type Stats struct {
	Temperature string
}

var (
	cache      *Stats
	cacheMutex sync.Mutex
	lastUpdate time.Time
)

func Get() (*Stats, error) {
	cacheMutex.Lock()
	defer cacheMutex.Unlock()

	if time.Since(lastUpdate) < 30*time.Minute && cache != nil {
		return cache, nil
	}

	weather, err := fetchWeather()
	if err != nil {
		return nil, err
	}

	cache = weather
	lastUpdate = time.Now()
	return cache, nil
}

func fetchWeather() (*Stats, error) {
	ip, err := getPublicIP()
	if err != nil {
		return nil, err
	}

	loc, err := getLocationByIP(ip)
	if err != nil {
		return nil, err
	}

	temp, err := getTemperature(loc.Latitude, loc.Longitude)
	if err != nil {
		return nil, err
	}

	return &Stats{
		Temperature: fmt.Sprintf("%.1f°C", temp),
	}, nil
}

type Location struct {
	Latitude  float64
	Longitude float64
}

func getPublicIP() (string, error) {
	resp, err := http.Get("https://api64.ipify.org?format=text")
	if err != nil {
		return "", fmt.Errorf("error getting public IP: %v", err)
	}
	defer func() { _ = resp.Body.Close() }()

	ip, err := io.ReadAll(resp.Body)
	if err != nil {
		return "", fmt.Errorf("error reading response: %v", err)
	}

	return string(ip), nil
}

func getLocationByIP(ip string) (*Location, error) {
	url := fmt.Sprintf("https://ip-api.com/json/%s?fields=lat,lon", ip)
	resp, err := http.Get(url)
	if err != nil {
		return nil, fmt.Errorf("error getting location: %v", err)
	}
	defer func() { _ = resp.Body.Close() }()

	var result struct {
		Lat float64 `json:"lat"`
		Lon float64 `json:"lon"`
	}

	body, err := io.ReadAll(resp.Body)
	if err := json.Unmarshal(body, &result); err != nil {
		return nil, fmt.Errorf("error decoding location: %v", err)
	}

	if result.Lat == 0 || result.Lon == 0 {
		return nil, fmt.Errorf("invalid location data")
	}

	return &Location{
		Latitude:  result.Lat,
		Longitude: result.Lon,
	}, nil
}

func getTemperature(lat, lon float64) (float64, error) {
	weatherURL := fmt.Sprintf("https://api.open-meteo.com/v1/forecast?latitude=%.4f&longitude=%.4f&current_weather=true", lat, lon)
	resp, err := http.Get(weatherURL)
	if err != nil {
		return 0, fmt.Errorf("error getting weather data: %v", err)
	}
	defer func() { _ = resp.Body.Close() }()

	var data struct {
		CurrentWeather struct {
			Temperature float64 `json:"temperature"`
		} `json:"current_weather"`
	}

	body, err := io.ReadAll(resp.Body)
	if err := json.Unmarshal(body, &data); err != nil {
		return 0, fmt.Errorf("error decoding weather data: %v", err)
	}

	return data.CurrentWeather.Temperature, nil
}
