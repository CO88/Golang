package main

import (
	"fmt"
	"time"
)
/*
	type error interface {
		Error() string
	}
 	fmt 패키지의 다양한 출력 루틴들은 error의 출력을 
 	요청 받았을 때 자동으로 이 메소드를 호출합니다.
*/

type MyError struct {
	when time.Time
	what string
}

func (e *MyError) Error() string {
	return fmt.Sprintf("at %v, %s", e.when, e.what)
}

func run() error {
	return &MyError { 
			time.Now(),
			"it didn't work",
			}
}


