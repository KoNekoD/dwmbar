package drawer_templates

import "time"

func GetClockMonthRu(month time.Month) string {
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

func GetClockWeekDayRu(weekday time.Weekday) string {
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

func GetClockMonthEn(month time.Month) string {
	var clockMonth string

	switch month {
	case time.January:
		clockMonth = "Jan"
	case time.February:
		clockMonth = "Feb"
	case time.March:
		clockMonth = "Mar"
	case time.April:
		clockMonth = "Apr"
	case time.May:
		clockMonth = "May"
	case time.June:
		clockMonth = "Jun"
	case time.July:
		clockMonth = "Jul"
	case time.August:
		clockMonth = "Aug"
	case time.September:
		clockMonth = "Sep"
	case time.October:
		clockMonth = "Oct"
	case time.November:
		clockMonth = "Nov"
	case time.December:
		clockMonth = "Dec"
	default:
		clockMonth = "invalid month"
	}

	return clockMonth
}

func GetClockWeekDayEn(weekday time.Weekday) string {
	var clockWeekDay string

	switch weekday {
	case time.Sunday:
		clockWeekDay = "Sun"
	case time.Monday:
		clockWeekDay = "Mon"
	case time.Tuesday:
		clockWeekDay = "Tue"
	case time.Wednesday:
		clockWeekDay = "Wed"
	case time.Thursday:
		clockWeekDay = "Thu"
	case time.Friday:
		clockWeekDay = "Fri"
	case time.Saturday:
		clockWeekDay = "Sat"
	default:
		clockWeekDay = "invalid weekday"
	}

	return clockWeekDay
}
