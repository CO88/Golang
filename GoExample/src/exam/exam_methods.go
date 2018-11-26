package main

import (
	"math"
)

// Interface
type Abser interface {
	Abs() float64
}

type MyFloat float64

type Vertex3D struct{
	X, Y, Z float64
}

// (v *Vertex3D)는 메소드 리시버(method receiver)
func (v *Vertex3D) hypot_3d() float64 {
	return math.Sqrt(v.X * v.X + v.Y * v.Y + v.Z * v.Z)
}

/*
	 func (v Vertex3D) scale(f float64)로 
 	메소드 리시버를 포인터 리시버로 사용하지 않으면
 	원래의 값(Vertex3D)을 바꿀 수 없다.
 	출력의 경우는 Vertex3D를 읽기만 하기 때문에, 
 	상관이 없음
*/
func (v *Vertex3D) scale(f float64) {
	v.X = v.X * f
	v.Y = v.Y * f
	v.Z = v.Z * f
}

//메소드는 구조체뿐만 아니라 아무 타입(Type)에나 붙일 수 있음
func (f MyFloat) Abs() float64 {
	if f < 0 {
		return float64(-f)
	}
	return float64(f)
}


