package main

import (
	"fmt"
)

var i int = 0

func main() {
	i := 123
	p := &i
	fmt.Println(p)
	fmt.Println(*p)
	fmt.Println(&p)
	fmt.Println(&i)

	var x, y int
	fmt.Println(&x)
	fmt.Println(&x == &x, &x == &y, &x == nil)

	t := new(int) //create an unnamed variable of type T
	fmt.Println("new function, default value is", *t)
	*t = 2
	fmt.Println("new function, changed value is", *t)

	//every time return new address, but have the same value
	fmt.Println(f() == f())
	fmt.Println(f())          //return the address
	fmt.Println(*f() == *f()) //have the same value
	*f() = 2
	fmt.Println(*f())

	i2 := 2
	i3 := &i2
	fmt.Println(&i2)
	fmt.Println(i3)
	fmt.Println(&i3)
	fmt.Println(&(*i3))

	f2 := 1
	inc(&f2)
	fmt.Println(inc(&f2))

	*f3() = 2
	fmt.Println(*f3())
}

func f() *int {
	v := 1
	return &v
}

func f3() *int {
	return &i
}

/*

 */
func inc(p *int) int {
	*p++
	return *p
}
