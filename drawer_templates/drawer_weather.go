package drawer_templates

// WeatherEmoji returns one or more emoji,
// generalizing the weather state by the code ww (0-99).
func WeatherEmoji(code int) string {
	switch code {
	// 00-19: No precipitation at the time of observation
	case 0, 1, 2, 3: // 0-3: Cloudiness, changes in cloudiness
		return "☁️"
	case 4, 5, 6, 7, 8: // 4-8: Haze, dust, faint dust/sand blizzard
		return "🌫️"
	case 9: // 9: Dust storm.
		return "🌪️"
	case 10, 11, 12: // 10-12: Fog (fog, fine mist)
		return "🌫️"
	case 13: // 13: Lightning is seen, no thunder
		return "🌩️"
	case 14, 15, 16: // 14-16: Precipitation is visible, but not at the station
		return "🌦️"
	case 17: // 17: Thunderstorm, but no precipitation.
		return "⛈️"
	case 18: // 18: Squalls
		return "💨"
	case 19: // 19: Vortex (tornado).
		return "🌪️"

	// 20-29: Precipitation/fog/thunderstorms WERE in the past hour, but not now
	case 20: // Drizzle or snow grains.
		return "🌧️"
	case 21: // Rain
		return "🌧️"
	case 22: // Snow
		return "❄️"
	case 23: // Rain and snow or ice pellets
		return "🌨️"
	case 24: // Freezing drizzle or freezing rain
		return "🧊🌧️"
	case 25: // Shower(s) of rain
		return "🌦️"
	case 26: // Shower(s) of snow (or rain+snow)
		return "🌨️"
	case 27: // Shower(s) of hail (or rain+hail)
		return "🌧️🧊"
	case 28: // Fog or ice fog
		return "🌫️"
	case 29: // Thunderstorm
		return "⛈️"

	// 30-39: Dusty (sandy) storm or heavy snowstorm with snow
	case 30, 31, 32, 33, 34, 35: // 30-35: Duststorm or sandstorm (moderate/severe, changes in intensity)
		return "🌪️"
	case 36, 37, 38, 39: // 36-39: Blowing snow (heavy blowing snow)
		return "🌨️"

	// 40-49: Fog or ice fog at the time of observation
	case 40, 41, 42, 43, 44, 45, 46, 47, 48, 49: // All about the fog
		return "🌫️"

	// 50-59: Drizzle.
	case 56, 57: // Freezing drizzle
		return "🧊🌧️"
	case 50, 51, 52, 53, 54, 55, 58, 59: // 50–55: Drizzle (intermittent / continuous, slight / moderate / heavy)
		return "🌧️"

	// 60–69: Rain
	case 66, 67: // Freezing rain
		return "🧊🌧️"
	case 68, 69: // Rain or drizzle + snow
		return "🌧️❄️"
	case 60, 61, 62, 63, 64, 65: // 60–65: Rain (intermittent / continuous, slight / moderate / heavy)
		return "🌧️"

	// 70-79: Solid precipitation not in the form of showers (snow, ice pellets...)
	case 76: // Diamond dust
		return "💎❄️"
	case 77: // Snow grains
		return "❄️"
	case 78: // Isolated star-like snow crystals
		return "❄️"
	case 79: // Ice pellets
		return "🧊"
	case 70, 71, 72, 73, 74, 75: // 70–75: snow (intermittent / continuous, slight / moderate / heavy)
		return "❄️"

	// 80-99: Heavy rainfall or thunderstorm (shower(s), thunderstorm)
	case 80, 81, 82: // 80–82: Rain shower(s), slight … violent
		return "🌦️"
	case 83, 84: // 83–84: Shower(s) of rain and snow
		return "🌦️❄️"
	case 85, 86: // 85–86: Snow shower(s)
		return "🌨️"
	case 87, 88, 89, 90: // 87–90: Showers of hail / snow pellets
		return "🧊"
	case 91, 92: // 91–92: Thunderstorm (slight or heavy rain)
		return "⛈️"
	case 93, 94: // 93–94: Thunderstorm with snow
		return "⛈️❄️"
	case 95, 97: // 95,97: Thunderstorm (slight/moderate or heavy) - no obvious hail.
		return "⛈️"
	case 96, 99: // 96, 99: Thunderstorm with hail
		return "⛈️🧊"
	case 98: // 98: Thunderstorm with dust/sand storm
		return "⛈️🌪️"
	}

	// If the code is out of the range 0-99
	return "❓"
}
