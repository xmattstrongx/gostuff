package checkpalindrome

func checkPalindrome(inputString string) bool {
	if len(inputString) == 1 {
		return true
	}
	if inputString[0] != inputString[len(inputString)-1] {
		return false
	}
	return checkPalindrome(inputString[1 : len(inputString)-1])
}

// Given the string, check if it is a palindrome.

// Example

// For inputString = "aabaa", the output should be
// checkPalindrome(inputString) = true;
// For inputString = "abac", the output should be
// checkPalindrome(inputString) = false;
// For inputString = "a", the output should be
// checkPalindrome(inputString) = true.
// Input/Output

// [time limit] 4000ms (go)
// [input] string inputString

// A non-empty string consisting of lowercase characters.

// Guaranteed constraints:
// 1 ≤ inputString.length ≤ 10.

// [output] boolean

// true if inputString is a palindrome, false otherwise.
