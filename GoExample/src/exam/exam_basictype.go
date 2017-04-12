
/*
	go의 기본 자료형
	
	bool 
	string
	int	int8 int16 int32 int64
	uint uint8 uint16 uint32 uint64 uintptr
	
	byte // uint8의 alias
	rune // int32의 alias
		 // 유니코드 코드 포인트 값을 표현
		 
	float32 float64
	complex64 complex128
*/

package main

import (
	"fmt"
	"math/cmplx"
)

var (
	Tobe bool		= false
	MaxInt uint64	= 1<<64 - 1
	zi complex128	= cmplx.Sqrt(-5 + 12i)
)

//struct
type Vertex struct {
	X int
	Y int
}

//struct pointer
func pointerCopy(src Vertex) {
	q := &src
	q.X = 1e9
	fmt.Println("q = ", q.X)
}

//Struct Literals
//구조체 리터럴은 필드와 값을 나열해서 구조체를 새로 할당하는 방법
func exam_Literals() {
	var(
		p = Vertex{1,2}		//has type Vertex
		q = &Vertex{1,2}	//has type *Vertex
		r = Vertex{X: 1}	//Y:0 is implicit
		s = Vertex{}		//X:0 and Y:0
	)
	
	fmt.Println(p, q, r, s)
}
