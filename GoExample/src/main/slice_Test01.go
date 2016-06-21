package main

import (
	"fmt"
)

func use_slice(){
	a := []int{1,2,3,4,5}
	
	fmt.Println(len(a), cap(a))
	
	a = append(a, 6, 7)
	
	fmt.Println(len(a), cap(a))
	
}

func iterate_slice(){
	a := []int{1,2,3,4,5}
	
	a = append(a,6,7)
	
	for i:=0; i<len(a); i++ {
		fmt.Println(a[i])
	}
}


