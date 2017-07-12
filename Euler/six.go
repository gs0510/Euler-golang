package main

import "fmt"

func sumSquared(num int64) int64 {
	temp := num*(num + 1)/int64(2)
	return temp*temp
}

func squaredSum(num int64) int64 {
	var sum int64 = 0
	for i:=1; i<=100; i++ {
		sum += int64(i*i)
	}
	return sum
}

func main() {
	diff := sumSquared(100) - squaredSum(100)
	fmt.Println(diff)
}