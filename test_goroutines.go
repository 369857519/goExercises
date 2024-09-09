package main

import (
	"fmt"
	"time"
)

func say(s string) {
	for i := 0; i < 5; i++ {
		time.Sleep(100 * time.Millisecond)
		fmt.Println(s)
	}
}

//ch <- v 把值发到通道 ch
//v:= <-ch 从ch接收数据，并赋值给 v

func sum(s []int, c chan int) {
	sum := 0
	for _, v := range s {
		sum += v
	}
	c <- sum //把 sum 发送到通道c
}

func fibonacci(n int, c chan int) {
	x, y := 0, 1
	for i := 0; i < n; i++ {
		c <- x
		x, y = y, x+y
	}
	close(c)
	fmt.Println("range finish")
}

func fibonacciSelect(c, quit chan int) {
	x, y := 0, 1
	for {
		select {
		case c <- x:
			x, y = y, x+y
		case <-quit:
			fmt.Println("quit")
			return

		}
	}
}

func main() {
	go say("world")
	say("hello")

	s := []int{1, 2, 6, 7, 3, 2, 3, 4, 6, 8}
	c := make(chan int)
	go sum(s[:len(s)/2], c)
	go sum(s[len(s)/2:], c)
	x, y := <-c, <-c
	fmt.Println(x, y, x+y)
	//通道缓冲区
	//带缓冲区的通道允许发送端的数据发送和接收端的数据获取处于异步状态，就是说发送端发送的数据可以放在缓冲区里面，可以等待接收端去获取数据，而不是立刻需要接收端去获取数据。
	ch := make(chan int, 2)
	ch <- 1
	ch <- 2
	fmt.Println(<-ch, <-ch)
	//缓冲区遍历
	//通过range官架子来遍历获取到的数据
	cb := make(chan int, 10)
	fmt.Println("range")
	go fibonacci(cap(cb), cb)
	// range函数遍历每个从通道接收到的参数
	for i := range cb {
		fmt.Println(i)
	}
	//select 语句，一个goroutin实现扽带多个通信操作
	cSelect := make(chan int)
	quit := make(chan int)
	go func() {
		for i := 0; i < 10; i++ {
			fmt.Println(<-cSelect)
		}
		quit <- 0
	}()
	fibonacciSelect(cSelect, quit)

}
