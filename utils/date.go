package utils

import (
	"strconv"
	"time"
)

func GetCurentYear() string {
	// Get the current date and time
	currentTime := time.Now()

	// Extract the year from the current date
	currentYear := currentTime.Year()

	// Convert the year to a string using strconv.Itoa()
	return strconv.Itoa(currentYear)
}
