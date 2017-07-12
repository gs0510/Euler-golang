package main

import "fmt"

func isDivisible(num int) int {
	if num%3 == 0 {
		return num
	} else if num%5 == 0 {
		return num
	} else {
		return 0
	}
}

func main() {
	sum := 0
	for i:=0; i<1000; i++ {
		sum += isDivisible(i)
	}
	fmt.Println(sum)
}