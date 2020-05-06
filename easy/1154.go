package easy

import (
	"fmt"
	"time"
)

func dayOfYear(date string) int {
	t, err := time.Parse(time.RFC3339, date+"T00:00:00Z")
	if err != nil {
		fmt.Printf("incorrect date format: %v", date)
		return 0
	}

	return t.YearDay()
}
