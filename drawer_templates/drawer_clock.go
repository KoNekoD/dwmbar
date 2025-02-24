package drawer_templates

import "time"

var monthRu = map[time.Month]string{
	time.January:   "Янв",
	time.February:  "Фев",
	time.March:     "Март",
	time.April:     "Апр",
	time.May:       "Май",
	time.June:      "Июнь",
	time.July:      "Июль",
	time.August:    "Авг",
	time.September: "Сен",
	time.October:   "Окт",
	time.November:  "Ноя",
	time.December:  "Дек",
}

var weekdayRu = map[time.Weekday]string{
	time.Sunday:    "Вс",
	time.Monday:    "Пн",
	time.Tuesday:   "Вт",
	time.Wednesday: "Ср",
	time.Thursday:  "Чт",
	time.Friday:    "Пт",
	time.Saturday:  "Сб",
}

var monthEn = map[time.Month]string{
	time.January:   "Jan",
	time.February:  "Feb",
	time.March:     "Mar",
	time.April:     "Apr",
	time.May:       "May",
	time.June:      "Jun",
	time.July:      "Jul",
	time.August:    "Aug",
	time.September: "Sep",
	time.October:   "Oct",
	time.November:  "Nov",
	time.December:  "Dec",
}

var weekdayEn = map[time.Weekday]string{
	time.Sunday:    "Sun",
	time.Monday:    "Mon",
	time.Tuesday:   "Tue",
	time.Wednesday: "Wed",
	time.Thursday:  "Thu",
	time.Friday:    "Fri",
	time.Saturday:  "Sat",
}

func GetClockMonthRu(month time.Month) string {
	if val, ok := monthRu[month]; ok {
		return val
	}
	return "invalid month"
}

func GetClockWeekDayRu(weekday time.Weekday) string {
	if val, ok := weekdayRu[weekday]; ok {
		return val
	}
	return "invalid weekday"
}

func GetClockMonthEn(month time.Month) string {
	if val, ok := monthEn[month]; ok {
		return val
	}
	return "invalid month"
}

func GetClockWeekDayEn(weekday time.Weekday) string {
	if val, ok := weekdayEn[weekday]; ok {
		return val
	}
	return "invalid weekday"
}
