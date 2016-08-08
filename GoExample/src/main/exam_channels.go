package main

import (
	"fmt"
	"time"
)

/*
	ch <- v		// v를 ch로 보냄
	v := <-ch	// ch로부터 값을 받아서 v로 넘김
	
	ch := make(chan int)
	
	기본적으로, 송/수신은 상대편이 준비될 때까지 블록이 됨

*/

var channel = make(chan int, 10)
var str string

func sum(a []int, c chan int) {
	sum := 0
	for _, v := range a {
		sum += v
		//fmt.Println(v)
	}
	c <- sum
}

func chanBuffered() {
	bufferedChan := make(chan int, 2)
	
	bufferedChan <- 1
	bufferedChan <- 2
	/* 
		channel에 버퍼사이즈가 2인데 그 이상을 넣으려고 해서
		Deadlock !!! 
	*/  
	//bufferedChan <- 3
	
	fmt.Println(<-bufferedChan)
	fmt.Println(<-bufferedChan)
}

func chanExam01() {
	str = "hello, world"
	channel <- 1
}

/*
	close(chan) 은 송신측에서만 할 수 있다. 
	닫힌 채널에 데이터를 보내면 패닉이 일어난다.
	또한, 채널은 파일과 다르게 항상 닫을 필요가 없습니다.
	수신측에게 더 이상 보낼 값이 없다고 말해야 할때만 
	행해지면 됩니다. range 루프를 종료시켜야 할 때 처럼
	
	수신측은 v, ok := <-ch 로 ok로 채널이 닫혀 있는지 확인 할 수가 있습니다.
*/
func fibonacciChan(n int, c chan int) {
	x, y := 0, 1
	for i := 0; i < n; i++ {
		c <- x
		x, y = y, x+y
	}
	close(c)
}

/*
	select 구문은 고루틴이 다수의 통신 동작으로부터 
	수행 준비를 기다릴 수 있게 합니다.
*/
func fibonacciSelect(c, quit chan int) {
	x, y := 0, 1
	for{
		select {
			case c <- x :
				x, y = y, x+y
			case <- quit :
				fmt.Println("quit")
				return
		}
	}
}

/*
	select 의 default 케이스는 현재 수행 준비가 
	완료된 케이스가 없을 때 수행됩니다
	
	select {
		case i := <-c :
			// i를 사용
		default :
			// c로부터의 수신은 블록된 상태
	}
*/
func selectDefaultTest() {
	tick := time.Tick(1e8)
	boom := time.After(5e8)
	
	for {
		select {
			case <-tick:
				fmt.Println("tick.")
			case <-boom:
				fmt.Println("BOOM!")
				return
			default:
				fmt.Println("     .")
				time.Sleep(5e7)
		}
	}
}

