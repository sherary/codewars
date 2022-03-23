/**
Given an array of integers, return a new array with each value doubled.

For example:

[1, 2, 3] --> [2, 4, 6]
*/

package fundamentals

// func Maps(x []int) []int {
// 	var res []int

// 	for _, nums := range x {
// 		res = append(res, nums*2)
// 	}

// 	return res
// }

//alternatives
func Maps(x []int) []int {
	result := make([]int, len(x))
	for i, v := range x {
		result[i] = v * 2
	}
	return result
}
