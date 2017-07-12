package main

import "fmt"

func fibo(num int64) int64 {
	var first int64 = 0
	var second int64 = 1
	var sum int64
	sum = 0
	for second < num {
		first , second = second, first + second
		if second%2 == 0 {
			sum += second
		}
	}
	return sum
}

func main() {
	fmt.Println(fibo(4000000))
}