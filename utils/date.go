package utils

import (
	"strconv"
	"time"
)

func GetCurentYear() string {
	// Get the current date and time
	current_time := time.Now()

	// Extract the year from the current date
	current_year := current_time.Year()

	// Convert the year to a string using strconv.Itoa()
	return strconv.Itoa(current_year)
}
