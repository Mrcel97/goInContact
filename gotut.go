package main

import (
	"fmt"
	"reflect"
)

func add(x, y float64) float64 {
	return x + y
}

func multiple(a, b string) (string, string) {
	return a, b
}

func main() {
	num1, num2 := 5.6, 9.5
	w1, w2 := "Hey", "there"

	fmt.Println(add(num1, num2))
	fmt.Println(multiple(w1, w2))

	// Cast:
	var a int = 62
	var b float64 = float64(a)

	x := a

	fmt.Println("\nTypes: ")
	fmt.Println("a is:", reflect.TypeOf(a))
	fmt.Println("b is:", reflect.TypeOf(b))
	fmt.Println("x is:", reflect.TypeOf(x))
}

/*
Variables out of a function are used to be defined with 'var'.
Variables inside a function commonly have no type specification. Go will know for us.
Once a type is assigned to a variable this type can't chance once program is compiled.

Ways to define variables:
1. num1, num2 := 5.6, 9.5 (by default will be float64)
2. var num1, num2 float64 = 5.6, 9.5
3. num1, num2 float64 = 5.6, 9.5
*/
