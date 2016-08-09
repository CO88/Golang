package main

import (
	"golang.org/x/tour/tree"
)
func InOrder(t *tree.Tree, ch chan int) {
	if t == nil { return }
	
	InOrder(t.Left, ch)
	ch <- t.Value
	InOrder(t.Right, ch)
}

func Walk(t *tree.Tree, ch chan int) {
	
	InOrder(t, ch)
	close(ch)
}

func Same(t1, t2 *tree.Tree) bool {
	
	ch1 := make(chan int)
	ch2 := make(chan int)
	
	go Walk(t1, ch1)
	go Walk(t2, ch2)
	
	for {
		val1, ok1 := <-ch1
		val2, ok2 := <-ch2
		
		if val1 != val2 || ok1 != ok2 {
			return false
		}
		if (ok1 == false) && (ok2 == false) {
			break
		}
	}
	return true
}