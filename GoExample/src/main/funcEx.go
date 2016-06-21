package main

import (
	"fmt"
	"math"
)

func swap(x, y string) (string, string) {
	return y, x
}

func split(sum int) (x, y int) {
	x = sum * 4 / 9
	y = sum - x
	return 
}

func shortInit() {
	var x, y, z int = 1, 2, 3
	c, python, java := true, false, "no!"
	
	fmt.Println(x,y,z,c,python,java)
}

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

// go에서의 while문
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
/*
	조건문
*/
func sqrt(x float64) string {
	if x < 0 {
		return sqrt(-x) + "i"
	}
	return fmt.Sprint(math.Sqrt(x))
}
func pow(x, n, lim float64) float64 {
	// v는 if안에서만 사용이 가능합니다.
	if v := math.Pow(x, n); v < lim {
		return v
	}
	return lim
}
func pow2(x, n, lim float64) float64 {
	if v := math.Pow(x, n); v < lim {
		return v
	} else {
		fmt.Println("%g >= %g\n", v, lim)
	}
	//cant use v here, though
	return lim
}



