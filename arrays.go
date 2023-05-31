package main

import "fmt"

func main() {
	var a [3]string
	a[0] = "hello"
	a[1] = "for"
	a[2] = "world"
	fmt.Println(a[1], a[2])
	fmt.Println(a)

	b := [6]int{1, 2, 3, 4, 5, 6}
	fmt.Println(b)
}
