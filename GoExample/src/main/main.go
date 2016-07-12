package main

import (
	"fmt"
	"math"
)

func hanoi(n int, a string, b string, c string) string {
		if n == 1 {
			return a +" -> "+c +"\n"
		} else {
			return hanoi(n-1,a,c,b) + a +" -> "+ c + "\n" + hanoi(n-1,b,a,c)
		}
} 

var x, y, z int
var c, python, java bool

func main(){
	//fmt.Println(hanoi(3,"a","b","c"))
	//use_slice()
	//iterate_slice()
	//a, b := swap("hello","world")
	//fmt.Println(a,b)
	//fmt.Println(split(17))
	//fmt.Println(x, y, z, c, python, java)
	//shortInit()
	
	//2016-06-21
	fmt.Println(sqrt(2), sqrt(-4))
	fmt.Println(pow(3, 2, 10), pow(3, 3, 20))
	
	//2016-07-12
	
	fmt.Println("Newton's method : ",  Sqrt(20))
	fmt.Println("math.Sqrt method : ", math.Sqrt(20))
}