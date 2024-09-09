package main

import "fmt"

func main() {
	var pow = []int{1, 2, 4, 8, 16, 32, 64, 128}
	// 遍历 pow 切片，i 是索引，v 是值
	for i, v := range pow {
		// 打印 2 的 i 次方等于 v
		fmt.Printf("2**%d = %d\n", i, v)
	}

	ch := make(chan int, 2)
	ch <- 1
	ch <- 2
	close(ch)

	for v := range ch {
		fmt.Println(v)
	}
}
