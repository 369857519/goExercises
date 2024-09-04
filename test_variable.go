package main

import "fmt"

func main() {
	//变量声明
	fmt.Println("Hello world")
	var a string = "Runoob"

	fmt.Println(a)

	var b, c int = 1, 2
	fmt.Println(b, c)

	var d int
	fmt.Println(d)
	//nil的几种类型
	var e *int
	var f []int
	var g map[string]int
	var h chan int
	var i func(string) int
	var j error
	fmt.Println(e, f, g, h, i, j)
	fmt.Println(e == nil, f == nil, h == nil, i == nil, j == nil)
	//默认值
	var ii int
	var ff float64
	var bb bool
	var ss string
	fmt.Printf("%v %v %v %q\n", ii, ff, bb, ss)
	//声明+默认值,仅能用于函数体中
	dd := 1
	fmt.Println(dd)
	//多变量声明
	var x, y int
	var (
		g_a int
		g_b int
	)
	var ae, af = 123, "hello"
	fmt.Println(x, y, g_a, g_b, ae, af)
	//引用类型，改变引用类型值可能会导致不同变量指向一个VALUE存储，从而互相影响
	fmt.Println(&x)

}
