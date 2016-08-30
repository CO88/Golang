package main

import (
	"fmt"
)

func hasDupeRune(s string) bool {
	runeSet := map[rune]struct{}{}
	for _, r := range s {
		if _, exist := runeSet[r]; exist {
			return true
		}
		// bool대신 빈 구조체 (struct{}{})를 사용하면 메모리 절약
		// 자료형 : struct{} (아무것도 없는 구조체) 데이터 : {} (아무것도 없음)
		runeSet[r] = struct{}{}
	}
	return false
}