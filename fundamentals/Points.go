/**
* Write a function that takes such collection
* and counts the points of our team in the championship.
*
* Rules for counting points for each match:
*
* if x > y: 3 points
* if x < y: 0 point
* if x = y: 1 point
*
* Notes:
*
* there are 10 matches in the championship
* 0 <= x <= 4
* 0 <= y <= 4
*
 */

package fundamentals

import "fmt"

func Points(games []string) int {
	for i, scores := range games {
		fmt.Println(i, scores)
	}
	return -1
}
