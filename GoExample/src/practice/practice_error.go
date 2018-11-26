package main

import (
	"fmt"
)

type ErrNegativeSqrt float64

func (e ErrNegativeSqrt) Error() string {
	return fmt.Sprintf("cannot Sqrt negative number: %g",e)
	//return fmt.Print(e)
}

func Sqrt(x float64) (float64, error) {
	
	const threshold = 1e-15
	
	if x < 0 {
		return 0, ErrNegativeSqrt(x)
	}
	
	z := float64(1)
	z_pre := float64(x)
	
	for {
		z = z - (z*z - x) / (2*z)
		
		if ((z_pre - z) < threshold) {
			return z, nil
		} else {
			z_pre = z	
		}
	}
}

