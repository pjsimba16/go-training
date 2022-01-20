package booking

import (
	"strconv"
	"time"
)

// Schedule returns a time.Time from a string containing a date
func Schedule(date string) time.Time {
	t, _ := time.Parse("1/2/2006 15:04:05", date)
	return t
}

// HasPassed returns whether a date has passed
func HasPassed(date string) bool {
	t, _ := time.Parse("January 2, 2006 15:04:05", date)
	if t.Before(time.Now()) {
		return true
	} else {
		return false
	}
}

// IsAfternoonAppointment returns whether a time is in the afternoon
func IsAfternoonAppointment(date string) bool {
	t, _ := time.Parse("Monday, January 2, 2006 15:04:05", date)
	h1, _ := time.Parse("15:04:05", "12:00:00")
	h2, _ := time.Parse("15:04:05", "18:00:00")
	hour := t.Hour()
	if (hour >= h1.Hour()) && (hour < h2.Hour()) {
		return true
	} else {
		return false
	}

}

// Description returns a formatted string of the appointment time
func Description(date string) string {
	time := Schedule(date)
	year, month, day := time.Date()
	dateString := (month.String()) + " " + strconv.Itoa(day) + ", " + strconv.Itoa(year) + ","
	weekday := time.Weekday().String()
	return "You have an appointment on " + weekday + ", " + dateString + " at " + strconv.Itoa(time.Hour()) + ":" + strconv.Itoa(time.Minute()) + "."
}

// AnniversaryDate returns a Time with this year's anniversary
func AnniversaryDate() time.Time {
	anniversaryDate := time.Date(time.Now().Year(), 9, 15, 00, 00, 00, 0, time.UTC)
	return anniversaryDate
}
