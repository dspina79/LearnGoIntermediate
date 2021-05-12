package main

import (
	"fmt"
	"time"
)

func printDate(t time.Time) string {
	return t.Format("Jan _2 2006")
}

func printFullDate(t time.Time) string {
	tf := "Mon Jan _2 2006 3:04:05 PM"
	return t.Format(tf)
}

func main() {
	p := fmt.Println
	t := time.Now()
	te, e := time.Parse("Jan _2 2006 3:04PM", "May 10 2020 2:18AM")

	if e != nil {
		p("Error", e)
	}

	p(t.Format(time.RFC3339))
	p(t.Format("3:04PM"))
	p(t.Format("Mon Jan _2 2006"))
	// if you use another format
	p("If you use an incorrect format")
	p(t.Format("Tue Jan 13 2000"))
	p(printDate(t))
	p(printFullDate(t))
	p("Other Date")
	p(printDate(te))
	p(printFullDate(te))

}
