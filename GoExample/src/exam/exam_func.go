package main

import (
	"fmt"
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





