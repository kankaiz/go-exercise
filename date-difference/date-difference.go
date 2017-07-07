// Description:
// Create an application that can read in pairs of dates in the following format -
// DD MM YYYY, DD MM YYYY
// Validate the input data, and compute the difference between the two dates in days.
// Output of the application should be of the form -
// DD MM YYYY, DD MM YYYY, difference
// Where the first date is the earliest, the second date is the latest and the difference is the number of
// days.
// Input can be from a file, or from standard input, as the developer chooses.
// Provide test data to exercise the application.

package main

import (
	"fmt"
	"time"
)

// MyDate date format is 22 06 2017
// month 03 or 3?
// day > 31...
type MyDate struct {
	day   int
	month time.Month
	year  int
}

// validDate checks the input date format
func validDate(date MyDate) bool {
	possibleDate := time.Date(date.year, date.month, date.day, 0, 0, 0, 0, time.UTC)
	// if the parsed date doesn't change then it is a valid date
	if possibleDate.Day() == date.day &&
		possibleDate.Month() == date.month &&
		possibleDate.Year() == date.year {
		return true
	}
	return false
}

func main() {
	var d1, m1, y1, d2, m2, y2 int
	//read from STDIN
	fmt.Scanf("%v %d %v, %v %d %v", &d1, &m1, &y1, &d2, &m2, &y2)
	date1 := MyDate{d1, time.Month(m1), y1}
	date2 := MyDate{d2, time.Month(m2), y2}

	if !validDate(date1) || !validDate(date2) {
		// invalid date format is
		fmt.Println("Please input a valid date")
	} else {
		// valid date format
		// business logic here
		fmt.Println(date1.day, date1.month, date2.year)
		time1 := time.Date(date1.year, date1.month, date1.day, 0, 0, 0, 0, time.UTC)
		time2 := time.Date(date2.year, date2.month, date2.day, 0, 0, 0, 0, time.UTC)

		if date1.year >= date2.year {
			fmt.Println((int(time1.Sub(time2) / (24 * time.Hour))))
		} else {
			fmt.Println((int(time2.Sub(time1) / (24 * time.Hour))))
		}
	}
}
