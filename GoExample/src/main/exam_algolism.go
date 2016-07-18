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