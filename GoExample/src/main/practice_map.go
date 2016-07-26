package main

import (
	"strings"
)

func WordCount(s string) map[string]int {
	
	//	strings.Fields(string)은 문자열을 공백을 구분자로
	//	단어를 뽑아내어 슬라이스 ( []string) 으로 값을 내보낸다
	var res = strings.Fields(s)
	var count_map = make(map[string]int)
	var v int
	var ok bool

	for _, value := range res {
		
		v, ok = count_map[value]
		if ok == false {
			count_map[value] = 1
		} else {
			count_map[value] = v + 1
		}
	}
	return count_map
}

