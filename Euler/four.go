package main

import "fmt"

func reverseInt(num int64) int64 {
	new_int := int64(0)
	for num > 0 {
		new_int = new_int*10 + num%10
		num = num/10
	}
	return new_int
}

func isPalindrome(num int64) bool {
	reverse := reverseInt(num)
	if reverse == num {
		return true
	}
	return false
}

func main() {
	for i := 999; i>=100; i-- {
		for j := i; j>=100; j-- {
			if isPalindrome(int64(i*j)) {
				fmt.Println(i*j)
				break
			}
		}
	}
}