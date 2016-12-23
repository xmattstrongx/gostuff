package time_conversion

import (
	"log"
	"time"
)

const inputLayout = "03:04:05PM"
const outputLayout = "15:04:05"

func ConvertTime(str string) string {

	t, err := time.Parse(inputLayout, str)
	if err != nil {
		log.Fatal(err)
	}

	return t.Format(outputLayout)
}
