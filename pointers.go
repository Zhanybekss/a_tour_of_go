package main

import "fmt"

func main() {
	i := 10
	p := &i
	fmt.Println(p)
	fmt.Println(*p)

	j := 22
	p = &j         // point to j
	*p = *p / 2    // divide j through the pointer
	fmt.Println(j) // see the new value of j
}
