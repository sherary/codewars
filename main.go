package main

import (
	"codewars/fundamentals"
	"fmt"
)

func main() {
	// fmt.Println(fundamentals.GetSize(4, 2, 6))    //88 48
	// fmt.Println(fundamentals.GetSize(1, 1, 1))    //6 1
	// fmt.Println(fundamentals.GetSize(1, 2, 1))    //10, 2
	// fmt.Println(fundamentals.GetSize(1, 2, 2))    //16, 4
	// fmt.Println(fundamentals.GetSize(10, 10, 10)) //600, 1000
	fmt.Println(fundamentals.CountPositivesSumNegatives([]int{1, 2, 3, 4, 5, 6, 7, 8, 9, 10, -11, -12, -13, -14, -15}))
	fmt.Println(fundamentals.CountPositivesSumNegatives([]int{}))
	fmt.Println(fundamentals.CountPositivesSumNegatives(nil))
}
