package main

import (

)

func hanoi(n int, a string, b string, c string) string {
		if n == 1 {
			return a +" -> "+c +"\n"
		} else {
			return hanoi(n-1,a,c,b) + a +" -> "+ c + "\n" + hanoi(n-1,b,a,c)
		}
} 

func fibonacci() func() int {
	var sum int = 0
	var a1 int = 0
	var a2 int = 1
	return func() int {
		//sum은 a2다음 숫자
		sum = a1 + a2
		a1 = a2 
		a2 = sum
		return sum
	}
}
