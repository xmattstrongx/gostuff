package main

import "fmt"

func main() {
	var leftAnagram, rightAnagram string
	fmt.Scanf("%s\n%s", &leftAnagram, &rightAnagram)

	leftHash := getCharacterMap(leftAnagram)
	rightHash := getCharacterMap(rightAnagram)

	anagramDifference := hashCompare(leftHash, rightHash)
	fmt.Println(anagramDifference)
}

func getCharacterMap(anagram string) map[string]int {
	characterCounter := make(map[string]int)
	for _, v := range anagram {
		_, ok := characterCounter[string(v)]
		if !ok {
			characterCounter[string(v)] = 1
		} else {
			characterCounter[string(v)]++
		}
	}

	return characterCounter
}

func hashCompare(leftHash map[string]int, rightHash map[string]int) int {
	var anagramDifference int
	if len(leftHash) >= len(rightHash) {
		anagramDifference = calculateAnagramDifference(leftHash, rightHash)
	} else {
		anagramDifference = calculateAnagramDifference(rightHash, leftHash)
	}

	return anagramDifference
}

func calculateAnagramDifference(greaterHash map[string]int, lessHash map[string]int) int {
	var anagramDifference int
	for key, value := range greaterHash {
		leftValue, ok := lessHash[key]
		if !ok {
			anagramDifference += value
		} else if leftValue < value {
			anagramDifference += value - leftValue
		} else if leftValue > value {
			anagramDifference += leftValue - value
		}
	}

	for key, value := range lessHash {
		_, ok := greaterHash[key]
		if !ok {
			anagramDifference += value
		}
	}

	return anagramDifference
}

// fcczanyyyx
// jthujrpdoqbbwhpmeoke

// f:1 y:3 c:2 a:1 n:1 z:1 = 9

// w:2 w:1 = 1
// r:2 r:1 = 1
// g:1 g:1
// e:2
// j:2
// h:2
// u:1
// p:2
// d:1
// o:2
// b:2
// s:1 s:1
// t:1
// l:1 l:1
// m:3 m:2 = 1
// v:1 v:1
// q:1
// i:1 i:1
// k:1
// x:1 x:2 = 1

// 9 + 4 + 17
