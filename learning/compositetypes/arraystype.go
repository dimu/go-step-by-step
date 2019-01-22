package main

import (
	"fmt"
)

/*
composite type——array
*/
func main() {
	var q [3]int = [3]int{1, 2, 3}
	var r [3]int = [3]int{1, 2} //缺省值为zero
	fmt.Println(len(r), r[2], q[2])

	t := [...]int{3, 4, 5} //[3]int, length determined by the initializer
	fmt.Printf("%T\n", t)

	//t = [4]int{1, 2, 3, 4} //type assignment error

	k := [...]int{49: -1}
	fmt.Printf("%T %d\n", k, k[49])

	fmt.Println(q == r) //arrays comparable, elements are equal
	r[2] = 3
	fmt.Println(q == r)
}
