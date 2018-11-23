package main

import (
	"fmt"
)

func main() {
	x := 15
	a := &x // point to memory address

	fmt.Println(a)
	fmt.Println(*a) // read through the memory address
	*a = 5
	fmt.Println(x)
}
