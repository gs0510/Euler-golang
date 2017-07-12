package main

import "fmt"

func main() {
	var num int64 = 1
	for i:=int64(2); i<=20; i++ {
		num = num*i
	}
	var flag bool = false
	for i:=int64(1000); i<=num; i++ {
		flag = false
		for j:=int64(2); j<=20; j++ {
			if i%j != 0 {
				i++;
				flag = true
				break
			}
		}
		if !flag {
			fmt.Println(i)
			break
		}
	}
}