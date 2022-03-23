package main

import (
	"codewars/sorting"
	"fmt"
)

func main() {
	// fmt.Println(fundamentals.GetSize(4, 2, 6))    //88 48
	// fmt.Println(fundamentals.GetSize(1, 1, 1))    //6 1
	// fmt.Println(fundamentals.GetSize(1, 2, 1))    //10, 2
	// fmt.Println(fundamentals.GetSize(1, 2, 2))    //16, 4
	// fmt.Println(fundamentals.GetSize(10, 10, 10)) //600, 1000
	// fmt.Println(fundamentals.CountPositivesSumNegatives([]int{1, 2, 3, 4, 5, 6, 7, 8, 9, 10, -11, -12, -13, -14, -15}))
	// fmt.Println(fundamentals.CountPositivesSumNegatives([]int{}))
	// fmt.Println(fundamentals.CountPositivesSumNegatives(nil))
	// fmt.Println(fundamentals.Maps([]int{1, 2, 3}))          //([]int{2, 4, 6}))
	// fmt.Println(fundamentals.Maps([]int{4, 1, 1, 1, 4}))    //([]int{8, 2, 2, 2, 8}))
	// fmt.Println(fundamentals.Maps([]int{2, 2, 2, 2, 2, 2})) //([]int{4, 4, 4, 4, 4, 4}))
	// fmt.Println(fundamentals.Combat(100.0, 50.0)) //50.0
	// fmt.Println(fundamentals.Combat(1.0, 100))    //0.0
	// fmt.Println(fundamentals.SortNumbers([]int{1, 2, 10, 50, 5})) //[]int{1,2,5,10,50}
	// fmt.Println(fundamentals.SortNumbers([]int{}))                //[]int{}
	// fmt.Println(fundamentals.SortNumbers(nil))                    //nil
	arr := []int{1, 39, 2, 9, 7, 54, 11}
	fmt.Println(sorting.Bubble(arr))
	fmt.Println(sorting.RecursiveBubble(arr, len(arr)))
	fmt.Println(sorting.Insertion(arr))
	fmt.Println(sorting.Selection(arr))
}
