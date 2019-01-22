package main

import (
	"fmt"
	"unicode/utf8"
)

func main() {
	s := "hello world"
	fmt.Println(&s)
	t := s
	fmt.Println(&t)
	s += "bye bye"
	fmt.Println(&s)

	s += "dfssfdefsfdfdfdsdfdfdsffffffffdfdsafdasfadsffffffffffssssssssssssssssssssssssss"
	fmt.Println(&s)
	//s[0] = 'a'
	fmt.Println(t)

	s1 := s[:5]
	fmt.Println(&s1, s1)
	s1 = "abc"
	fmt.Println(&s1, s1)

	s2 := "高大上"
	fmt.Println(len(s2))

	//raw string literals，使用``表示
	s3 := `go go go
		age: 18
		name: "123"
			
	`
	fmt.Println(s3)

	s4 := "\xe4\xb8\x96\xe7\x95\x8c" //世界的utf-8编码
	fmt.Println(s4)
	s4 = "\u4e16\u754c"
	fmt.Println(s4)
	s4 = "\U00004e16\U0000754c"
	fmt.Println(s4)

	s4 = "baby, 你好!"
	fmt.Println(len(s4))                    //13
	fmt.Println(utf8.RuneCountInString(s4)) //9

	s4 = "プログラム"
	fmt.Printf("% x\n", s4)
	fmt.Printf("%x\n", []rune(s4))
}
