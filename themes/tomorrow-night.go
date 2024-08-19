package themes

import (
	"main/drawer_theme"
)

func CreateThemeTomorrowNight() *drawer_theme.Theme {
	return drawer_theme.CreateTheme(
		"#1d1f21",
		"#cc6666",
		"#b5bd68",
		"#e6c547",
		"#e78a4e",
		"#81a2be",
		"#b294bb",
		"#70c0ba",
		"#373b41",
	)
}
