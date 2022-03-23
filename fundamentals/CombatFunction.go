/**
Create a combat function that takes the player's current health
and the amount of damage recieved,
and returns the player's new health.
Health can't be less than 0.
*/

package fundamentals

import "math"

// func Combat(health, damage float64) float64 {
// 	var result = health - damage
// 	if result < 0 {
// 		return 0
// 	}
// 	return result
// }

// alternatives
func Combat(health, damage float64) float64 {
	return math.Max(health-damage, 0.0)
}
