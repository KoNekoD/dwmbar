package themes

import (
	"main/drawer_theme"
)

func CreateThemeGruvBox() *drawer_theme.Theme {
	return drawer_theme.CreateTheme(
		"#1d2021",
		"#ea6962",
		"#a9b665",
		"#d8a657",
		"#e78a4e",
		"#7daea3",
		"#d3869b",
		"#89b482",
		"#d4be98",
	)
}
