package components

var (
	jCalendarE int32 = 1917
	gCalendarB int32 = 1918
)

func CalendarType(year int32) string {
	if year <= jCalendarE {
		return "julian"

	} else if year >= gCalendarB {
		return "gregorian"
	}

	return ""
}

func IsLeap(year int32) bool {
	var isLeap func(year int32) bool

	if year <= jCalendarE {
		isLeap = func(year int32) bool {
			return year%4 == 0
		}

	} else if year >= gCalendarB {
		isLeap = func(year int32) bool {
			return year%400 == 0 || year%4 == 0 && year%100 != 0
		}
	}

	return isLeap(year)
}
