// Create a method IsUpperCase to see whether the string is ALL CAPS.
//For example:

// type MyString string
// MyString("c").IsUpperCase() == false
// MyString("C").IsUpperCase() == true
// MyString("hello I AM DONALD").IsUpperCase() == false
// MyString("HELLO I AM DONALD").IsUpperCase() == true
// MyString("ACSKLDFJSgSKLDFJSKLDFJ").IsUpperCase() == false
// MyString("ACSKLDFJSGSKLDFJSKLDFJ").IsUpperCase() == true
// In this Kata, a string is said to be in ALL CAPS
// whenever it does not contain any lowercase letter
// so any string containing no letters at all is trivially considered to be in ALL CAPS.

package fundamentals

import "strings"

type MyString string

func (s MyString) IsUpperCase() bool {
	// Your code here!
	return s == MyString(strings.ToUpper(string(s)))
}
