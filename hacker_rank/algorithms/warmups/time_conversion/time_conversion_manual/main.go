package main

import (
	"fmt"
	"strconv"
)

func main() {
	var time string
	fmt.Scanf("%s", &time)

	meridiem, timeWithoutMeridiem := getMeridiemAndTimeWithoutMeridiem(time)

	if meridiem == "PM" {
		timeWithoutMeridiem = getMilitaryPMTime(timeWithoutMeridiem)
	} else {
		timeWithoutMeridiem = getMilitaryAMTime(timeWithoutMeridiem)
	}

	fmt.Println(timeWithoutMeridiem)
}

func getMeridiemAndTimeWithoutMeridiem(time string) (string, string) {
	meridiem := time[len(time)-2:]
	timeWithoutMeridiem := time[:len(time)-2]
	return meridiem, timeWithoutMeridiem
}

func getMilitaryPMTime(time string) string {
	hours := time[:2]
	numericHours, _ := strconv.ParseInt(hours, 10, 0)
	if numericHours < 12 {
		numericHours += 12
	}
	return strconv.Itoa(int(numericHours)) + time[2:]
}

func getMilitaryAMTime(time string) string {
	if time[:2] == "12" {
		return "00" + time[2:]
	}
	return time
}
