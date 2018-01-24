package main

import (
	"fmt"
	"reflect"
)

func main() {
	var chs = make(chan int)
	var worker = func(c chan int) {
		for i := 0; i < 5; i++ {
			c <- i
		}
		close(c)
	}

	go worker(chs)

	var selectCase = make([]reflect.SelectCase, 1)
	selectCase[0].Dir = reflect.SelectRecv //设置信道是接收
	selectCase[0].Chan = reflect.ValueOf(chs)

	numDone := 0
	for numDone < 1 {
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
