package themes

import (
	"main/drawer_theme"
)

func CreateThemeTomorrowDark() *drawer_theme.Theme {
	return drawer_theme.CreateTheme(
		"#1D1F21",
		"#a54242",
		"#8c9440",
		"#de935f",
		"#e78a4e",
		"#5f819d",
		"#85678f",
		"#5e8d87",
		"#707880",
	)
}
