package drawer_theme

type Color string

type Theme struct {
	Black   string
	Red     string
	Green   string
	Yellow  string
	Orange  string
	Blue    string
	Magenta string
	Cyan    string
	White   string
}

func CreateTheme(
	black string,
	red string,
	green string,
	yellow string,
	orange string,
	blue string,
	magenta string,
	cyan string,
	white string,
) *Theme {
	return &Theme{
		Black:   black,
		Red:     red,
		Green:   green,
		Yellow:  yellow,
		Orange:  orange,
		Blue:    blue,
		Magenta: magenta,
		Cyan:    cyan,
		White:   white,
	}
}
