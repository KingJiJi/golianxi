package main

import (
	"fmt"
	"reflect"
)

func main() {
	var n int = 1
	var chs = make([]chan int, n)

	var worker = func(n int, c chan int) {
		for i := 0; i < n; i++ {
			c <- i
		}
		close(c)
	}

	for i := 0; i < n; i++ {
		chs[i] = make(chan int)
		go worker(3+i, chs[i])
	}

	var selectCase = make([]reflect.SelectCase, n)
	for i := 0; i < n; i++ {
		selectCase[i].Dir = reflect.SelectRecv //设置信道是接收
		selectCase[i].Chan = reflect.ValueOf(chs[i])
	}

	numDone := 0
	for numDone < n {
		chosen, recv, recvOk := reflect.Select(selectCase)
		if recvOk {
			fmt.Println(chosen, recv.Int(), recvOk)
			//0 0 true
			//0 1 true
			//0 2 true
		} else {
			numDone++
		}
	}
}
