package booking

import (
	"fmt"
	"time"
)

// Schedule returns a time.Time from a string containing a date.
func Schedule(date string) time.Time {
	t, _ := time.Parse("1/2/2006 15:04:05", date)

	return t
}

// HasPassed returns whether a date has passed.
func HasPassed(date string) bool {
	dateTime, _ := time.Parse("January 2, 2006 15:04:05", date)

	return time.Now().After(dateTime)
}

// IsAfternoonAppointment returns whether a time is in the afternoon.
func IsAfternoonAppointment(date string) bool {
	dateTime, _ := time.Parse("Monday, January 2, 2006 15:04:05", date)

	return dateTime.Hour() >= 12 && dateTime.Hour() < 18
}

// Description returns a formatted string of the appointment time.
func Description(date string) string {
	return "You have an appointment on " + Schedule(date).Format("Monday, January 2, 2006, at 15:04") + "."
}

// AnniversaryDate returns a Time with this year's anniversary.
func AnniversaryDate() time.Time {
	return Schedule("09/15/" + fmt.Sprintf("%d", time.Now().Year()) + " 00:00:00")
}
