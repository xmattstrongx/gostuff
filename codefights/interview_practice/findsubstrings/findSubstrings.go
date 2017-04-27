package findsubstrings

import (
	"fmt"
	"strings"
)

func findSubstrings(words []string, parts []string) []string {
	if len(parts) == 0 {
		fmt.Println(words)
		return words
	} else if len(words) == 0 {
		return []string{}
	}

	var firstIndex int
	for i, word := range words {
		longestPart := ""
		firstIndex = -1
		for _, part := range parts {
			if strings.Contains(word, part) {
				currentIndex := strings.Index(word, part)
				if len(longestPart) == 0 {
					firstIndex = currentIndex
				}
				if len(longestPart) < len(part) {
					longestPart = part
					firstIndex = currentIndex
				}
				if len(longestPart) == len(part) {
					if currentIndex < firstIndex {
						firstIndex = currentIndex
					}
				}

			}
		}
		if longestPart != "" {
			newWord := word[:firstIndex+len(longestPart)] + "]" + word[firstIndex+len(longestPart):]
			newWord = newWord[:firstIndex] + "[" + newWord[firstIndex:]
			words[i] = newWord
		}
	}

	return words
}

// You have two arrays of strings, words and parts. Return an array that contains the strings from words, modified so that any occurrences of the substrings from parts are surrounded by square brackets [], following these guidelines:

// If several parts substrings occur in one string in words, choose the longest one. If there is still more than one such part, then choose the one that appears first in the string.

// Example

// For words = ["Apple", "Melon", "Orange", "Watermelon"] and parts = ["a", "mel", "lon", "el", "An"], the output should be
// findSubstrings(words, parts) = ["Apple", "Me[lon]", "Or[a]nge", "Water[mel]on"].

// While "Watermelon" contains three substrings from the parts array, "a", "mel", and "lon", "mel" is the longest substring that appears first in the string.

// Input/Output

// [time limit] 4000ms (go)
// [input] array.string words

// An array of strings consisting only of uppercase and lowercase English letters.

// Guaranteed constraints:
// 0 ≤ words.length ≤ 5000,
// 1 ≤ words[i].length ≤ 30.

// [input] array.string parts

// An array of strings consisting only of uppercase and lower English letters. Each string is no more than 5 characters in length.

// Guaranteed constraints:
// 0 ≤ parts.length ≤ 5000,
// 1 ≤ parts[i].length ≤ 5,
// parts[i] ≠ parts[j].

// [output] array.string
