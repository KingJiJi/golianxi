package main

import (
	"fmt"
)

func main() {
	var c = make(chan int, 1)
	go func(c chan int) {
		for i := 0; i < 10; i++ {
			c <- i
		}
		close(c)
	}(c)

L:
	for {
		select {
		case val, ok := <-c:
			if ok {
				fmt.Println("接收: ", val)
			} else {
				fmt.Println("信道关闭了")
				break L
			}
		}
	}
	//接收:  0
	//接收:  1
	//接收:  2
	//接收:  3
	//接收:  4
	//接收:  5
	//接收:  6
	//接收:  7
	//接收:  8
	//接收:  9
	//信道关闭了
}
