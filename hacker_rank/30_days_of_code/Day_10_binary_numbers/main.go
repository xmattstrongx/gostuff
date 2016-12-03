package main

import "fmt"

func main() {
	var baseTenInput int
	fmt.Scanf("%d", &baseTenInput)

	baseTwoInput := fmt.Sprintf("%b", baseTenInput)

	total := getLongestStreakOfOnes(baseTwoInput)
	fmt.Println(total)
}

func getLongestStreakOfOnes(input string) int {
	var count int
	streakCounts := []int{}

	a := []rune(input)

	for i := 0; i < len(a); i++ {
		if a[i] == '1' {
			count++
		} else if a[i] == '0' {
			streakCounts = append(streakCounts, count)
			count = 0
		}
	}
	streakCounts = append(streakCounts, count)

	return getLargestStreak(streakCounts)
}

func getLargestStreak(streakCounts []int) int {
	var largest int
	for _, num := range streakCounts {
		// fmt.Println(num)
		if num > largest {
			largest = num
		}
	}
	return largest
}
