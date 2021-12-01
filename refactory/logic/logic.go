package main

import (
	"fmt"
	"math"
	"strconv"
)

func main() {
	// kata := "malam"

	// pal := Palindrome(kata)

	// if pal {
	// 	fmt.Println("Yes")
	// } else {
	// 	fmt.Println("no")
	// }

	// =======================
	// LeapYear(1900, 2000)

	// ========================
	// rev := ReverseWord("I am A Great human")
	// fmt.Println(rev)

	// ==========================
	// arr := []int{15, 1, 3}
	// num := NearestFibonacci(arr)
	// fmt.Println(num)

	//========================
	// FB := FizzBuzz(15)
	// fmt.Println(FB)
}

func Palindrome(kata string) bool {
	if len(kata) <= 1 {
		return true
	}
	if kata[0] == kata[len(kata)-1] {
		return Palindrome(kata[1 : len(kata)-1])
	} else {
		return false
	}
}

func LeapYear(a, b int) {
	for i := a; i <= b; i += 4 {
		fmt.Print(i)
		if i >= b-1 {
			break
		}
		fmt.Print(",")
	}

}

func ReverseWord(kata string) string {
	var reverse string
	for _, v := range kata {
		reverse = string(v) + reverse
	}
	return reverse
}

func NearestFibonacci(arr []int) interface{} {
	if len(arr) == 0 {
		fmt.Println(0)
		return true
	}

	num := 0
	for _, v := range arr {
		num = num + v
	}

	first := 0
	second := 1
	third := first + second

	for third < num {
		first = second
		second = third
		third = first + second
	}

	if math.Abs(float64(third-num)) >= math.Abs(float64(second-num)) {
		return second
	} else {
		return third
	}
}

func FizzBuzz(a int) []string {
	var result []string

	for i := 1; i <= a; i++ {
		if i%3 == 0 && i%5 == 0 {
			result = append(result, "FizzBuzz")
		} else if i%5 == 0 {
			result = append(result, "Fizz")
		} else if i%3 == 0 {
			result = append(result, "Buzz")
		} else {
			result = append(result, strconv.Itoa(i))
		}
	}
	return result
}
