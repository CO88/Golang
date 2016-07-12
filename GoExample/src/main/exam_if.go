package main

import (
	"fmt"
	"math"
)

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

