package main

import (
	"fmt"
	"time"
)

func AddGigasecond(t time.Time) time.Time {
	t = t.Add(time.Second * 1000000000)
	return t
}

func main() {
	layout := "January 2, 2006 at 15:04:05 (03:04:05PM)"
	date := "January 24, 2015 at 22:00:00 (10:00:00PM)"
	t, _ := time.Parse(layout, date)
	fmt.Println(t)
	fmt.Println(t.Format(layout))
	fmt.Println(AddGigasecond(t))
	fmt.Println(AddGigasecond(t).Format(layout))
}

// Gigasecond / time
