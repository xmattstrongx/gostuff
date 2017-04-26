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
