package main

import (
	"fmt"
)

func Sqrt(x float64) float64 {
	
	const threshold = 1e-15
	
	z := float64(x/2)
	tmp := float64(x)

	//음수 일 때 처리
	if x < 0 {
		return x;
	}
	
	for (tmp - z > threshold) {
		tmp = z
		//NewTon's method 
		z = z - (z*z -x)/(2*z)
		fmt.Println(" z = ",z)
	}
	
	return z
}

