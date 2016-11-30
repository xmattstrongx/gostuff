package main

import "fmt"

func main() {
	var size int64
	fmt.Scanf("%d", &size)

	phoneBook := readInPhoneBookData(size)
	names := getDataToExamine(size)

	for _, name := range names {
		if _, ok := phoneBook[name]; ok {
			fmt.Printf("%s=%d\n", name, phoneBook[name])
		} else {
			fmt.Printf("Not found\n")
		}
	}
}

func readInPhoneBookData(size int64) map[string]int64 {
	var key string
	var value int64
	var i int64
	phoneBook := make(map[string]int64)

	for i = 0; i < size; i++ {
		fmt.Scanf("%s %d", &key, &value)

		if _, ok := phoneBook[key]; !ok {
			phoneBook[key] = value
		}
	}
	return phoneBook
}

func getDataToExamine(size int64) []string {
	var i int64
	names := make([]string, size)
	for i = 0; i < size; i++ {
		fmt.Scanf("%s", &names[i])
	}
	return names
}
