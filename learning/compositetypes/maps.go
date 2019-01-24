package main

import (
	"fmt"
)

func main() {
	map2 := make(map[string]string)

	map1 := map[string]string{
		"name": "kitty",
		"age":  "24",
	}

	map2["name"] = "kitty"
	map2["age"] = "24"
	map2["age"] = "25" //后面的值覆盖前面的值

	fmt.Println(map1)
	fmt.Println(map2)

	delete(map2, "age")
	fmt.Println(map2)

	fmt.Println(map2["li"]) //没有的key值，返回默认值，字符串时为空字符串

	//循环遍历map
	for key, value := range map1 {
		fmt.Printf("%s\t%s\n", key, value)
	}

	var map3 map[string]int
	fmt.Println(map3 == nil)
	//	map3["age"] = 11 //can't assignment to entry in nil map,必须先分配空间再赋值

}
