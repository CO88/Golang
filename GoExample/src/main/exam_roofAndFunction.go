package main

import (
	"fmt"
)

func Sqrt(x float64) float64 {
	
	z := float64(x/2)
	tmp := float64(x)
	// z := 1.0
	
	for (tmp - z > 0.0000000000000001) {
		tmp = z
		//NewTon's method 
		z = z - (z*z -x)/(2*z)
		fmt.Println(" z = ",z)
	}
	
	return z
}

