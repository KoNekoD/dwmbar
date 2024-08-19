package themes

import (
	"main/drawer_theme"
)

func CreateThemeTomorrow() *drawer_theme.Theme {
	return drawer_theme.CreateTheme(
		"#1D1F21",
		"#fb4934",
		"#98971a",
		"#d7992a",
		"#e78a4e",
		"#458588",
		"#b16286",
		"#689d6a",
		"#373b41",
	)
}
