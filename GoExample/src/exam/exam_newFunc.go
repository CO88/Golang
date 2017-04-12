
package main

import (
	"fmt"
)

func exam_newFunc() {
	
	v := new(Vertex)
	fmt.Println(v)
	
	v.X, v.Y = 11, 9
	fmt.Println("new Vertex:", v)
	//fmt.Println("old Vertex:", Vertex)
}

