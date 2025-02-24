package weather_state

import (
	"encoding/json"
	"fmt"
	"io"
	"net/http"
	"sync"
	"time"
)

var (
	cache      Stats
	cacheMutex sync.Mutex
	lastUpdate time.Time
)

type WeatherResponse struct {
	Latitude             float64 `json:"latitude"`
	Longitude            float64 `json:"longitude"`
	GenerationtimeMs     float64 `json:"generationtime_ms"`
	UtcOffsetSeconds     int     `json:"utc_offset_seconds"`
	Timezone             string  `json:"timezone"`
	TimezoneAbbreviation string  `json:"timezone_abbreviation"`
	Elevation            float64 `json:"elevation"`
	CurrentUnits         struct {
		Time          string `json:"time"`
		Interval      string `json:"interval"`
		Temperature2M string `json:"temperature_2m"`
		WeatherCode   string `json:"weather_code"`
	} `json:"current_units"`
	Current struct {
		Time          string  `json:"time"`
		Interval      int     `json:"interval"`
		Temperature2M float64 `json:"temperature_2m"`
		WeatherCode   int     `json:"weather_code"`
	} `json:"current"`
}

type Stats struct {
	Temperature string
	Code        int
}

func Get() (Stats, error) {
	cacheMutex.Lock()
	defer cacheMutex.Unlock()

	if time.Since(lastUpdate) < 30*time.Minute {
		return cache, nil
	}

	if err := fetchWeather(); err != nil {
		return cache, err
	}

	lastUpdate = time.Now()
	return cache, nil
}

func fetchWeather() error {
	ip, err := getPublicIP()
	if err != nil {
		return err
	}

	loc, err := getLocationByIP(ip)
	if err != nil {
		return err
	}

	weatherURL := fmt.Sprintf("https://api.open-meteo.com/v1/forecast?latitude=%.4f&longitude=%.4f&current=temperature_2m,weather_code", loc.Latitude, loc.Longitude)
	resp, err := http.Get(weatherURL)
	if err != nil {
		return fmt.Errorf("error getting weather data: %v", err)
	}
	defer func() { _ = resp.Body.Close() }()
	var data WeatherResponse
	body, err := io.ReadAll(resp.Body)
	if err := json.Unmarshal(body, &data); err != nil {
		return fmt.Errorf("error decoding weather data: %v", err)
	}

	temp := data.Current.Temperature2M
	code := data.Current.WeatherCode

	cache = Stats{Temperature: fmt.Sprintf("%.1f°C", temp), Code: code}

	return nil
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
	url := fmt.Sprintf("http://ip-api.com/json/%s?fields=lat,lon", ip)
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

	return &Location{Latitude: result.Lat, Longitude: result.Lon}, nil
}
