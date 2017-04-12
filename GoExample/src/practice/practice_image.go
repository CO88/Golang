package main	

import (
	"image"
	"image/color"
)

type Image struct {
	w int
	h int
}

/*
	method receiver를 pointer receiver로 할 경우 
	image interface의 method인 Bounds를 구현하지 않았다는 
	오류가 발생
	==> 
		m := Image{150,30}
		pic.ShowImage(&m)
	 이 부분에서 ShowImage() 함수의 파라미터를 구조체의
	 주소값이 아닌 값을 넣어 버리면 메소드도 복사를 해야
	 하기 때문에 method receiver를 사용해야함
	 
	 아니면, ShowImage() 함수의 파라미터의 값을 image 인터페이스
	 를 암시적으로 구현한 구조체의 주소값인 &m을 넣으면
	 pointer receiver, mehtod receiver둘 다 사용 가능
*/ 
func (i *Image) Bounds() image.Rectangle {
	return image.Rect(0, 0, i.w, i.h)
}

func (i *Image) ColorModel() color.Model {
	return color.RGBA64Model
}

func (i *Image) At(x, y int) color.Color {
	
	pic := func(x, y int) uint8 {
		return uint8((x+y)/2)
	}
	
	v := pic(x,y)
	
	return color.RGBA{v, v, 255, 255}
}
