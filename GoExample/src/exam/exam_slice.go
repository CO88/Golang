package main

import (
	"fmt"
)

func test_slice() {
	p := []int{2,3,5,7,11,13}
	fmt.Println("p==", p)
	
	for i := 0; i < len(p); i++ {
		fmt.Printf("p[%d] == %d\n",i,p[i])
	}
}

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

func slicing_slices() {
	
	p := []int{2, 3, 5, 7, 11, 13}
	
	fmt.Println("p ==", p)
	fmt.Println("p[1:4] ==", p[1:4])
	
	// missing low index implies 0
	fmt.Println("p[:3] ==", p[:3])
	
	// missing high index implies len(s)
	fmt.Println("p[4:] ==", p[4:])
}

func print_slice(s string, x []int) {
	fmt.Printf("%s len=%d cap=%d %v\n",s, len(x), cap(x), x)	
}

func isEmptySlice_int(x []int) bool {
	
	//fmt.Println(x, len(z), cap(z))
	if len(x) == 0 || x == nil {
		return true
	} else {
		return false
	}
}

func iterates_slice1(x []int) {
	
	//i에는 index가 v에는 value가 들어옴
	// index, value := range slice
	for i, v := range x {
		fmt.Printf("x[%d] = %d\n",i,v)
	}
}

func iterates_slice2() {
	pow := make([]int, 10)
	
	for i := range pow {
		pow[i] = 1 << uint(i)
	}
	
	// _을 이용하여 INDEX값 무시	
	for _, value := range pow {
		fmt.Printf("%d\n", value)
	}
}


