package drawer_templates

import "time"

func GetClockMonth(month time.Month) string {
	var clockMonth string

	switch month {
	case time.January:
		clockMonth = "Янв"
	case time.February:
		clockMonth = "Фев"
	case time.March:
		clockMonth = "Март"
	case time.April:
		clockMonth = "Апр"
	case time.May:
		clockMonth = "Май"
	case time.June:
		clockMonth = "Июнь"
	case time.July:
		clockMonth = "Июль"
	case time.August:
		clockMonth = "Авг"
	case time.September:
		clockMonth = "Сен"
	case time.October:
		clockMonth = "Окт"
	case time.November:
		clockMonth = "Ноя"
	case time.December:
		clockMonth = "Дек"
	default:
		clockMonth = "invalid month"
	}

	return clockMonth
}

func GetClockWeekDay(weekday time.Weekday) string {
	var clockWeekDay string

	switch weekday {
	case time.Sunday:
		clockWeekDay = "Вс"
	case time.Monday:
		clockWeekDay = "Пн"
	case time.Tuesday:
		clockWeekDay = "Вт"
	case time.Wednesday:
		clockWeekDay = "Ср"
	case time.Thursday:
		clockWeekDay = "Чт"
	case time.Friday:
		clockWeekDay = "Пт"
	case time.Saturday:
		clockWeekDay = "Сб"
	default:
		clockWeekDay = "invalid weekday"
	}

	return clockWeekDay
}
