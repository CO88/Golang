package main

import (
	"math"
)

//Hypot = Hypotenuse's abbreviation
func Hypot(a, b float64) float64 {
	return math.Sqrt(a*a + b*b)
}

func CAbs(x complex128) float64 {
	return Hypot(real(x), imag(x))
}


func Cbrt(x complex128) complex128 {
	
	const limit = 1e-15
	z := complex128(x/2)
	z_pre := complex128(x)

	
	for ( CAbs(z - z_pre) > limit) {
		z_pre = z
		//NewTon's method
		z = z - (z * z * z -x) / (3 * z * z)
	}
	return z
}


