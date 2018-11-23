package main

import (
	"fmt"
	"math"
	"math/rand"
)

// In Go, main() wil be always called by default
func main() {
	foo()
}

func foo() {
	// In Go if the first letter is capitalized 'P' on Println() and 'S' on Sqrt() function will be exported by Go.
	// In lower case, will be considered as an internal function.
	fmt.Println("The square root of 4 is", math.Sqrt(4))

	fmt.Println("A number from 0-90:", rand.Intn(100))
}
