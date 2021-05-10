package main

import (
	"fmt"
	"time"
)

func getDaysMonthsYears(durationHours time.Duration) (int64, int64, int64) {
	var daySpan int64 = int64(durationHours.Hours()) / 24
	var months int64 = daySpan / 30
	var years int64 = daySpan / 365

	return daySpan, months, years
}

func main() {
	t := time.Now()
	fmt.Println(t)

	dob := time.Date(1979, 10, 4, 10, 4, 0, 0, time.UTC)
	fmt.Println("The day of the week was", dob.Weekday())
	fmt.Println("Was the date before now? ", dob.Before(t))
	fmt.Println("Was the date after now? ", dob.After(t))

	timeDiff := t.Sub(dob)
	fmt.Println("Time Since DOB")
	fmt.Println(timeDiff)
	fmt.Println("Minutes:", timeDiff.Minutes())
	fmt.Println("Hours:", timeDiff.Hours())

	d, m, y := getDaysMonthsYears(timeDiff)

	fmt.Println(d, " days")
	fmt.Println(m, " months")
	fmt.Println(y, " years")

	// Unix epoch
	secondsSinceUnix := time.Now().Unix()
	fmt.Println("Seconds since the Unix epoch: ", secondsSinceUnix)
	fmt.Println("Time now based on Unix epoch: ", time.Unix(secondsSinceUnix, 0))
}
