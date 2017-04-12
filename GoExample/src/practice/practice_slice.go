package main

import (

)

func Pic(dx, dy int) [][]uint8 {
	//dy개 만큼의 길이를 가지는 슬라이스를 리턴
	//dx 개의 8비트 부호없는 8비트 정수 타입을 가지는 슬라이스
	result := make([][]uint8, dy)
	
	for i := 0; i < dy; i++ {
		result[i] = make([]uint8, dx)
		for j := 0; j < dx; j++ {
			result[i][j] = uint8((i+j)/2)
		}
	}
	return result
}
