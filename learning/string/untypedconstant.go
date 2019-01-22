package main

import (
	"fmt"
)

func main() {
	var f float64 = 3 + 0i

	f = 2

	fmt.Printf("%f\n", f)

	i := 0
	fmt.Printf("%T\n", i)

	fmt.Printf("%T\n", 0i)
}
