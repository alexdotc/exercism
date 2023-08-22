package booking

import "fmt"
import "time"

// Schedule returns a time.Time from a string containing a date.
func Schedule(date string) time.Time {
	layout := "1/2/2006 15:04:05"
	t, _ := time.Parse(layout, date)
	return t
}

// HasPassed returns whether a date has passed.
func HasPassed(date string) bool {
	layout := "January 2, 2006 15:04:05"
	t, _ := time.Parse(layout, date)
	y := time.Now()
	return t.Compare(y) < 0
}

// IsAfternoonAppointment returns whether a time is in the afternoon.
func IsAfternoonAppointment(date string) bool {
	layout := "Monday, January 2, 2006 15:04:05"
	t, _ := time.Parse(layout, date)
	switch {
		case t.Hour() == 18 && t.Minute() == 0 && t.Second() == 0:
			return true
		case t.Hour() >= 12 && t.Hour() < 18:
			return true
	}
	return false
}

// Description returns a formatted string of the appointment time.
func Description(date string) string {
	layout := "Monday, January 2, 2006, at 15:04"
	return fmt.Sprintf("You have an appointment on %s.", Schedule(date).Format(layout))
}

// AnniversaryDate returns a Time with this year's anniversary.
func AnniversaryDate() time.Time {
	return time.Date(time.Now().Year(), 9, 15, 0, 0, 0, 0, time.UTC)
}
