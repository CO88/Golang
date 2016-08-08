package main

import (
	"golang.org/x/tour/tree"
)

func Walk(t *tree.Tree, ch chan int) {
	 
	 
	var preOrder func(t *tree.Tree)
	preOrder = func(t *tree.Tree) {
		if t == nil {
			return 
		}
		preOrder(t.Left)
		ch <- t.Value
		preOrder(t.Right)
	}
	
	preOrder(t)
	close(ch)
}

func Same(t1, t2 *tree.Tree) bool {
	
	ch1 := make(chan int)
	ch2 := make(chan int)
	
	go Walk(t1, ch1)
	go Walk(t2, ch2)
	
	for {
		i1, ok1 := <-ch1
		i2, ok2 := <-ch2
		
		if i1 != i2 || ok1 != ok2 {
			return false
		}
		if (ok1 == false) && (ok2 == false) {
			break
		}
	}
	
	
	return true
}