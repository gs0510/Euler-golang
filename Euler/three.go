package main

import (
	"fmt"
	"math"
)

func isPrime(num int64) bool {
	var limit int64 = int64(math.Sqrt(float64(num)))
	var i int64
	for i = 2; i<=limit; i++ {
		if num%i == 0 {
			return false
		}
	}
	return true
}



func main() {
	var i int64
	var limit int64 = int64(math.Sqrt(float64(600851475143)))
	for i=limit; i>1; i-- {
		if 600851475143%i == 0 && isPrime(i) {
				fmt.Println(i)
				break
			}
		}
}