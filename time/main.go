package main

import (
	"fmt"
	"time"
)

//20160727204520.005Z
func main() {
	bullshit := "20160927205458.988Z"
	t, _ := time.Parse("20060102150405.000Z", bullshit)
	fmt.Printf("After converting Bullshit timestamp: %s\n", t)

	fmt.Println("Unix format:", t.Format(time.UnixDate))

	bullshit = "20160927212245Z"
	t, _ = time.Parse("20060102150405Z", bullshit)
	fmt.Printf("After converting Bullshit timestamp: %s\n", t)

	fmt.Println("Unix format:", t.Format(time.UnixDate))

	formats := []string{
		"01-02-2006",
		"01/02/2006",
		"20060102",
		"2006-01-02",
		"2006-1-2",
		"1-2-2006",
		"1/2/2006",
	}

	var openDJTime string
	var openDJTimeMatt string

	dateValue := "05-08-2016"

	for _, value := range formats {
		t, err := time.Parse(value, dateValue)
		if err == nil {
			openDJTimeMatt = t.Format("20060102150405.000Z")
			openDJTime = fmt.Sprintf("%04d%02d%02d000000Z", t.Year(), t.Month(), t.Day())
		}
	}

	fmt.Printf("After converting openDJTime timestamp: %s\n", openDJTime)
	fmt.Printf("After converting openDJTime MATT timestamp: %s\n", openDJTimeMatt)

}
