//key-points about slice
package main

import (
	"fmt"
)

func reverse(s []int) {
	for i, j := 0, len(s)-1; i < j; i, j = i+1, j-1 {
		s[i], s[j] = s[j], s[i]
	}
}

func nonempty(xx []string) []string {
	i := 0
	for _, s := range xx {
		if s != "" {
			xx[i] = s
			i++
		}
	}

	return xx[:i]
}

func nonempty2(xx []string) []string {
	out := xx[:0]
	for _, s := range xx {
		if s != "" {
			out = append(out, s)
		}
	}

	return out
}

func remove(slice []int, i int) []int {
	copy(slice[i:], slice[i+1:])
	return slice[:len(slice)-1]
}

func main() {
	months := [...]string{"January", "February", "March", "April", "May", "June", "July", "August", "September", "Octorber", "November", "December"}
	Q2 := months[4:7]
	summer := months[6:9]
	fmt.Println(Q2)
	fmt.Println(summer)

	months[4] = "none" //改变底层数组数据
	fmt.Println(Q2)    //查看改动后，对分片的影响

	fmt.Println(cap(Q2), len(Q2))

	//fmt.Println(Q2[:20]) //slice out of range

	endlessSummer := summer[:5]
	fmt.Println(endlessSummer)

	a := [...]int{9, 8, 7, 6, 5, 4, 3, 2, 1}
	//	reverse(a) //会报类型错误 type [9]int as type []int
	reverse(a[:])
	fmt.Println(a)

	x := []int{}
	x = append(x, 1)
	x = append(x, 2, 3)
	x = append(x, 4, 5, 6)
	x = append(x, x...) //省略号，不定长参数
	fmt.Println(x)

	data := []string{"hello", "", "world!"}
	fmt.Printf("%q\n", nonempty(data))
	fmt.Printf("%q\n", data)

	s1 := []int{5, 6, 7, 8, 9}
	fmt.Println(remove(s1, 2))
}
