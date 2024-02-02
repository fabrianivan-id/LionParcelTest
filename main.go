package main

import (
	"fmt"
)

func isPalindrome(str string) string {
	for i := 0; i < len(str)/2; i++ {
		if str[i] != str[len(str)-1-i] {
			return "False"
		}
	}
	return "True"
}

func max(arr []int) int {
	max := arr[0]
	for _, v := range arr {
		if v > max {
			max = v
		}
	}
	return max
}

func printTriangle(n int) {
	for i := 1; i <= n; i++ {
		for j := 1; j <= i; j++ {
			fmt.Print("*")
		}
		fmt.Println()
	}
}

func factorial(n int) int {
	if n == 0 {
		return 1
	}
	return n * factorial(n-1)
}

func main() {
	fmt.Println(isPalindrome("radar"))
	fmt.Println(isPalindrome("hello"))

	fmt.Println(max([]int{3, 5, 1, 9, 2}))

	printTriangle(5)

	fmt.Println(factorial(5))
}
