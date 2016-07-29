package main

import (
	"fmt"
	"math"
	"math/cmplx" 
	"golang.org/x/tour/pic"
	"golang.org/x/tour/wc"
)

var x, y, z int
var c, python, java bool

func main(){
	fmt.Println(hanoi(3,"a","b","c"))
	//use_slice()
	//iterate_slice()
	//a, b := swap("hello","world")
	//fmt.Println(a,b)
	//fmt.Println(split(17))
	//fmt.Println(x, y, z, c, python, java)
	//shortInit()
	
	//2016-06-21
	fmt.Println(sqrt(2), sqrt(-4))
	fmt.Println(pow(3, 2, 10), pow(3, 3, 20))
	
	//2016-07-12
	fmt.Println("Newton's method : ",  Sqrt(-25))
	fmt.Println("math.Sqrt method : ", math.Sqrt(25))
	
	//2016-07-18
	fmt.Println(Vertex{1,2})
	
	v := Vertex{1,2}				//X = 1 , Y = 2
	v.X = 4
	fmt.Println(v.X, " : ", v.Y)	//X = 4 , Y = 2
	pointerCopy(v)
	exam_Literals()
	
	exam_newFunc()
	test_slice()
	slicing_slices()
	
	a := make([]int, 5)
	print_slice("a", a)
	b := make([]int, 0, 5)
	print_slice("b", b)
	c := b[:2]
	print_slice("c", c)
	d := c[2:5]
	print_slice("d", d)
	
	//2016-07-19
	fmt.Println("Is this slice (a) empty?? :",isEmptySlice_int(a))
	fmt.Println("Is this slice (b) empty?? :",isEmptySlice_int(b))
	fmt.Println("Is this slice (c) empty?? :",isEmptySlice_int(c))
	fmt.Println("Is this slice (d) empty?? :",isEmptySlice_int(d))
	
	//2016-07-20
	iterates_slice1(a)
	iterates_slice2()
	
	pic.Show(Pic)
	
	//2016-07-22
	var location_map = createMap()
	fmt.Println(location_map)
	
	//2016-07-25
	literals_Test()
	mutating_map()
	
	functionValue()
	
	//2016-07-26
	wc.Test(WordCount)
	pos, neg := adder(), adder()
	for i := 0; i < 10; i++ {
		fmt.Println(
			pos(i),
			neg(-2*i),
		)
	}
	
	//A Tour of GO #44
	//연습 : 피보나치 클로져
	f := fibonacci();
	for i := 0; i < 10; i++ {
		fmt.Println(f())
	}
	//함수의 주소값이 출력됨
	//아래 2가지의 출력문은 값이 동일함
	fmt.Println(fibonacci())
	fmt.Println(f)
	
	//2016-07-27
	switchExample01()
	switchExample02()
	switchExample03()
	
	fmt.Println(Cbrt(2+2i))
	fmt.Println(cmplx.Pow(2+2i, 1/3.0))
	
	fmt.Println(Cbrt(8))
	fmt.Println(cmplx.Pow(8, 1/3.0))
	
	//2017-07-29 (#50 ~
	var ai Abser
	v3 := &Vertex3D{3,4,5}
	myf := MyFloat(-math.Sqrt2)
	
	v3.scale(3)
	fmt.Println(v3.X,",",v3.Y,",",v3.Z)
	
	ai = myf // a MyFloat implements Abser
	
	// a Vertex3D, does NOT implements Abser
	//ai = &v3 
	
	fmt.Println(ai.Abs())
	



}