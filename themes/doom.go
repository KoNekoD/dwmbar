package themes

import (
	"main/drawer_theme"
)

func CreateThemeDoom() *drawer_theme.Theme {
	return drawer_theme.CreateTheme(
		"#1c1f24",
		"#ff6c6b",
		"#98be65",
		"#da8548",
		"#c51605",
		"#51afef",
		"#c678dd",
		"#5699af",
		"#202328",
	)
}
