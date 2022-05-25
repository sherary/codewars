/**
* Your function takes two arguments:
*
* current father's age (years)
* current age of his son (years)
* Сalculate how many years ago the father was twice as old as his son (or in how many years he will be twice as old).
 */

package fundamentals

import "math"

func TwiceAsOld(dadYearsOld, sonYearsOld int) int {
	// Implement me
	return int(math.Abs(float64(dadYearsOld - sonYearsOld*2)))
}
