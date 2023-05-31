package main

import "fmt"

func main() {
	array := [5]int{1, 2, 3, 4, 5}
	fmt.Println(array)
	a := array[0:2]
	b := array[2:5]
	fmt.Println(a, b)
	a[0] = 99
	fmt.Println(a)
	fmt.Println(array)
}
