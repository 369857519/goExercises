package main

import (
	"fmt"
	"unsafe"
)

const c_name1 = "abc"

const (
	Unknown = 0
	Female  = 1
	Male    = 2
)

func main() {
	//常量、枚举、可以使用的函数
	fmt.Println(c_name1)
	fmt.Println(Unknown, Female, Male)
	const (
		a = "abc"
		b = len(a)
		c = unsafe.Sizeof(a)
	)
	fmt.Println(a, b, c)
	//iota
	const (
		ia = iota
		ib
		ic
		id = ""
		ie = iota
	)
	fmt.Println(ia, ib, ic, id, ie)
	const (
		ii = 1 << iota
		ij
		ik
	)
	fmt.Println(ii, ij, ik)
}
