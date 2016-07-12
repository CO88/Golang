package main

import (
	"fmt"
)
/*
	반복문
*/
func for_example01() {
	sum := 0
	for i:=0; i < 10; i++ {
		sum += i
	}
	fmt.Println(sum)
}

/*
 	go에서의 while문
*/
func for_example02() {
	sum := 1
	for sum < 1000 {
		sum += sum
	}
	fmt.Println(sum)
}
func for_example03() {
	for {
		// 무한 루프
	}
}
