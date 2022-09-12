package main

import (
	"fmt"
	"time"
)

func main() {
	fmt.Println("Time module")

	// In time module Now function is used to get current time. It will return a tim.Time type.
	presentTime := time.Now()
	// It will print everything about today with precision in nano seconds.
	fmt.Println(presentTime)
	// Formatting the time is quite challenging. You won't be getting correct details in output with below string in format
	fmt.Println(presentTime.Format("03-05-2007 10:04:05 Tuesday"))
	// You have to use exactly the same string to get correct results for this format
	// In date's place, for correct date to be in output, use exact same string "01-02-2006"
	// For time, use "15:04:05"
	// For weekday, use "Monday"
	fmt.Println(presentTime.Format("01-02-2006 15:04:05 Monday"))

	// Above comments are true for everywhere while you are using formatting time.
	createDate := time.Date(2020, time.August, 11, 23, 12, 0, 0, time.UTC)
	fmt.Println(createDate.Format("01-02-2006 15:04:05 Monday"))

	// Or you can use predefined formats well
	/*
		UnixDate	“Mon Jan _2 15:04:05 MST 2006”
		RubyDate	“Mon Jan 02 15:04:05 -0700 2006”
		RFC822	“02 Jan 06 15:04 MST”
		RFC822Z	“02 Jan 06 15:04 -0700”
		RFC850	“Monday, 02-Jan-06 15:04:05 MST”
		RFC1123	“Mon, 02 Jan 2006 15:04:05 MST”
		RFC1123Z	“Mon, 02 Jan 2006 15:04:05 -0700”
		RFC3339	“2006-01-02T15:04:05Z07:00”
		RFC3339Nano	“2006-01-02T15:04:05.999999999Z07:00”
	*/
	fmt.Println(createDate.Format(time.UnixDate))
	fmt.Println(createDate.Format(time.RFC822))
}
