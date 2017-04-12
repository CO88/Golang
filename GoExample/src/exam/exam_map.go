package main

import (
	"fmt"
)

type Vertex_Location struct {
	Lat, Long float64
}

func createMap() map[string]Vertex_Location {
	
	var res = make(map[string]Vertex_Location)
	
	return res
}

func literals_Test() {
	var m = map[string]Vertex_Location {
		/*
		"Bell Labs" : Vertex_Location {
			40.68433, -74.39967,
		},
		"Google" : Vertex_Location {
			37.42202, -122.08408,
		},
		*/
		//가장 상위의 타입이 타입명이라면 리터럴에서 타입명을 생략해도 됨
		"Bell Labs" :	{40.68433, -74.39967},
		"Google" :		{37.42202, -122.08408},
	}
	
	fmt.Println(m["Bell Labs"])
	fmt.Println(m["Google"])
}

func mutating_map() {
	m := make(map[string]int)
	
	m["Answer"] = 42
	fmt.Println("The value:", m["Answer"])
	
	
	m["Answer"] = 48
	fmt.Println("The value:", m["Answer"])
	
	delete(m, "Answer")
	fmt.Println("The value:", m["Answer"])
	
	//test value
	//존재하지 않는 key의 반환 값은 타입에 맞는 zero value이다.
	//그래서 value에 0(Integer)을 넣어보고 키의 존재 여부 체크
	m["Answer"] = 0
	v, ok := m["Answer"]
	fmt.Println("The value:", v, "Present?", ok)
	
	delete(m, "Answer")
	v, ok = m["Answer"]
	fmt.Println("The value:", v, "Present?", ok)
}