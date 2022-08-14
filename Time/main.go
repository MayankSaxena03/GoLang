package main

import (
	"fmt"
	"time"
)

func main() {
	present := time.Now()
	fmt.Println("Present Time = ", present)
	/* Time Formats */
	/* The values of date, time and day should remain constant */
	fmt.Println("Time Format = ", present.Format("2006-01-02"))
	fmt.Println("Time Format = ", present.Format("15:04:05"))
	fmt.Println("Time Format = ", present.Format("Mon Jan 2 15:04:05 2006"))
	fmt.Println("Time Format = ", present.Format("Monday January 2 15:04:05 2006"))
	fmt.Println("Time Format = ", present.Format("2006-01-02 15:04:05"))
	fmt.Println("Time Format = ", present.Format("2006-01-02 15:04:05 -0700 MST"))
	fmt.Println("Time Format = ", present.Format("Monday 2006-01-02 15:04"))
	/*
		Important Constants
			Year: "2006" "06"
			Month: "Jan" "January" "01" "1"
			Day of the week: "Mon" "Monday"
			Day of the month: "2" "_2" "02"
			Day of the year: "__2" "002"
			Hour: "15" "3" "03" (PM or AM)
			Minute: "4" "04"
			Second: "5" "05"
			AM/PM mark: "PM"
	*/

	/* Set your time */
	myTime := time.Date(2001, time.March, 3, 07, 30, 50, 0, time.Local) /* Format to create your time variable */
	fmt.Println("My Time = ", myTime)
	fmt.Println("My Time Format = ", myTime.Format("Monday 2006-01-02 15:04:05"))
}
