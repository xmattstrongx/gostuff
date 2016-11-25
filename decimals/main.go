package main

import (
	"fmt"
)

func main() {
	//find first decimal value
	var f float64 = 123456.789
	new := int((f - float64(int(f))) * 10)
	fmt.Println(new)
}
