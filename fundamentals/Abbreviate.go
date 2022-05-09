/**
* Write a function to convert a name into initials.
* This kata strictly takes two words with one space in between them.
*
* The output should be two capital letters with a dot separating them.
*
* It should look like this:
*
* Sam Harris => S.H
*
* patrick feeney => P.F
 */

package fundamentals

import "strings"

func AbbrevName(name string) string {
	str := strings.Split(name, " ")
	return strings.ToUpper(str[0][0:1] + "." + str[1][0:1])
}
