package main

import "fmt"

func main() {
	a := 0
	b := &a
	n := &b
	ultimatepointone(&n)
	fmt.Println(a)
}

func ultimatepointone(n **int) {
	**n = 1
}
