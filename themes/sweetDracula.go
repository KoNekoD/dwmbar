package themes

import (
	"main/drawer_theme"
)

func CreateThemeSweetDracula() *drawer_theme.Theme {
	return drawer_theme.CreateTheme(
		"#1a1e21",
		"#ff6c6b",
		"#50fa7b",
		"#f1fa8c",
		"#e78a4e",
		"#bd93f9",
		"#ff85ab",
		"#8be9fd",
		"#f8f8f2",
	)
}
