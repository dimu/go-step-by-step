package main

import (
	"fmt"
	"time"
)

type Employee struct {
	ID        int //字段名大写表示导出字段，小写为非导出字段，这样达到访问控制机制
	Name      string
	Address   string
	DoB       time.Time
	Position  string
	Salary    int
	ManagerID int
}

func main() {
	var employee Employee
	//如何输出变量的内存地址？
	fmt.Println(employee)
	fmt.Printf("the address of employee var is %x\n", &employee)
	employee.Position = "java programmer"
	position := &employee.Position
	fmt.Println("")
	*position = "senior " + *position
	fmt.Println(employee.Position)
	fmt.Println(*position)

	employee1 := &employee
	employee1.Position += "(wonderful) "
	fmt.Println(employee1.Position)
	(*employee1).Position += "(second) "
	fmt.Println(employee1.Position)
}
