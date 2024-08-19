package themes

import (
	"main/drawer_theme"
)

func CreateThemeDracula() *drawer_theme.Theme {
	return drawer_theme.CreateTheme(
		"#21222c",
		"#ff5555",
		"#50fa7b",
		"#f1fa8c",
		"#e78a4e",
		"#bd93f9",
		"#ff79c6",
		"#8be9fd",
		"#f8f8f2",
	)
}
