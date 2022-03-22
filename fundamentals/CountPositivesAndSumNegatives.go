/**
Given an array of integers.

Return an array, where the first element is the count of positives numbers
and the second element is sum of negative numbers.
0 is neither positive nor negative.

If the input is an empty array or is null, return an empty array.


*/

package fundamentals

func CountPositivesSumNegatives(numbers []int) []int {
	var res []int
	count, sum := 0, 0
	if len(numbers) == 0 || numbers == nil {
		return res
	}
	for _, num := range numbers {
		if num > 0 {
			count++
		}

		if num < 0 {
			sum += num
		}
	}
	res = append(res, count, sum)
	return res
}
