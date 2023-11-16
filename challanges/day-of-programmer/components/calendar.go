package components

var Calendar = map[string]map[string]interface{}{
	"january": {
		"from": int32(1),
		"to":   int32(31),
	},
	"february": {
		"leap": map[string]interface{}{
			"from": int32(1),
			"to":   int32(29),
		},
		"nonLeap": map[string]interface{}{
			"from": int32(1),
			"to":   int32(28),
		},
		"transitionYear": map[string]interface{}{
			"from": int32(14),
			"to":   int32(28),
		},
	},
	"march": {
		"from": int32(1),
		"to":   int32(31),
	},
	"april": {
		"from": int32(1),
		"to":   int32(30),
	},
	"may": {
		"from": int32(1),
		"to":   int32(31),
	},
	"june": {
		"from": int32(1),
		"to":   int32(30),
	},
	"july": {
		"from": int32(1),
		"to":   int32(31),
	},
	"august": {
		"from": int32(1),
		"to":   int32(31),
	},
	"september": {
		"from": int32(1),
		"to":   int32(30),
	},
	"october": {
		"from": int32(1),
		"to":   int32(31),
	},
	"november": {
		"from": int32(1),
		"to":   int32(30),
	},
	"december": {
		"from": int32(1),
		"to":   int32(31),
	},
}

func ListMonths() []string {
	return []string{
		"january",
		"february",
		"march",
		"april",
		"may",
		"june",
		"july",
		"august",
		"september",
		"october",
		"november",
		"december",
	}
}
