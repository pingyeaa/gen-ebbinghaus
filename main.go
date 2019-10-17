package main

import (
	"fmt"
	"os"
	"time"
)

func main() {
	var date string
	for key, value := range os.Args {
		if key == 1 {
			date = value
		}
	}

	var now time.Time
	if date == "" {
		now = time.Now()
	} else {
		now, _ = time.Parse("2006-01-02", date)
	}

	var output string
	days := []int{1, 2, 4, 7, 15, 30, 60, 180, 365}
	for key, value := range days {
		newDate := now.AddDate(0, 0, value)
		if key == 0 {
			output += newDate.Format("2006.01.02")
		} else {
			output += "/" + newDate.Format("2006.01.02")
		}
	}

	fmt.Println(output)
}
