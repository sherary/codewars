package main

import (
	"codewars/fundamentals"
	"fmt"
)

func main() {
	fmt.Println(fundamentals.GetSize(4, 2, 6))    //88 48
	fmt.Println(fundamentals.GetSize(1, 1, 1))    //6 1
	fmt.Println(fundamentals.GetSize(1, 2, 1))    //10, 2
	fmt.Println(fundamentals.GetSize(1, 2, 2))    //16, 4
	fmt.Println(fundamentals.GetSize(10, 10, 10)) //600, 1000
}
