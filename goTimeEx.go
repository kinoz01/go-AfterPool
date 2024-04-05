package main

import (
	"fmt"
	"time"
	//"strconv"
)

const (
	layout  = "1/2/2006 15:04:05"
	layout2 = "January 2, 2006 15:04:05"
	layout3 = "Monday, January 2, 2006 15:04:05"
	layout4 = "Monday, January 2, 2006, at 15:04."
	layout5 = "2006-01-02 15:04:05 -0700 MST"
	layout6 = "2006"
)

// Schedule returns a time.Time from a string containing a date.
func Schedule(date string) time.Time {
	t, _ := time.Parse(layout, date)
	// layout and date should have the same format to
	// print date in this format
	// 2024-09-15 00:00:00 +0000 UTC.
	return t
}

// HasPassed returns whether a date has passed.
func HasPassed(date string) bool {
	t, _ := time.Parse(layout2, date)
	return t.Before(time.Now())	
}

// IsAfternoonAppointment returns whether a time is in the afternoon.
func IsAfternoonAppointment(date string) bool {
	t, _ := time.Parse(layout3, date)
	if t.Hour() >= 12 && t.Hour() < 18 {
		return true
	}
	return false
}

// Description returns a formatted string of the appointment time.
func Description(date string) string {
	t, _ := time.Parse(layout, date)
	// fmt.Println(t)
	return "You have an appointment on " + t.Format(layout4)
	// Change the layout in return and see the format changes in the output
}

// AnniversaryDate returns a Time with this year's anniversary.
func AnniversaryDate() time.Time {
	timeNow := time.Now()
	year := timeNow.Format(layout6)
	date := year + "-09-15 00:00:00 +0000 UTC"
	t, _ := time.Parse(layout5, date)
	return t
	// Method 1: // t, _ := time.Parse("2006-01-2", fmt.Sprintf("%d-09-15", time.Now().Year()))
	// Method 2: // return time.Date(time.Now().Year(), time.September, 15, 0, 0, 0, 0, time.UTC)
}

func main() {
	fmt.Println(Schedule("7/25/2019 13:45:00"))

	fmt.Println(HasPassed("December 9, 2026 18:59:59"))

	fmt.Println(IsAfternoonAppointment("Thursday, July 25, 2019 15:45:00"))

	fmt.Println(Description("7/25/2019 13:45:00"))

	fmt.Println(AnniversaryDate())
}

//  Booking up for Beauty
