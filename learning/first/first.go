// first
package main

import (
	"fmt"
	"math"
)

//package variable
var sex bool = true

/*Go's basic types are bool, string, int, int8, int16, int32, int64
unit unit8 unit16 unit32 uint64 uintptr
byte //alias for uint8
rune //alias for int32
float32 float64
complex64 complex128
*/
//var test uint16 = 65537 //overflow unit16

//Zero values:variables declared without an explicit initial value are given
//zero value: 0 for numeric types, false for boolean type, "" for strings

//type conversions: go requires an explicit conversion
var i int = 12
var f float32 = float32(i)
var u uint32 = uint32(f)

//Type inference
//the variable's type if inferred from the value on the right hand side

func main() {
	fmt.Println("Hello World!")
	fmt.Println(math.Pi)
	a, b := swap("first", "second")
	fmt.Println(a, b)
	fmt.Println(split(17))

	//function variable
	var count int

	//short variable declarations, only inside a function
	total := 256
	fmt.Println(sex, count, total)

	//type reference
	v := 0.33 + 0.44i
	fmt.Printf("v is of type %T\n", v)

	//constants declared, can't use := syntax
	const con = 123
	fmt.Println("const values ", con)

	fmt.Println("total is", forloop())

	//定义在同一包下不同文件下的常见
	fmt.Println(param)

}

//golang can return multiple results, enhancement
func swap(x, y string) (string, string) {
	return y, x
}

//named return values, 返回值名称与方法体一致
func split(sum int) (x, y int) {
	x = sum * 4 / 9
	y = sum - x
	return
}

func forloop() int {
	sum := 0
	//for loop no parentheses and the braces {} are required
	for i := 0; i < 10; i++ {
		sum += i
	}

	//the init and post statement are optional, so it can also write as follow:
	//for ; i < 10; {
	//}

	//infinite loop
	//	for {

	//	}

	return sum
}
