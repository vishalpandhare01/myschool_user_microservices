package validation

import (
	"regexp"
	"time"
)

func ValidateDate(date string) bool {
	// Regular expression to match the DD-MM-YYYY format
	var validDateFormat = `^(\d{2})-(\d{2})-(\d{4})$`
	re := regexp.MustCompile(validDateFormat)

	// Check if the date matches the format
	if !re.MatchString(date) {
		return false
	}

	// Try to parse the date in DD-MM-YYYY format
	parsedDate, err := time.Parse("02-01-2006", date)
	if err != nil {
		return false
	}

	// Check if the parsed date matches the original date (to handle edge cases like 31-02-2020)
	return parsedDate.Format("02-01-2006") == date
}
