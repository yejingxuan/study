package main

import "fmt"

func main() {
	fmt.Println("Hello, World!")
	dataType()
}

func dataType() {
	//int没有初始化就为零值
	var i int
	//f没有初始化就为零值
	var f float64
	//bool零值为 false
	var b bool
	//string没有初始化为""
	var s string

	fmt.Printf("%v %v %v %q", i, f, b, s)
}
