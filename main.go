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
	// fmt.Println(fundamentals.IsPalindrome("a"))     //true
	// fmt.Println(fundamentals.IsPalindrome("aba"))   //true
	// fmt.Println(fundamentals.IsPalindrome("Abba"))  //true
	// fmt.Println(fundamentals.IsPalindrome("hello")) //false
	// arr := []int{1, 39, 2, 9, 7, 54, 11}
	// fmt.Println(sorting.Bubble(arr))
	// fmt.Println(sorting.RecursiveBubble(arr, len(arr)))
	// fmt.Println(sorting.Insertion(arr))
	// fmt.Println(sorting.Selection(arr))
	// fmt.Println(sorting.CountSort(arr))

	// fmt.Println(fundamentals.MyString("a").IsUpperCase())                                                                     //).To(Equal(false))
	// fmt.Println(fundamentals.MyString("abcdefghijklmnopqrstuvwxyz").IsUpperCase())                                            //).To(Equal(false))
	// fmt.Println(fundamentals.MyString("ABCDEFGHIJKLMNOPQRSTUVWXYz").IsUpperCase())                                            //).To(Equal(false))
	// fmt.Println(fundamentals.MyString("false").IsUpperCase())                                                                 //).To(Equal(false))
	// fmt.Println(fundamentals.MyString("true").IsUpperCase())                                                                  //).To(Equal(false))
	// fmt.Println(fundamentals.MyString("False").IsUpperCase())                                                                 //).To(Equal(false))
	// fmt.Println(fundamentals.MyString("True").IsUpperCase())                                                                  //).To(Equal(false))
	// fmt.Println(fundamentals.MyString("WHAT DOES THE FOX SAY").IsUpperCase())                                                 //).To(Equal(true))
	// fmt.Println(fundamentals.MyString("HTML CSS JAVASCRIPT PYTHON C PERL LISP JAVA GO RUBY NODEJS RUST SCALA").IsUpperCase()) // ).To(Equal(true))

	fmt.Println(fundamentals.AbbrevName("Sam Harris"))     //To(Equal("S.H"))
	fmt.Println(fundamentals.AbbrevName("Patrick Feenan")) //To(Equal("P.F"))
	fmt.Println(fundamentals.AbbrevName("Evan Cole"))      //To(Equal("E.C"))
	fmt.Println(fundamentals.AbbrevName("P Favuzzi"))      //To(Equal("P.F"))
	fmt.Println(fundamentals.AbbrevName("David Mendieta")) //To(Equal("D.M"))
}
