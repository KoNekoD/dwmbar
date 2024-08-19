package themes

import (
	"main/drawer_theme"
)

func CreateThemeSweetMars() *drawer_theme.Theme {
	return drawer_theme.CreateTheme(
		"#1a1e21",
		"#cc241d",
		"#98971a",
		"#d79921",
		"#e78a4e",
		"#458588",
		"#b16286",
		"#689d6a",
		"#a89984",
	)
}
