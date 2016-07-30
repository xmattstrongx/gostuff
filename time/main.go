package main

import (
	"fmt"
	"time"
)

//20160727204520.005Z
func main() {
	bullshit := "20160727204520.005Z"
	t, _ := time.Parse("20060102150405.000Z", bullshit)
	fmt.Printf("After converting Bullshit timestamp: %s\n", t)

	fmt.Println("Unix format:", t.Format(time.UnixDate))
}
