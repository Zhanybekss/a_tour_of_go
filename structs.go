package main

import "fmt"

type Laptop struct {
	model       string
	memory_size int
}

func main() {
	laptop := Laptop{
		model:       "Asus",
		memory_size: 256,
	}
	fmt.Println(laptop)
}
