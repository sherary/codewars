/**

Write a function that checks if a given string (case insensitive) is a palindrome.

*/
package fundamentals

import "strings"

func IsPalindrome(str string) bool {
	if len(str) == 1 {
		return true
	}
	var length = len(str)
	var reversetdString string = ""

	for i := length - 1; i >= 0; i-- {
		reversetdString = reversetdString + string(str[i])
	}

	if strings.ToLower(str) == strings.ToLower(reversetdString) {
		return true
	}

	return false
}
