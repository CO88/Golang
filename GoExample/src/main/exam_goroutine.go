package main

import (
	"fmt"
	"time"
)

func say(s string) {
	
	//share
	var j int = 0;
	
	for i := 0; i < 5; i++ {
		time.Sleep(100 * time.Millisecond)
		j++
		fmt.Println(s,"  ",j)
	}
}

