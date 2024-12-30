package main

import (
	"fmt"
	"time"
)

func main() {
	// var dateAndTime time.Time = time.Now()
	// fmt.Println(dateAndTime)

	// var date string = dateAndTime.Format("02-01-06")
	// var time string = dateAndTime.Format("03:04:05 PM")
	// fmt.Println(date)
	// fmt.Println(time)

	// var dateStr string = "02/12/2025"

	// realDate, _ := time.Parse("02/01/2006", dateStr)

	// fmt.Printf("%T", realDate)

	currentDate := time.Now()

	tomoro := currentDate.Add(24 * time.Hour)

	fmt.Println(tomoro)

}
