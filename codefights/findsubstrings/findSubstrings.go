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
