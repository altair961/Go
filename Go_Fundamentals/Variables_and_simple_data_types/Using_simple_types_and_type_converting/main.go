package main

import "fmt"

func main() {
	var a string
	a = "foo"

	fmt.Println(a)

	var b int = 99
	fmt.Println(b)

	var c = true
	fmt.Println(c)

	d := 3.1415 //float 64
	fmt.Println(d)

	var e int = int(d)
	fmt.Println(e)

	// var f int8 = 54
	// var g int16 = f
	// fmt.Println(e)
}
